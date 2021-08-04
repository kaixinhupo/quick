package model

// TableMeta 表元数据
type TableMeta struct {
	TableName           string       //表名
	TableComment        string       //表注释
	ModelName           string       //模型名称
	ModelNameFirstLower string       //模型名称首字母小写
	BasePath            string       //路由
	Fields              []*FieldMeta // 所有字段
	Queries             []*FieldMeta // 查询字段
}

// FieldMeta 字段元数据
type FieldMeta struct {
	Col               string //字段
	ColType           string //字段类型
	ColTypeLower      string //字段类型小写形式
	ColTypeLen        string //字段类型长度
	ColComment        string //字段注释
	ColNull           bool   //是否为空
	ColPk             bool   //是否主键
	ColSystem         bool   //ColPk、ColCreate、ColUpdate之一是否为true
	ColCreate         bool   //创建时间字段
	ColUpdate         bool   //修改时间字段
	Property          string //模型属性名称
	PropertyType      string //模型属性类型
	PropertyCamel     string //模型属性名称camel格式
	PropertyNullValue string //模型属性值为空时的值
	First             bool   //是否查询字段的第一个
}
