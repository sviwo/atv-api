package boot

import (
	"context"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	iot20180120 "github.com/alibabacloud-go/iot-20180120/v6/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/gogf/gf/v2/frame/g"
	"sviwo/pkg/aliyun"
)

func InitAliyunIotClient(ctx context.Context) error {
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
	aliyun.IotClient, err = iot20180120.NewClient(config)
	return err
}
