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

type GenContext struct {
    Config *config.GenConfig
    Meta *model.TableMeta
}

func (impl *GencodeServiceImpl) Generate(meta []*model.TableMeta) error {
    uid := uuid.NewString()
    output := impl.initOutputDir(uid)
    _, err := os.Stat(output); if err!=nil {
        if os.IsNotExist(err) {
            os.MkdirAll(output,os.ModePerm)
        }
    } 
    ctx := & GenContext {
        Config: impl.config,
    }
    for _, m := range meta {
        ctx.Meta = m
        err := impl.generate(ctx,output);if err!=nil {
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
func (c *GencodeServiceImpl) generate(ctx *GenContext, output string) error {
    err := c.generateInfo(ctx,output) ;if err != nil {
        return  err
    }
    err = c.generateEntity(ctx,output) ;if err != nil {
        return  err
    }
    err = c.generateModel(ctx,output) ;if err != nil {
        return  err
    }
    err = c.generateRepository(ctx,output) ;if err != nil {
        return  err
    }
    err = c.generateRepositoryImpl(ctx,output) ;if err != nil {
        return  err
    }
    err = c.generateService(ctx,output) ;if err != nil {
        return  err
    }
    err = c.generateServiceImpl(ctx,output) ;if err != nil {
        return  err
    }
    err = c.generateController(ctx,output) ;if err != nil {
        return  err
    }
    return nil
}

// 生成info文件
func (c *GencodeServiceImpl) generateInfo(ctx *GenContext, output string) error {
    path := filepath.Join(output,"info.txt")
    return c.generateByTemplate(path,ctx,"info.txt.tpl")
}

// 生成entity文件
func (c *GencodeServiceImpl) generateEntity(ctx *GenContext, output string) error {
    meta := ctx.Meta
    dir := filepath.Join(output,"dao","entity")
    err := os.MkdirAll(dir,os.ModeDir); if err != nil {
        return err
    }
    path := filepath.Join(dir,meta.ModelName+ "Entity.go")
    return c.generateByTemplate(path,ctx,"entity.go.tpl")
}

// 生成entity文件
func (c *GencodeServiceImpl) generateModel(ctx *GenContext, output string) error {
    meta := ctx.Meta
    dir := filepath.Join(output,"model")
    err := os.MkdirAll(dir,os.ModeDir); if err != nil {
        return err
    }
    path := filepath.Join(dir,meta.ModelName+ "Model.go")
    return c.generateByTemplate(path,ctx,"model.go.tpl")
}




// 生成repository文件
func (c *GencodeServiceImpl) generateRepository(ctx *GenContext, output string) error {
    meta := ctx.Meta
    dir := filepath.Join(output,"dao","repository")
    err := os.MkdirAll(dir,os.ModeDir); if err != nil {
        return err
    }
    path := filepath.Join(dir,meta.ModelName+ "Repository.go")
    return c.generateByTemplate(path, ctx, "repository.go.tpl")
}

// 生成repository实现文件
func (c *GencodeServiceImpl) generateRepositoryImpl(ctx *GenContext, output string) error {
    meta := ctx.Meta
    dir := filepath.Join(output,"dao","database")
    err := os.MkdirAll(dir,os.ModeDir); if err != nil {
        return err
    }
    path := filepath.Join(dir,meta.ModelName+ "RepositoryImpl.go")
    return c.generateByTemplate(path, ctx, "repository-impl.go.tpl")
}

// 生成service文件
func (c *GencodeServiceImpl) generateService(ctx *GenContext, output string) error {
    meta := ctx.Meta
    dir := filepath.Join(output,"service","contract")
    err := os.MkdirAll(dir,os.ModeDir); if err != nil {
        return err
    }
    path := filepath.Join(dir,meta.ModelName+ "Service.go")
    return c.generateByTemplate(path, ctx, "service.go.tpl")
}

// 生成service实现文件
func (c *GencodeServiceImpl) generateServiceImpl(ctx *GenContext, output string) error {
    meta := ctx.Meta
    dir := filepath.Join(output,"service","biz")
    err := os.MkdirAll(dir,os.ModeDir); if err != nil {
        return err
    }
    path := filepath.Join(dir,meta.ModelName+ "ServiceImpl.go")
    return c.generateByTemplate(path, ctx, "service-impl.go.tpl")
}

// 生成controller文件
func (c *GencodeServiceImpl) generateController(ctx *GenContext, output string) error {
    meta := ctx.Meta
    dir := filepath.Join(output,"controller")
    err := os.MkdirAll(dir,os.ModeDir); if err != nil {
        return err
    }
    path := filepath.Join(dir,meta.ModelName+ "Controller.go")
    return c.generateByTemplate(path, ctx, "controller.go.tpl")
}


// 根据模板生成文件
func (c *GencodeServiceImpl) generateByTemplate(path string, ctx *GenContext, template string) error {
    f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0777) ;if err != nil {
        return err
    }
    defer f.Close()
    buf := bufio.NewWriter(f)
    err = c.engine.ExecuteWriter(buf,template,"",ctx);if err != nil {
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
    self :=  &GencodeServiceImpl {
        engine : _engine,
        config: config,
	}
    return self
}