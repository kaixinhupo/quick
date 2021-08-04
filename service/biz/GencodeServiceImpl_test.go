package biz

import (
	"testing"

	"github.com/kaixinhupo/quick/infrastructure/config"
	"github.com/kaixinhupo/quick/model"
)

var _service = setupService()

func setupService() *GencodeServiceImpl {
	gc := &config.GenConfig{
		Module:      "github.com/kaixinhupo/quick",
		TemplateDir: "D:\\Go\\home\\src\\github.com\\kaixinhupo\\quick\\templates",
		OutputDir:   "D:\\Go\\home\\src\\github.com\\kaixinhupo\\quick\\output",
	}
	return NewGencodeService(gc)
}

func prepareMeta() []*model.TableMeta {
	var fields = make([]*model.FieldMeta, 5)
	fields[0] = &model.FieldMeta{
		Col:           "id",
		ColType:       "bigint",
		ColComment:    "ID",
		ColNull:       false,
		Property:      "Id",
		ColSystem:     true,
		PropertyType:  "int64",
		PropertyCamel: "id",
		ColPk:         true,
	}
	fields[1] = &model.FieldMeta{
		Col:           "username",
		ColType:       "varchar",
		ColComment:    "用户名",
		ColNull:       false,
		Property:      "Username",
		PropertyType:  "string",
		PropertyCamel: "username",
		ColPk:         false,
	}
	fields[2] = &model.FieldMeta{
		Col:           "password",
		ColType:       "varchar",
		ColComment:    "密码",
		ColNull:       false,
		Property:      "Password",
		PropertyType:  "string",
		PropertyCamel: "password",
		ColPk:         false,
	}
	fields[3] = &model.FieldMeta{
		Col:           "created_at",
		ColType:       "datetime",
		ColComment:    "创建时间",
		ColNull:       false,
		ColSystem:     true,
		ColCreate:     true,
		Property:      "CreatedAt",
		PropertyType:  "core.LocalTime",
		PropertyCamel: "createdAt",
		ColPk:         false,
	}
	fields[4] = &model.FieldMeta{
		Col:           "updated_at",
		ColType:       "datetime",
		ColComment:    "更新时间",
		ColNull:       false,
		ColUpdate:     true,
		ColSystem:     true,
		Property:      "UpdatedAt",
		PropertyCamel: "updatedAt",
		PropertyType:  "core.LocalTime",
		ColPk:         false,
	}
	var meta = make([]*model.TableMeta, 1)
	meta[0] = &model.TableMeta{
		TableName:    "user",
		TableComment: "用户",
		ModelName:    "User",
		Fields:       fields,
	}
	return meta
}

func TestGenerate(t *testing.T) {
	tm := prepareMeta()
	_service.Generate(tm)
}
