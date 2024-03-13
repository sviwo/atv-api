package aliyun

import (
	"encoding/json"
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	iot20180120 "github.com/alibabacloud-go/iot-20180120/v6/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"golang.org/x/net/context"
	"strings"
)

var iotClient *iot20180120.Client

func InitAliyunClient() error {
	var ctx = gctx.New()
	accessKeyID := g.Cfg().MustGet(ctx, "aliyun.accessKeyID").String()
	accessKeySecret := g.Cfg().MustGet(ctx, "aliyun.accessKeySecret").String()
	host := g.Cfg().MustGet(ctx, "aliyun.iot.api.host").String()
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: &accessKeyID,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: &accessKeySecret,
	}
	config.Endpoint = tea.String(host)
	var err error
	iotClient, err = iot20180120.NewClient(config)
	return err
}

// 命令下发
func PubRequest(ctx context.Context, productKey string, deviceName string, topic string, messageContent string) error {
	pubRequest := &iot20180120.PubRequest{
		ProductKey:     tea.String(productKey),
		IotInstanceId:  tea.String(g.Cfg().MustGet(ctx, "aliyun.iot.iotInstanceId").String()),
		DeviceName:     tea.String(deviceName),
		TopicFullName:  tea.String(topic),          ///${productKey}/${deviceName}/user/get
		MessageContent: tea.String(messageContent), //eyJ0ZXN0IjoidGFzayBwdWIgYnJvYWRjYXN0In0=
	}
	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		resp, _err := iotClient.PubWithOptions(pubRequest, runtime)
		if _err != nil {
			return _err
		}
		glog.Infof(ctx, "parseJson Unmarshal err:%v", util.ToJSONString(resp))
		return nil
	}()
	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		// 错误 message
		glog.Errorf(ctx, tea.StringValue(error.Message))
		// 诊断地址
		var data interface{}
		d := json.NewDecoder(strings.NewReader(tea.StringValue(error.Data)))
		d.Decode(&data)
		if m, ok := data.(map[string]interface{}); ok {
			recommend, _ := m["Recommend"]
			fmt.Println(recommend)
		}
		_, tryErr = util.AssertAsString(error.Message)
		if tryErr != nil {
			return tryErr
		}
	}
	return tryErr
}
