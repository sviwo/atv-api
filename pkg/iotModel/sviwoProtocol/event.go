package sviwoProtocol

// 相关topci信息
const ()

// 事件上报结构体
type (
	// 事件上报请求报文
	ReportEventReq struct {
		//Sys             SysInfo                `json:"sys"`
		RequestId       string         `json:"requestId"`
		ProductKey      string         `json:"productKey"`
		DeviceName      string         `json:"deviceName"`
		DeviceType      string         `json:"deviceType"`
		GmtCreate       int64          `json:"gmtCreate"`
		IotId           string         `json:"iotId"`
		CheckFailedData interface{}    `json:"checkFailedData"`
		Identifier      string         `json:"identifier"`
		Name            string         `json:"name"`
		Type            string         `json:"type"`
		Time            int64          `json:"time"`
		Value           map[string]any `json:"value"`
	}
	ReportEventParams struct {
		Value    map[string]any `json:"value"`
		CreateAt int64          `json:"time"`
	}
	// 事件上报响应报文
	ReportEventReply struct {
		Code int `json:"code"`
		Data struct {
		} `json:"data"`
		Id      string `json:"id"`
		Message string `json:"message"`
		Method  string `json:"method"`
		Version string `json:"version"`
	}
)
