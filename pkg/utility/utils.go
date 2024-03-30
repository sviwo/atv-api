package utility

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/container/glist"
	"github.com/gogf/gf/v2/crypto/gaes"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gbase64"
	"github.com/gogf/gf/v2/encoding/gurl"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"gopkg.in/gomail.v2"
	"strconv"
	"strings"
	"sviwo/internal/boot"
	"sviwo/internal/consts/enums"
	"time"
)

// 密码加密
func EncryptPassword(password, salt string, num int) string {
	if num <= 0 {
		return gmd5.MustEncryptString(password + gmd5.MustEncryptString(salt))
	}
	for i := 0; i < num; i++ {
		password = gmd5.MustEncryptString(password)
	}
	return gmd5.MustEncryptString(password + gmd5.MustEncryptString(salt))
}

// 获取当前请求接口域名
func GetDomain(r *ghttp.Request) (string, error) {
	pathInfo, err := gurl.ParseURL(r.GetUrl(), -1)
	if err != nil {
		g.Log().Error(context.Background(), err)
		err = gerror.New("解析附件路径失败")
		return "", err
	}
	return fmt.Sprintf("%s://%s:%s/", pathInfo["scheme"], pathInfo["host"], pathInfo["port"]), nil
}

// GetUserAgent 获取user-agent
func GetUserAgent(ctx context.Context) string {
	return ghttp.RequestFromCtx(ctx).Header.Get("User-Agent")
}

// GetDbConfig get db boot
func GetDbConfig() (cfg *gdb.ConfigNode, err error) {
	cfg = g.DB().GetConfig()
	err = ParseDSN(cfg)
	return
}

// ParseDSN parses the DSN string to a Config
func ParseDSN(cfg *gdb.ConfigNode) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = gerror.New(r.(string))
		}
	}()
	dsn := cfg.Link
	if dsn == "" {
		return
	}
	foundSlash := false
	// gfast:123456@tcp(192.168.0.212:3306)/gfast-v2
	for i := len(dsn) - 1; i >= 0; i-- {
		if dsn[i] == '/' {
			foundSlash = true
			var j, k int

			// left part is empty if i <= 0
			if i > 0 {
				// [username[:password]@][protocol[(address)]]
				// Find the last '@' in dsn[:i]
				for j = i; j >= 0; j-- {
					if dsn[j] == '@' {
						// username[:password]
						// Find the first ':' in dsn[:j]
						for k = 0; k < j; k++ {
							if dsn[k] == ':' {
								cfg.Pass = dsn[k+1 : j]
								cfg.User = dsn[:k]
								break
							}
						}
						break
					}
				}

				// gfast:123456@tcp(192.168.0.212:3306)/gfast-v2
				// [protocol[(address)]]
				// Find the first '(' in dsn[j+1:i]
				var h int
				for k = j + 1; k < i; k++ {
					if dsn[k] == '(' {
						// dsn[i-1] must be == ')' if an address is specified
						if dsn[i-1] != ')' {
							if strings.ContainsRune(dsn[k+1:i], ')') {
								panic("invalid DSN: did you forget to escape a param value?")
							}
							panic("invalid DSN: network address not terminated (missing closing brace)")
						}
						for h = k + 1; h < i-1; h++ {
							if dsn[h] == ':' {
								cfg.Host = dsn[k+1 : h]
								cfg.Port = dsn[h+1 : i-1]
								break
							}
						}
						break
					}
				}
			}
			for j = i + 1; j < len(dsn); j++ {
				if dsn[j] == '?' {
					cfg.Name = dsn[i+1 : j]
					break
				} else {
					cfg.Name = dsn[i+1:]
				}
			}
			break
		}
	}
	if !foundSlash && len(dsn) > 0 {
		panic("invalid DSN: missing the slash separating the database name")
	}
	return
}

// 获取附件真实路径
func GetRealFilesUrl(r *ghttp.Request, path string) (realPath string, err error) {
	if gstr.ContainsI(path, "http") {
		realPath = path
		return
	}
	realPath, err = GetDomain(r)
	if err != nil {
		return
	}
	realPath = realPath + path
	return
}

// 获取附件相对路径
func GetFilesPath(fileUrl string) (path string, err error) {
	gvar_type, err := g.Cfg().Get(context.Background(), "file.type")
	if err != nil {
		return "", err
	}
	upType := gstr.ToLower(gvar_type.String())

	gvar_UpPath, err := g.Cfg().Get(context.Background(), "file.local.UpPath")
	if err != nil {
		return "", err
	}
	upPath := gstr.Trim(gvar_UpPath.String(), "/")
	if upType != "local" || (upType == "local" && !gstr.ContainsI(fileUrl, upPath)) {
		path = fileUrl
		return
	}
	pathInfo, err := gurl.ParseURL(fileUrl, 32)
	if err != nil {
		g.Log().Error(context.Background(), err)
		err = gerror.New("解析附件路径失败")
		return
	}
	pos := gstr.PosI(pathInfo["path"], upPath)
	if pos >= 0 {
		path = gstr.SubStr(pathInfo["path"], pos)
	}
	return
}

// 货币转化为分
func CurrencyLong(currency interface{}) int64 {
	strArr := gstr.Split(gconv.String(currency), ".")
	switch len(strArr) {
	case 1:
		return gconv.Int64(strArr[0]) * 100
	case 2:
		if len(strArr[1]) == 1 {
			strArr[1] += "0"
		} else if len(strArr[1]) > 2 {
			strArr[1] = gstr.SubStr(strArr[1], 0, 2)
		}
		return gconv.Int64(strArr[0])*100 + gconv.Int64(strArr[1])
	}
	return 0
}

//func GetExcPath() string {
//	file, _ := exec.LookPath(os.Args[0])
//	// 获取包含可执行文件名称的路径
//	path, _ := filepath.Abs(file)
//	// 获取可执行文件所在目录
//	index := strings.LastIndex(path, string(os.PathSeparator))
//	ret := path[:index]
//	return strings.Replace(ret, "\\", "/", -1)
//}

// 查看数组中是否有对应的值
func InArray(needle string, haystack []string) bool {
	for _, v := range haystack {
		if needle == v {
			return true
		}
	}
	return false
}

func ToStr(value interface{}) string {
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	case time.Time:
		it := value.(time.Time)
		key = it.String()
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}

// 发送邮件
func SendEmail(subject, body, to string) error {
	m := gomail.NewMessage()
	//发送人
	m.SetHeader("From", g.Cfg().MustGet(context.TODO(), "email.username").String())
	//接收人
	m.SetHeader("To", to)
	//主题
	m.SetHeader("Subject", subject)
	//内容
	m.SetBody("text/html", body)
	// 发送邮件
	if err := boot.NewDialer.DialAndSend(m); err != nil {
		glog.Error(context.Background(), "验证发送失败，错误原因==", err)
		return gerror.NewCode(enums.VftCodeSendFailed)
	}
	return nil
}

/*
GfTokenDecryptToken GfToken解密
*/
func GfTokenDecryptToken(ctx context.Context, token string) string {
	if token == "" {
		panic(gerror.NewCode(enums.RequestMissingParam))
	}
	parts := strings.SplitN(token, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		panic(gerror.NewCode(enums.RequestMethodTypeError))
	} else if parts[1] == "" {
		panic(gerror.NewCode(enums.RequestMethodTypeError))
	}
	token64, err := gbase64.Decode([]byte(parts[1]))
	if err != nil {
		panic(err)
	}
	decryptToken, err2 := gaes.Decrypt(token64, g.Cfg().MustGet(ctx, "gfToken.encryptKey").Bytes())
	if err2 != nil {
		panic(err2)
	}
	return gstr.Split(string(decryptToken), "_")[0]
}

/*
BuildTree 构建树形结构
*/
func BuildTree(array []map[string]interface{}) (treeDataList []interface{}) {
	temp := make(map[string]map[string]interface{})
	//此list是为了保持数据顺序使用
	orderList := glist.New()
	for _, m := range array {
		id := gconv.String(m["id"])
		temp[id] = m
		orderList.PushBack(id)
	}
	for range temp {
		//根据orderList有序的id获取对应的map对象
		mapObj := temp[gconv.String(orderList.PopFront())]
		//使用此map对象的parentId获取他的父级
		parentObj := temp[gconv.String(mapObj["parentId"])]
		if parentObj == nil {
			//对象本身没有父级则说明已经是顶级，则将自己追加为返回值元素
			treeDataList = append(treeDataList, mapObj)
		} else {
			//对象本身有父级，则将自己追加为父级的children
			parentObj["children"] = append(gconv.SliceAny(parentObj["children"]), mapObj)
		}
	}
	return
}

// RemoveDuplicationMap 数组去重
func RemoveDuplicationMap(arr []string) []string {
	set := make(map[string]struct{}, len(arr))
	j := 0
	for _, v := range arr {
		_, ok := set[v]
		if ok {
			continue
		}
		set[v] = struct{}{}
		arr[j] = v
		j++
	}
	return arr[:j]
}

func RemoveRepeatedElementAndEmpty(arr []int) []int {
	newArr := make([]int, 0)
	for _, item := range arr {
		repeat := false
		if len(newArr) > 0 {
			for _, v := range newArr {
				if v == item {
					repeat = true
					break
				}
			}
		}
		if repeat {
			continue
		}
		newArr = append(newArr, item)
	}
	return newArr
}

// FileSize 字节的单位转换 保留两位小数
func FileSize(fileSize int64) string {
	units := []string{"B", "KB", "MB", "GB", "TB", "EB"}
	var size = float64(fileSize)
	var i int
	for i = 0; size > 1024; i++ {
		size /= 1024
	}
	return fmt.Sprintf("%.2f %s", size, units[i])
}
