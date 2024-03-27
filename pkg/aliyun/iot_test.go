package aliyun

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/grand"
	"testing"
)

func TestSetDevicePropertyRequest(t *testing.T) {
	maps := make(map[string]interface{})
	maps["VehSpeed"] = 88.00
	marshal, err := json.Marshal(maps)
	if err != nil {
		return
	}
	err = SetDevicePropertyRequest(context.Background(), "k0ugjmf1ois", "sviwo_atv", string(marshal))
	if err != nil {
		return
	}
}

//func TestInitAliyunIotClient(t *testing.T) {
//	tests := []struct {
//		name    string
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if err := InitAliyunIotClient(); (err != nil) != tt.wantErr {
//				t.Errorf("InitAliyunIotClient() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}

func TestRegisterDevice(t *testing.T) {
	ctx := context.Background()
	data, err := RegisterDevice(ctx, "k0ugjmf1ois", "sviwo-"+grand.S(10))
	if err != nil {
		panic(err)
	}
	glog.Info(ctx, data)
}
