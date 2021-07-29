package config

import "github.com/kaixinhupo/quick/infrastruture/db"

func DatasourceConfig() *db.XormConfig {
	return &db.XormConfig{
		DatasourceType: "mysql",
		DatasourceName: "quick:Quick!0729@tcp(123.57.235.86:13307)/quick?charset=utf8mb4&collation=utf8mb4_unicode_ci",
	}
}