package sviwoProtocol

type (
	// ack 标记是否需要回复，1需要回复，0不需要回复
	SysInfo struct {
		Ack int `json:"ack"`
	}
	// 属性节点，包含属性值和时间
	PropertyNode struct {
		Value      interface{} `json:"value"`
		CreateTime int64       `json:"time"`
	}
	//公共参数
	CommonParam struct {
		RequestId       string      `json:"requestId"`
		ProductKey      string      `json:"productKey"`
		DeviceName      string      `json:"deviceName"`
		DeviceType      string      `json:"deviceType"`
		GmtCreate       int64       `json:"gmtCreate"`
		IotId           string      `json:"iotId"`
		CheckFailedData interface{} `json:"checkFailedData"`
	}
)

const (
	NeedAck = 1
)
