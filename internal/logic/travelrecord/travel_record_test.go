package travelrecord

import (
	"context"
	"github.com/gogf/gf/v2/test/gtest"
	_ "sviwo/internal/cmd"
	_ "sviwo/internal/logic/product"
	"sviwo/internal/model"
	"sviwo/internal/service"
	"testing"
)

func TestCreateOnline(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		//cache.SetAdapter(context.Background())
		data := model.TravelRecordOnline{
			DeviceName: "asdas546a4s6d5",
		}
		err := service.TravelRecord().CreateOnline(context.Background(), data)
		if err != nil {
			t.Log(err)
		}
		t.Log(data)
	})

}
