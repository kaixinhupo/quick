package model

type TableMeta struct {
	TableName    string
	TableComment string
	ModelName    string
	Fields       []*FieldMeta
	Queries      []*FieldMeta
}

type FieldMeta struct {
	Col               string
	ColType           string
	ColLen            int
	ColComment        string
	ColNull           bool
	ColPk             bool
	ColSystem         bool
	ColCreate         bool
	ColUpdate         bool
	Property          string
	PropertyType      string
	PropertyCamel     string
	PropertyNullValue string
	First             bool
}