package product

import (
	"context"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"sviwo/internal/model"
	"sviwo/internal/service"
	"sviwo/pkg/aliyun"
	"sviwo/pkg/cache"
	"testing"
)

func TestSet(t *testing.T) {
	aliyun.InitAliyunIotClient(context.Background())
	cache.SetAdapter(context.Background())
	in := &model.DevicePropertyInput{
		DeviceKey: "sviwo_atv",
		Params: map[string]any{
			"speaker": 0,
			"Light":   0,
		},
	}
	out, err := service.DevDeviceProperty().Set(context.TODO(), in)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(out)
}
