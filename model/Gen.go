package model

type TableMeta struct {
	TableName      string
	TableComment   string
	ModelName      string
	ModelNameLower string
	ModelComment   string
	Fields         []*FieldMeta
}

type FieldMeta struct {
	Col          string
	ColType      string
	ColLen       int
	ColComment   string
	ColNull      bool
	ColPk        bool
	Property     string
	PropertyType string
}