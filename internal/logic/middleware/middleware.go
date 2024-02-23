package middleware

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gutil"
	"github.com/gogf/gf/v2/util/gvalid"
	"io"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net/url"
	"reflect"
	"strings"
	rcode "sviwo/internal/logic/biz/enums"
	"sviwo/internal/model"
	"sviwo/internal/service"
	"sviwo/utility/common"
	ecc "sviwo/utility/encrypt"
	"sviwo/utility/response"
)

type sMiddleware struct{}

func init() {
	service.RegisterMiddleware(New())
}

func New() *sMiddleware {
	return &sMiddleware{}
}

func (s *sMiddleware) ErrorHandler(r *ghttp.Request) {
	r.Middleware.Next()
	err := r.GetError()
	if err == nil {
		return
	}
	glog.Error(r.GetCtx(), err)
	r.Response.ClearBuffer()
	var gvalidErr gvalid.Error
	errors.As(err, &gvalidErr)
	if !gutil.IsEmpty(gvalidErr) {
		rule, err := gvalidErr.FirstRule()
		if "required" == rule {
			response.Json(
				r, rcode.New(rcode.IllegalArgument.Code(), err.Error()),
				nil,
			)
		}
	} else if reflect.TypeOf(err.(gerror.ICode).Code()).Name() == reflect.TypeOf(rcode.OpenResponseEnum{}).Name() {
		response.Json(r, err.(gerror.ICode).Code(), nil)
	} else {
		response.FailMsg(r)
	}
}

// 统一返回处理中间件
func (s *sMiddleware) ResponseHandler(r *ghttp.Request) {
	r.Middleware.Next()
	// 如果已经有返回内容，那么该中间件什么也不做
	if r.Response.BufferLength() > 0 {
		return
	}
	var (
		err  = r.GetError()
		res  = r.GetHandlerResponse()
		code = gerror.Code(err)
	)
	if err != nil {
		if code == gcode.CodeNil {
			code = rcode.ApiException
		}
		response.JsonExit(r, code, nil)
	} else {
		response.SuccessMsg(r, res)
	}
}

// 自定义上下文对象
func (s *sMiddleware) Ctx(r *ghttp.Request) {
	// 初始化，务必最开始执行
	customCtx := &model.Context{
		Data: make(g.Map),
	}
	service.BizCtx().Init(r, customCtx)
	// 将自定义的上下文对象传递到模板变量中使用
	r.Assigns(g.Map{
		"Context": customCtx,
	})
	r.Middleware.Next()
}

func (s *sMiddleware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

// use middleware func for router
func (s *sMiddleware) DecodeData(r *ghttp.Request) {
	contentType := r.Request.Header.Get("Content-Type")
	isJsonRequest := strings.Contains(contentType, "application/json")
	isFileRequest := strings.Contains(contentType, "multipart/form-data")
	isFormUrl := strings.Contains(contentType, "application/x-www-form-urlencoded")
	//get user private key
	privateKey := ""
	if r.Request.Method == "GET" { //if use get

		parseQuery(r, privateKey)

	} else if isJsonRequest { //if use post json

		parseJson(r, privateKey)

	} else if isFormUrl { //if use post form

		parseForm(r, privateKey)

	} else if isFileRequest { //if use file

		parseFile(r, privateKey)
	}
	r.Middleware.Next()
}

func parseQuery(r *ghttp.Request, privateKey string) {
	encryptString := r.Get("data", "").String()
	if len(encryptString) < 1 {
		panic(rcode.RequestParamTypeError)
	}
	queryData, err := ecc.EccDecryptByBase64(encryptString, privateKey)
	if err != nil {
		panic(rcode.RequestParamTypeError)
	}
	dataMap := gjson.New(queryData).Map()
	var args []string
	var logs []string
	for k, v := range dataMap {
		r.SetQuery(k, v) //r.SetForm(key, val)
		val := common.ToStr(v)
		args = append(args, fmt.Sprintf("%s=%s", k, url.QueryEscape(val)))
		logs = append(logs, fmt.Sprintf("%s=%s", k, val))
	}
	glog.Infof(r.GetCtx(), "parseQuery  url:%s, md5key:%s, encryptString:%s, decrypt data:%s", r.Request.URL.String(), privateKey, encryptString, strings.Join(logs, "&"))
	r.Request.URL.RawQuery = strings.Join(args, "&")
}

func parseJson(r *ghttp.Request, privateKey string) {
	payload := r.GetBody() // 把body再写回去,不然别的地方取不到
	///解密body数据 请求的json是{"encryptString":{value}} value含有gcm的12字节nonce,实际长度大于32
	if payload != nil && len(payload) > 20 {
		var jsonData encryptJson
		glog.Infof(r.GetCtx(), "parseJson url:%s md5key:%s,old data:%s,", r.Request.URL.String(), privateKey, string(payload))
		err := json.Unmarshal(payload, &jsonData)
		if err != nil {
			glog.Infof(r.GetCtx(), "parseJson Unmarshal err:%v", err)
			panic(err)
		}
		payloadText := jsonData.EncryptString
		if len(payloadText) > 0 {
			payloadByte, err := ecc.EccDecryptByBase64(payloadText, privateKey)
			if err != nil {
				glog.Infof(r.GetCtx(), "parseJson GcmDecryptByte err:%v", err)
				panic(err)
			}
			payload = payloadByte
			glog.Infof(r.GetCtx(), "parseJson url:%s md5key:%s,encryptString:%s,decrypt data:%s", r.Request.URL.String(), privateKey, jsonData.EncryptString, string(payload))
		}
		r.Request.Body = io.NopCloser(bytes.NewBuffer(payload))
	}
}

type encryptJson struct {
	EncryptString string `json:"encryptString"`
}

func parseForm(r *ghttp.Request, privateKey string) {
	payload := r.GetBody() // 把body再写回去,不然别的地方取不到
	if payload != nil && len(payload) > 20 {
		var jsonData encryptJson
		glog.Infof(r.GetCtx(), "parseForm url:%s md5key:%s,old data:%s,", r.Request.URL.String(), privateKey, string(payload))
		values, err := url.ParseQuery(string(payload))
		if err != nil {
			glog.Errorf(r.GetCtx(), "parseForm ParseQuery err:%v", err)
			panic(err)
		}
		payloadText := values.Get("encryptString")
		if len(payloadText) > 0 {
			mapData, err := decryptString(payloadText, privateKey)
			if err != nil {
				glog.Errorf(r.GetCtx(), "parseForm err:%v", err)
				panic(err)
			}
			for k, v := range mapData {
				values.Add(k, common.ToStr(v))
			}
			formData := values.Encode()
			glog.Infof(r.GetCtx(), " parseForm url:%s md5key:%s,encryptString:%s,decrypt data:%s", r.Request.URL.String(), privateKey, jsonData.EncryptString, formData)
			payload = []byte(formData)
		}
	}
	r.Request.Body = io.NopCloser(bytes.NewBuffer(payload))
}

func parseFile(r *ghttp.Request, privateKey string) {
	contentType := r.Request.Header.Get("Content-Type")
	_, params, _ := mime.ParseMediaType(contentType)
	boundary, ok := params["boundary"]
	if !ok {
		panic(rcode.RequestParamTypeError)
	}
	bodyBuf := &bytes.Buffer{}
	wr := multipart.NewWriter(bodyBuf)
	mr := multipart.NewReader(r.Request.Body, boundary)
	for {
		p, err := mr.NextPart() //p的类型为Part
		if err == io.EOF {
			break
		}
		if err != nil {
			glog.Infof(r.GetCtx(), "NextPart err:%v", err)
			break
		}
		fileByte, err := ioutil.ReadAll(p)
		if err != nil {
			glog.Infof(r.GetCtx(), "ReadAll err:%v", err)
			break
		}
		pName := p.FormName()
		fileName := p.FileName()
		if len(fileName) < 1 {
			if pName == "encryptString" {
				formData, err := decryptString(privateKey, string(fileByte))
				if err != nil {
					glog.Errorf(r.GetCtx(), "writeFile gcmDecryptString err:%v", err)
					break
				}
				for k, v := range formData {
					val := common.ToStr(v)
					err = wr.WriteField(k, val)
					if err != nil {
						glog.Errorf(r.GetCtx(), "writeFile WriteField :%s=%s, err:%v", k, val, err)
						break
					}
				}
			} else {
				wr.WriteField(pName, string(fileByte))
			}
		} else {
			tmp, err := wr.CreateFormFile(pName, fileName)
			if err != nil {
				glog.Errorf(r.GetCtx(), "parseFile CreateFormFile err:%v", err)
				continue
			}
			_, err = tmp.Write(fileByte)
			if err != nil {
				panic(err)
			}
		}
	}
	//写结尾标志
	_ = wr.Close()
	r.Request.Header.Set("Content-Type", wr.FormDataContentType())
	r.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBuf.Bytes()))
}

func decryptString(privateKey, encryptString string) (map[string]interface{}, error) {
	formData := make(map[string]interface{}, 0)
	if len(encryptString) < 1 {
		return formData, nil
	}
	plaintext, err := ecc.EccDecryptByBase64(encryptString, privateKey)
	if err != nil {
		return formData, err
	}
	if len(plaintext) < 3 {
		//plaintext should json  {}
		return formData, nil
	}
	err = json.Unmarshal([]byte(plaintext), &formData)
	if err != nil {
		return formData, err
	}
	return formData, nil
}