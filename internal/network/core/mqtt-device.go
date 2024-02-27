package core

import (
	"context"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/gogf/gf/v2/frame/g"
	"strings"
	"sviwo/internal/consts"
	"sviwo/internal/mqtt"
)

var subMap = map[consts.Topic]func(context.Context, MQTT.Client, MQTT.Message){
	//device 处理的topic
	consts.TopicDeviceData: deviceDataHandler,

	//dataBus  处理的topic
	consts.CommonDataBusPrefix + consts.DataBusOnline:         onlineHandler,
	consts.CommonDataBusPrefix + consts.DataBusOffline:        offlineHandler,
	consts.CommonDataBusPrefix + consts.DataBusPropertyReport: reportPropertiesHandler,
}

func StartSubscriber(ctx context.Context) error {
	for topic, f := range subMap {
		if err := mqtt.Subscribe(ctx, string(topic), f); err != nil {
			return err
		}
	}
	return nil
}

func deviceDataHandler(ctx context.Context, client MQTT.Client, message MQTT.Message) {
	topicInfo := strings.Split(message.Topic(), "/")
	if len(topicInfo) != 3 {
		g.Log().Error(ctx, fmt.Sprintf("topic:%s is illegal, message(%s) ignored", message.Topic(), string(message.Payload())))
		return
	}
	//具体业务

}
