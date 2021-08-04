package db

import (
	"log"

	"xorm.io/xorm"
	l "xorm.io/xorm/log"
	"xorm.io/xorm/names"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kaixinhupo/quick/infrastructure/config"
)

var _defaultEngine *xorm.Engine

func DefaultEngine(config *config.XormConfig) *xorm.Engine {
	if _defaultEngine == nil {
		engine, err := xorm.NewEngine(config.DatasourceType, config.DatasourceName)
		if err != nil {
			log.Fatalln("connect db fail ", err)
			return nil
		}
		engine.ShowSQL(true)
		engine.Logger().SetLevel(l.LOG_DEBUG)
		engine.SetMapper(names.GonicMapper{})
		err = engine.Ping()
		if err != nil {
			log.Fatalln("connect db fail:", err.Error())
			return nil
		}
		_defaultEngine = engine
	}

	return _defaultEngine
}
