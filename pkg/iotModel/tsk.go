package iotModel

import "github.com/gogf/gf/v2/os/gtime"

type TsdTables struct {
	TableName  string      `json:"tableName"        dc:"表名"`
	DbName     string      `json:"dbName"        dc:"数据库名"`
	StableName string      `json:"stableName"        dc:"超级表名"`
	CreateTime *gtime.Time `json:"createTime" dc:"创建时间"`
}

type TsdTableInfo struct {
	Field  string `json:"field"        dc:"字段名"`
	Type   string `json:"type"        dc:"类型"`
	Length int    `json:"length"        dc:"长度"`
	Note   string `json:"note" dc:"note"`
}

type TsdTableDataInfo struct {
	Filed []string                 `json:"filed"        dc:"字段"`
	Info  []map[string]interface{} `json:"info"        dc:"数据"`
}
