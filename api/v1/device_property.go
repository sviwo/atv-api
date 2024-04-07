package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"sviwo/internal/model"
)

type DevicePropertyReq struct {
	g.Meta `path:"/property/set" method:"post" summary:"设备属性设置" tags:"设备"`
	*model.DevicePropertyInput
}
type DevicePropertyRes struct {
	*model.DevicePropertyOutput
}
