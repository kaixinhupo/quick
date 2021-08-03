package biz

import (
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/kaixinhupo/quick/infrastruture/config"
	"github.com/kaixinhupo/quick/infrastruture/core"
	"github.com/kaixinhupo/quick/infrastruture/errors"
	"github.com/kaixinhupo/quick/model"
	"xorm.io/xorm"
	"xorm.io/xorm/schemas"
)

type DatabaseMetaServiceImpl struct {
    engine *xorm.Engine
    config *config.GenConfig
}



func (impl *DatabaseMetaServiceImpl) ReadMeta(tables []string) ([]*model.TableMeta,error) {
    if len(tables)==0 {
        return nil,errors.NewBizError(core.CodeInvalidParam,"tables参数不能为空")
    }
    tbls,err := impl.engine.DBMetas(); if err != nil {
        return nil, err
    }
    rst := make([]*model.TableMeta,0) 
    if len(tbls)>0 {
       dict := make(map[string]int)
       for _, k := range tables {
           dict[k]=1
       }
       inx := 0
       for _, tbl := range tbls {
           if dict[tbl.Name] ==1 {
            m := &model.TableMeta {
                TableName: tbl.Name,
                TableComment: getComment(tbl.Comment,tbl.Name) ,
                ModelName: getModelName(tbl.Name),
                Fields: impl.getFields(tbl),
            }
            generateQueries(m)
            rst = append(rst,m)
            inx ++;
           }
       }
    }
    return rst,nil
}

func (impl *DatabaseMetaServiceImpl) getFields(table *schemas.Table) []*model.FieldMeta {
    cols := table.Columns()
    rst := make([]*model.FieldMeta,len(cols))
    for i, col := range cols {
        field := &model.FieldMeta {
            Col: col.Name,
            ColType: col.SQLType.Name,
            ColLen: col.SQLType.DefaultLength,
            ColComment: col.Comment,
            ColNull: col.Nullable,
            ColPk: col.IsPrimaryKey,
            ColCreate: impl.isCreate(col.Name),
            ColUpdate: impl.isUpdate(col.Name),
            ColSystem: false,
            Property: getProperty(col.Name),
            PropertyType: getPropertyType(col.SQLType.Name),
            PropertyCamel: getPropertyCamel(col.Name),
            PropertyNullValue: getPropertyNullValue(col.SQLType.Name),
       }
       updateColSystem(field)
       rst[i] = field
    }
    return rst
}

func (impl *DatabaseMetaServiceImpl)isCreate(col string) bool {
    return impl.config.CreateTimeCol == col
}

func (impl *DatabaseMetaServiceImpl)isUpdate(col string) bool {
    return impl.config.UpdateTimeCol == col
}

func updateColSystem(field *model.FieldMeta) {
    if field.ColPk || field.ColCreate || field.ColUpdate {
        field.ColSystem = true
    }
}

func generateQueries(table *model.TableMeta) {
    queries := make([]*model.FieldMeta,0) 
    for _, f := range table.Fields {
        if !f.ColSystem && f.PropertyType != "core.LocalTime" {
            queries =  append(queries,f)
        }
    }
    if len(queries) >0 {
        queries[0].First = true
    }
    table.Queries = queries
}

func getPropertyCamel(colName string) string {
    return strcase.ToLowerCamel(colName)
}

func getPropertyType(colType string) string {
    colType = strings.ReplaceAll(colType,"UNSIGNED ","")
    t := typeMap[colType]
    if t == "" {
        t ="[]byte"
    }
    return t
}

func getPropertyNullValue(colType string) string {
    colType = strings.ReplaceAll(colType,"UNSIGNED ","")
    t := nullValueMap[colType]
    if t == "" {
        t ="nil"
    }
    return t
}


func getProperty(colName string) string {
    return strcase.ToCamel(colName)
}

func getModelName(tableName string) string {
    return strcase.ToCamel(tableName)
}

func getComment(comment string,defaultVal string) string {
    return optional(comment,defaultVal)
}

func optional(val string,els string) string {
    if val == "" {
        return els
    }
    return val 
}

func NewDatabaseMetaService(engine *xorm.Engine, config *config.GenConfig) *DatabaseMetaServiceImpl {
    return  &DatabaseMetaServiceImpl {
        engine: engine,
        config: config,
	}
}

var typeMap = map[string]string {
    "BIT":"int8",
    "TINYINT":"int8",
    "SMALLINT":"int16",
    "MEDIUMINT":"int32",
    "INT":"int",
    "INTEGER":"int",
    "BIGINT":"int64",
    "CHAR":"string",
    "VARCHAR":"string",
    "TINYTEXT":"string",
    "TEXT":"string",
    "MEDIUMTEXT":"string",
    "LONGTEXT":"string",
    "BINARY":"[]byte",
    "VARBINARY":"[]byte",
    "DATE":"core.LocalTime",
    "DATETIME":"core.LocalTime",
    "TIME":"string",
    "TIMESTAMP":"core.LocalTime",
    "REAL":"float64",
    "FLOAT":"float32",
    "DOUBLE":"float64",
    "DECIMAL":"float64",
    "NUMERIC":"float64",
    "TINYBLOB":"[]byte",
    "BLOB":"[]byte",
    "MEDIUMBLOB":"[]byte",
    "LONGBLOB":"[]byte",
}

var nullValueMap = map[string]string {
    "BIT":"0",
    "TINYINT":"0",
    "SMALLINT":"0",
    "MEDIUMINT":"0",
    "INT":"0",
    "INTEGER":"0",
    "BIGINT":"0",
    "CHAR":`""`,
    "VARCHAR":`""`,
    "TINYTEXT":`""`,
    "TEXT":`""`,
    "MEDIUMTEXT":`""`,
    "LONGTEXT":`""`,
    "TIME":`""`,
    "REAL":"0",
    "FLOAT":"0",
    "DOUBLE":"0",
    "DECIMAL":"0",
    "NUMERIC":"0",
}