// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-03-15 14:53:40
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Device is the golang structure of table sw_device for DAO operations like Where/Data.
type Device struct {
	g.Meta      `orm:"table:sw_device, do:true"`
	DeviceId    interface{} //
	ProductId   interface{} // 产品ID
	DeviceCode  interface{} // 设备编号（同于车架号）
	DeviceName  interface{} // 设备名称
	DeviceModel interface{} // 设备型号
	Nickname    interface{} // 产品昵称（目前只有ATV，则等同于车辆昵称）
	CreateTime  *gtime.Time //
	UpdateTime  *gtime.Time //
	IsDelete    interface{} // 是否删除：true=已删除，false=正常
}
