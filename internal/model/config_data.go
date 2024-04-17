package model

import "github.com/gogf/gf/v2/os/gtime"

type ConfigDoInput struct {
	ConfigName     string `p:"configName"`     //参数名称
	ConfigKey      string `p:"configKey"`      //参数键名
	ConfigType     string `p:"configType"`     //状态
	ModuleClassify string `p:"moduleClassify"` //字典分类编码
	PaginationInput
}

type SysConfigRes struct {
	ConfigId       uint        `json:"configId"    dc:"参数主键"`
	ConfigName     string      `json:"configName"  dc:"参数名称"`
	ConfigKey      string      `json:"configKey"   dc:"参数键名"`
	ConfigValue    string      `json:"configValue" dc:"参数键值"`
	ConfigType     int         `json:"configType"  dc:"系统内置（Y是 N否）"`
	ModuleClassify string      `json:"moduleClassify" dc:"字典分类编码"`
	CreatedBy      uint        `json:"createdBy"    dc:"创建者"`
	UpdatedBy      uint        `json:"updatedBy"    dc:"更新者"`
	Remark         string      `json:"remark"      dc:"备注"`
	CreatedAt      *gtime.Time `json:"createdAt"   dc:"创建时间"`
	UpdatedAt      *gtime.Time `json:"updatedAt"   dc:"修改时间"`
}

type SysConfigOut struct {
	ConfigId       uint        `json:"configId"    dc:"参数主键"`
	ConfigName     string      `json:"configName"  dc:"参数名称"`
	ConfigKey      string      `json:"configKey"   dc:"参数键名"`
	ConfigValue    string      `json:"configValue" dc:"参数键值"`
	ConfigType     int         `json:"configType"  dc:"系统内置（Y是 N否）"`
	ModuleClassify string      `json:"moduleClassify" dc:"字典分类编码"`
	CreatedBy      uint        `json:"createdBy"    dc:"创建者"`
	UpdatedBy      uint        `json:"updatedBy"    dc:"更新者"`
	Remark         string      `json:"remark"      dc:"备注"`
	CreatedAt      *gtime.Time `json:"createdAt"   dc:"创建时间"`
	UpdatedAt      *gtime.Time `json:"updatedAt"   dc:"修改时间"`
}

type AddConfigInput struct {
	ConfigName     string `p:"configName"`
	ConfigKey      string `p:"configKey"`
	ConfigValue    string `p:"configValue"`
	ConfigType     int    `p:"configType"`
	Remark         string `p:"remark"`
	ModuleClassify string `p:"moduleClassify"`
}

type EditConfigInput struct {
	ConfigId       int    `p:"configId"`
	ConfigName     string `p:"configName"`
	ConfigKey      string `p:"configKey"`
	ConfigValue    string `p:"configValue"`
	ConfigType     int    `p:"configType"`
	Remark         string `p:"remark"`
	ModuleClassify string `p:"moduleClassify"`
}

type EditConfigReq struct {
	ConfigId    int    `p:"configId" v:"required#ID不能为空"`
	ConfigKey   string `p:"configKey" v:"required#KEY不能为空"`
	ConfigValue string `p:"configValue" v:"required#值不能为空"`
}

// CacheConfig 缓存配置
type CacheConfig struct {
	Adapter string `json:"adapter"`
	FileDir string `json:"fileDir"`
}
