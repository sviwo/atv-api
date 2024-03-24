package analysis

import (
	"context"
	_ "github.com/taosdata/driver-go/v3/taosRestful"
	_ "github.com/taosdata/driver-go/v3/taosWS"
	"sviwo/internal/service"
	"sviwo/pkg/general"

	"testing"
)

func TestGetDeviceDataByTsd(t *testing.T) {
	var reqData general.SelectReq
	reqData.Param = make(map[string]interface{})
	reqData.Param["deviceKey"] = "t202200001"
	gotNumber, err := service.AnalysisDeviceDataTsd().GetDeviceData(context.Background(), reqData)
	if err != nil {
		t.Log(err)
	}
	t.Log(gotNumber)

}
