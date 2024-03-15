package sviwoProtocol

// 平台下发topic
const ()

// 设备上报
const (
	//设备上报事件请求topic /k0ugjmf1ois/sviwo_atv/thing/event/property/post
	PropertySubRequestTopic = "/*/*/thing/event/property/post"

	//设备上下线通知topic /as/mqtt/status/k0ugjmf1ois/sviwo_atv
	PropertyOnOfflineTopic = "/as/mqtt/status/*/*"

	//设备事件通知topic /k0ugjmf1ois/sviwo_atv/thing/event/lowbattey/post
	EvnetTopic = "/*/*/thing/event/*/post"
)
