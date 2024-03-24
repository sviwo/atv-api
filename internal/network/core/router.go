package core

import (
	"context"
	"errors"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/gogf/gf/v2/frame/g"
	"pack.ag/amqp"
	"regexp"
	"strings"
	"sviwo/internal/consts"
	"sviwo/internal/model"
	"sviwo/internal/mqtt"
	"sviwo/internal/network/core/logic/baseLogic"
	"sviwo/pkg/dcache"
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
	gPool.Go(func(ctx context.Context) error {
		topic := message.ApplicationProperties["topic"].(string)
		topicInfo := strings.Split(topic, "/")
		if len(topicInfo) < 3 {
			//todo 是否入库，前端展示
			return errors.New(fmt.Sprintf("topic:%s is illegal, message(%s) ignored", topic, string(message.GetData())))
		}
		//解析消息详情
		productKey, deviceKey := topicInfo[1], topicInfo[2]
		if topicInfo[1] == "as" {
			productKey, deviceKey = topicInfo[4], topicInfo[5]
		}

		// 获取设备详情，拿出来消息协议，然后按照产品定义的消息协议解析消息
		deviceInfo, err := dcache.GetDeviceDetailInfo(deviceKey)
		if err != nil {
			g.Log().Debugf(ctx, "device info error: %v, topic:%s, message:%s, message ignored", err.Error(), topic, string(message.GetData()))
			return nil
		}
		if deviceInfo == nil {
			g.Log().Debugf(ctx, "device info is nil, topic:%s, message:%s, message ignored", topic, string(message.GetData()))
			return nil
		}
		// 设备禁用不处理
		if deviceInfo.Status == model.DeviceStatusNoEnable {
			g.Log().Debug(ctx, deviceKey, "device is no enable")
			return nil
		}
		dcache.UpdateStatus(ctx, deviceInfo) //更新设备状态

		for k, handleF := range subMapInfo.subTopics {
			k = transTopic(k)
			isMatch, _ := regexp.MatchString(k, topic)
			// 匹配输入字符串
			if isMatch {
				//记录TD设备日志
				logType := handleF.logType
				if len(topicInfo) == 7 && topicInfo[5] == "property" {
					logType = consts.MsgTypePropertyReport
				}
				go baseLogic.InertTdLog(ctx, logType, deviceKey, string(message.GetData()))
				if err := handleF.f(ctx, topicModel.TopicHandlerData{
					Topic:        topic,
					ProductKey:   productKey,
					DeviceKey:    deviceKey,
					PayLoad:      message.GetData(),
					DeviceDetail: deviceInfo,
				}); err != nil {
					if err.Error() != "ignore" {
						return errors.New(fmt.Sprintf("handleF error: %s, topic:%s, message:%s ", err.Error(), topic, string(message.GetData())))

					} else {
						g.Log().Infof(ctx, "handleF: %s, topic:%s, message:%s, message ignored", err.Error(), topic, string(message.GetData()))
						return nil
					}
				}
				break
			}
		}
		return nil
	})

	//if f, ok := subMapInfo.subTopics[topic]; ok {
	//	fmt.Println("------待处理的消息data received------:", string(message.GetData()), " properties:", message.ApplicationProperties)
	//	err := f.f(ctx, topicModel.TopicHandlerData{Topic: topic, PayLoad: message.GetData(), DeviceDetail: nil})
	//	if err != nil {
	//		return
	//	}
	//} else {
	//	fmt.Println("--------未处理的data received-------:", string(message.GetData()), " properties:", message.ApplicationProperties)
	//}
}

func transTopic(topic string) string {
	topic = strings.ReplaceAll(topic, "+", ".+")
	return strings.ReplaceAll(topic, "#", "*")
}
