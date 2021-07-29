package config

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

func DatasourceConfig() *XormConfig {
	return &Config.Xorm
}

var Config, _ = NewAppConfig()

func NewAppConfig() (*AppConfig, error) {
	confPath := os.Getenv("CONF_PATH")
	if confPath == "" {
		dir := filepath.Dir(os.Args[0])
		toml := filepath.Join(dir,"config.toml")
		confPath, _ = filepath.Abs(toml)
	}
	if confPath == "" {
		return nil, errors.New("not found application config file")
	}
	log.Println("config file path:", confPath)
	confBytes, err := ioutil.ReadFile(confPath) 
	if err != nil {
		return nil, errors.New("read config file error：" + err.Error())
	}

	appConfig := &AppConfig{}

	if _, err := toml.Decode(string(confBytes), appConfig); err != nil {
		log.Println("decode config file error ",err.Error())
		return nil, err
	}

	return appConfig, nil
}

type AppConfig struct {
	Xorm XormConfig `toml:"xorm"`
}

type XormConfig struct {
	// type of datasource,eg sqlite3,mysql
	DatasourceType string `toml:"datasourceType"`
	// connection string of the datasource
	DatasourceName string `toml:"datasourceName"`
}