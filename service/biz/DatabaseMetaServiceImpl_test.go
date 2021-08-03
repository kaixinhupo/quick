package biz

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/iancoleman/strcase"
	"github.com/kaixinhupo/quick/infrastruture/config"
	"github.com/kaixinhupo/quick/infrastruture/db"
)

func setupDatabaseMetaService() *DatabaseMetaServiceImpl {
	xormConfig := config.XormConfig {
		DatasourceType: "mysql",
		DatasourceName: "quick:Quick!0729@tcp(123.57.235.86:13307)/quick?charset=utf8mb4&collation=utf8mb4_unicode_ci",
	}
	genConfig := &config.GenConfig {
		Module: "github.com/kaixinhupo/quick",
        TemplateDir:"D:\\Go\\home\\src\\github.com\\kaixinhupo\\quick\\templates",
        OutputDir:"D:\\Go\\home\\src\\github.com\\kaixinhupo\\quick\\output",
		CreateTimeCol: "created_at",
		UpdateTimeCol: "updated_at",
	}
	engine := db.DefaultEngine(&xormConfig)
	return NewDatabaseMetaService(engine, genConfig)
}

func TestReadMeta(t *testing.T) {
	svc := setupDatabaseMetaService()
	tables := make([]string, 1)
	tables[0] = "user"
	meta, err := svc.ReadMeta(tables)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	json, _ := json.Marshal(meta)
	fmt.Println(string(json))
}

func TestToCamel(t *testing.T) {
	s := "tbl_user"
	s = strcase.ToCamel(s)
	fmt.Println(s)
	s = "tb-user"
	s = strcase.ToCamel(s)
	fmt.Println(s)
}
