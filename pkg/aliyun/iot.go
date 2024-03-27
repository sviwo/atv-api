package aliyun

import (
	"encoding/json"
	"fmt"
	iot20180120 "github.com/alibabacloud-go/iot-20180120/v6/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"golang.org/x/net/context"
	"net/http"
	"strings"
	"sviwo/internal/consts/enums"
)

var IotClient *iot20180120.Client

// 命令下发
func SetDevicePropertyRequest(ctx context.Context, productKey string, deviceName string, messageContent string) error {
	pubRequest := &iot20180120.SetDevicePropertyRequest{
		ProductKey:    tea.String(productKey),
		IotInstanceId: tea.String(g.Cfg().MustGet(ctx, "aliyun.iot.iotInstanceId").String()),
		DeviceName:    tea.String(deviceName),
		Items:         tea.String(messageContent), //eyJ0ZXN0IjoidGFzayBwdWIgYnJvYWRjYXN0In0=
	}
	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		resp, _err := IotClient.SetDevicePropertyWithOptions(pubRequest, runtime)
		if _err != nil {
			return _err
		}
		glog.Printf(ctx, "parseJson Unmarshal err:%v", util.ToJSONString(resp))
		//glog.Infof(ctx, "parseJson Unmarshal err:%v", util.ToJSONString(resp))
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

// RegisterDevice 拿到注册产品下的设备需要的参数
func RegisterDevice(ctx context.Context, productKey, deviceName string) (data *iot20180120.RegisterDeviceResponseBodyData, err error) {
	registerDeviceRequest := &iot20180120.RegisterDeviceRequest{
		ProductKey:    tea.String(productKey),
		DeviceName:    tea.String(deviceName),
		IotInstanceId: tea.String(g.Cfg().MustGet(ctx, "aliyun.iot.iotInstanceId").String()),
	}
	resp, err := IotClient.RegisterDeviceWithOptions(registerDeviceRequest, &util.RuntimeOptions{})
	if err != nil {
		return nil, err
	}
	if *resp.StatusCode != http.StatusOK || !*resp.Body.Success {
		glog.Error(ctx, util.ToJSONString(resp.Body.ErrorMessage))
		return nil, gerror.NewCode(enums.RequestThirdInterFaceFail)
	}
	data = resp.Body.Data
	return
}

// DeleteDevice 删除注册到产品下的设备
func DeleteDevice(ctx context.Context, productKey, deviceName string) error {
	deleteDeviceRequest := &iot20180120.DeleteDeviceRequest{
		ProductKey: tea.String(productKey),
		DeviceName: tea.String(deviceName),
	}
	resp, err := IotClient.DeleteDeviceWithOptions(deleteDeviceRequest, &util.RuntimeOptions{})
	if err != nil {
		return err
	}
	if *resp.StatusCode != http.StatusOK || !*resp.Body.Success {
		glog.Error(ctx, util.ToJSONString(resp.Body.ErrorMessage))
		return gerror.NewCode(enums.Fail)
	}
	return nil
}
