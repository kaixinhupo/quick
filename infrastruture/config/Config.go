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

func GenerateConfig() *GenConfig {
	Config.Gen.initGen()
	return &Config.Gen
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
	Gen  GenConfig  `toml:"gen"`
}

type XormConfig struct {
	// type of datasource,eg sqlite3,mysql
	DatasourceType string `toml:"datasourceType"`
	// connection string of the datasource
	DatasourceName string `toml:"datasourceName"`
}

type GenConfig struct {
	TemplateDir string `toml:"templateDir"`
	OutputDir string `toml:"outputDir"`
	isInit bool
}

func (gc *GenConfig) initGen() {
	if !gc.isInit {
		dir := filepath.Dir(os.Args[0])
		if gc.TemplateDir == "" {
			templates := filepath.Join(dir,"templates")
			gc.TemplateDir, _ = filepath.Abs(templates)
		}
		if gc.OutputDir == "" {
			output := filepath.Join(dir,"output")
			gc.OutputDir, _ = filepath.Abs(output)
		}
		gc.isInit = true
	}
}