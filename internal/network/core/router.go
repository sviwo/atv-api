package core

import (
	"context"
	"errors"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"pack.ag/amqp"
	"strings"
	"sviwo/internal/mqtt"
	"sviwo/pkg/gpool"
	"sviwo/pkg/iotModel/topicModel"
	"sync"
)

type filterMsgFunc struct {
	name string
	f    func(message MQTT.Message) MQTT.Message
}

type (
	SubMap struct {
		sync.RWMutex
		// 订阅主题
		subTopics map[string]handleFunc
		// 订阅处理前filter
		subTopicBeforeHandlerChain []filterMsgFunc
		// 订阅处理后filter
		subTopicAfterHandlerChain []filterMsgFunc
	}
	handleFunc struct {
		logType string
		f       func(context.Context, topicModel.TopicHandlerData) error
	}
)

var subMapInfo = SubMap{
	subTopics:                  make(map[string]handleFunc),
	subTopicBeforeHandlerChain: make([]filterMsgFunc, 0),
	subTopicAfterHandlerChain:  make([]filterMsgFunc, 0),
}

// 控制携程数量
var gPool = gpool.NewGPool(10000)

// HandleMessage 所有订阅的入口
func (s *SubMap) HandleMessage(ctx context.Context, handleF handleFunc) func(context.Context, MQTT.Client, MQTT.Message) {
	return func(ctx context.Context, client MQTT.Client, message MQTT.Message) {
		gPool.Go(func(ctx context.Context) error {
			// 根据topic拿到deviceKey和productKey
			topicInfo := strings.Split(message.Topic(), "/")
			if len(topicInfo) < 3 {
				//todo 是否入库，前端展示
				return errors.New(fmt.Sprintf("topic:%s is illegal, message(%s) ignored", message.Topic(), string(message.Payload())))
			}
			if topicInfo[1] != "sys" && topicInfo[1] != "ota" {
				//todo 非sys开头的topic不处理
				return errors.New(fmt.Sprintf("topic:%s is not supported,message(%s) ignored", message.Topic(), string(message.Payload())))
			}
			return nil
		})
	}
}

func RegisterSubTopicHandler(topic, logType string, handler func(context.Context, topicModel.TopicHandlerData) error) error {
	subMapInfo.Lock()
	defer subMapInfo.Unlock()
	if _, ok := subMapInfo.subTopics[topic]; ok {
		return fmt.Errorf("topic %s already registered", topic)
	}
	subMapInfo.subTopics[topic] = handleFunc{f: handler, logType: logType}
	return nil
}

func StartSubscriber(ctx context.Context) error {
	for topic := range subMapInfo.subTopics {
		if err := mqtt.Subscribe(ctx, topic, subMapInfo.HandleMessage(ctx, subMapInfo.subTopics[topic])); err != nil {
			return err
		}
	}
	return nil
}

func HandleMessage(ctx context.Context, message *amqp.Message) {
	topic := message.ApplicationProperties["topic"].(string)
	if f, ok := subMapInfo.subTopics[topic]; ok {
		fmt.Println("------待处理的消息data received------:", string(message.GetData()), " properties:", message.ApplicationProperties)
		err := f.f(ctx, topicModel.TopicHandlerData{Topic: topic, PayLoad: message.GetData(), DeviceDetail: nil})
		if err != nil {
			return
		}
	} else {
		fmt.Println("--------未处理的data received-------:", string(message.GetData()), " properties:", message.ApplicationProperties)
	}
}
