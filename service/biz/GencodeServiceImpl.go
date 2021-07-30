package biz

import (
	"bufio"
	"log"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/kaixinhupo/quick/infrastruture/config"
	"github.com/kaixinhupo/quick/infrastruture/core"
	"github.com/kaixinhupo/quick/model"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/view"
)


type GencodeServiceImpl struct {
    config *config.GenConfig
    engine *view.HandlebarsEngine
}

func (impl *GencodeServiceImpl) Generate(meta []*model.TableMeta) error {
    uid := uuid.NewString()
    output := impl.initOutputDir(uid)
    _, err := os.Stat(output); if err!=nil {
        if os.IsNotExist(err) {
            os.MkdirAll(output,os.ModePerm)
        }
    } 
    for _, m := range meta {
        err := impl.generate(m,output);if err!=nil {
            return err
        }
    }
    err = impl.zipDir(output,uid); if err !=nil {
       return err
    }
    os.RemoveAll(output)
    return nil
}

// 初始化输出目录
func (c *GencodeServiceImpl) initOutputDir(uid string) string {
    output := c.config.OutputDir
    log.Println("output dir:",output)
    return filepath.Join(output,uid)
}

func (c *GencodeServiceImpl) zipDir(output string,uid string) error {
    zip := filepath.Join(c.config.OutputDir,uid+".zip")
    return core.Compress(output,zip)
}

// 生成文件
func (c *GencodeServiceImpl) generate(meta *model.TableMeta,output string) error {
    err := c.generateInfo(meta,output) ;if err != nil {
        return  err
    }
    err = c.generateEntity(meta,output) ;if err != nil {
        return  err
    }
    return nil
}

// 生成info文件
func (c *GencodeServiceImpl) generateInfo(meta *model.TableMeta,output string) error {
    path := filepath.Join(output,"info.txt")
    return c.generateByTemplate(path,meta,"info.txt.tpl")
}

func (c *GencodeServiceImpl) generateEntity(meta *model.TableMeta,output string) error {
    dir := filepath.Join(output,"dao","entity")
    err := os.MkdirAll(dir,os.ModeDir); if err != nil {
        return err
    }
    path := filepath.Join(dir,meta.ModelName+ "Entity.go")
    return c.generateByTemplate(path,meta,"entity.go.tpl")
}

// 生成info文件
func (c *GencodeServiceImpl) generateByTemplate(path string, meta *model.TableMeta,template string) error {
    f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0777) ;if err != nil {
        return err
    }
    defer f.Close()
    buf := bufio.NewWriter(f)
    err = c.engine.ExecuteWriter(buf,template,"",meta);if err != nil {
        return err
    }
    return buf.Flush()
}

func NewGencodeService(config *config.GenConfig) *GencodeServiceImpl {
    _engine := iris.Handlebars(config.TemplateDir,".tpl")
    err := _engine.Load(); if err != nil {
        log.Println("error load template:",err.Error())
        return nil
    }
    return &GencodeServiceImpl {
        engine : _engine,
        config: config,
	}
}