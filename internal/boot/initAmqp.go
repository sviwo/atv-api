package boot

import (
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"sviwo/pkg/amqp"
	"time"
)

func InitAmqp(ctx context.Context) error {
	go func() {
		accessKey := g.Cfg().MustGet(ctx, "aliyun.accessKeyID").String()
		accessSecret := g.Cfg().MustGet(ctx, "aliyun.accessKeySecret").String()
		host := g.Cfg().MustGet(ctx, "aliyun.iot.amqp.host").String()
		clientId := g.Cfg().MustGet(ctx, "aliyun.iot.clientId").String()
		iotInstanceId := g.Cfg().MustGet(ctx, "aliyun.iot.iotInstanceId").String()
		consumerGroupId := g.Cfg().MustGet(ctx, "aliyun.iot.amqp.consumerGroupId").String()
		address := "amqps://" + host + ":5671"
		timestamp := time.Now().Nanosecond() / 1000000
		//userName组装方法，请参见AMQP客户端接入说明文档。
		userName := fmt.Sprintf("%s|authMode=aksign,signMethod=Hmacsha1,consumerGroupId=%s,authId=%s,iotInstanceId=%s,timestamp=%d|",
			clientId, consumerGroupId, accessKey, iotInstanceId, timestamp)
		stringToSign := fmt.Sprintf("authId=%s&timestamp=%d", accessKey, timestamp)
		hmacKey := hmac.New(sha1.New, []byte(accessSecret))
		hmacKey.Write([]byte(stringToSign))
		//计算签名，password组装方法，请参见AMQP客户端接入说明文档。
		password := base64.StdEncoding.EncodeToString(hmacKey.Sum(nil))

		amqpManager := &amqp.AmqpManager{
			Address:  address,
			UserName: userName,
			Password: password,
		}

		//如果需要做接受消息通信或者取消操作，从Background衍生context。
		ctx = context.Background()
		amqpManager.StartReceiveMessage(ctx)
	}()
	return nil
}
