package controller

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/kaixinhupo/quick/infrastruture/config"
	"github.com/kaixinhupo/quick/infrastruture/core"
	"github.com/kaixinhupo/quick/model"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/view"
)

type GenController struct {
    engine *view.HandlebarsEngine
    config *config.GenConfig
}

func (c *GenController) GetAsync(ctx iris.Context) mvc.Result {
    ctxCopy := ctx.Clone()
    go func (ctx iris.Context)  {
        time.Sleep(5 * time.Second)
        log.Printf("Done! in path: %s", ctx.Path())
    } (ctxCopy)

    return mvc.Response{
        Text: "ok",
    }
}

func (c *GenController) GetTables(ctx iris.Context) mvc.Result {
    uid := uuid.NewString()
    output := c.initOutputDir(uid)
    _, err := os.Stat(output); if err!=nil {
        if os.IsNotExist(err) {
            os.MkdirAll(output,os.ModePerm)
        }
    } 
    meta := c.prepareMeta()
    var e error = nil
    for _, m := range meta {
        e = c.generate(m,output); if e!=nil {
            break
        }
    }
    if e !=nil {
        return mvc.Response{
            Text: "fail",
            Err: e,
        }
    }
    err = c.zipDir(output,uid); if err !=nil {
        return mvc.Response{
            Text: "fail",
            Err: err,
        }
    }
    os.RemoveAll(output)
    return mvc.Response{
        Text: "ok",
    }
}


func (c *GenController) GetInfo(ctx iris.Context) mvc.Result {
    var buff bytes.Buffer
    var context = make(map[string]interface{})
    context["ModelName"] = "User"
    context["ModelNameLower"] = "user"

	err := c.engine.ExecuteWriter(&buff, "info.txt.tpl", "", context); if err != nil {
        return mvc.Response {
            Err: err,
        }
    }

    return mvc.Response{
        Text: buff.String(),
    }
}

func (c *GenController) GetEntity(ctx iris.Context) mvc.Result {
    var buff bytes.Buffer
    var context = c.prepareMeta()[0]
	err := c.engine.ExecuteWriter(&buff, "entity.go.tpl", "", context); if err != nil {
        return mvc.Response {
            Err: err,
        }
    }
    return mvc.Response{
        Text: buff.String(),
    }
}

func (c *GenController)  zipDir(output string,uid string) error {
    zip := filepath.Join(c.config.OutputDir,uid+".zip")
    return core.Compress(output,zip)
}

// 生成文件
func (c *GenController) generate(meta *model.TableMeta,output string) error {
    err := c.generateInfo(meta,output) ;if err != nil {
        return  err
    }
    err = c.generateEntity(meta,output) ;if err != nil {
        return  err
    }

    return nil
}

// 生成info文件
func (c *GenController) generateInfo(meta *model.TableMeta,output string) error {
    path := filepath.Join(output,"info.txt")
    return c.generateByTemplate(path,meta,"info.txt.tpl")
}

func (c *GenController) generateEntity(meta *model.TableMeta,output string) error {
    dir := filepath.Join(output,"dao","entity")
    err := os.MkdirAll(dir,os.ModeDir); if err != nil {
        return err
    }
    path := filepath.Join(dir,meta.ModelName+ "Entity.go")
    return c.generateByTemplate(path,meta,"entity.go.tpl")
}

// 生成info文件
func (c *GenController) generateByTemplate(path string, meta *model.TableMeta,template string) error {
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

// 准备元数据
func (c *GenController) prepareMeta() []*model.TableMeta {
    var fields = make([]*model.FieldMeta,5)
    fields[0] = & model.FieldMeta{
        Col:"id",                 
        ColType:"bigint",           
        ColLen:0,               
        ColComment:"ID",            
        ColNull:false ,           
        Property:"Id",          
        PropertyType:"int64",   
        ColPk: true,     
    }
    fields[1] = & model.FieldMeta{
        Col:"username",                 
        ColType:"varchar",           
        ColLen:50,               
        ColComment:"用户名",            
        ColNull:false ,           
        Property:"Username",          
        PropertyType:"string",   
        ColPk: false,          
    }
    fields[2] = & model.FieldMeta{
        Col:"password",                 
        ColType:"varchar",           
        ColLen:255,               
        ColComment:"密码",            
        ColNull:false ,           
        Property:"Password",          
        PropertyType:"string",  
        ColPk: false,              
    }
    fields[3] = & model.FieldMeta{
        Col:"created_at",                 
        ColType:"datetime",           
        ColLen:0,               
        ColComment:"创建时间",            
        ColNull:false ,           
        Property:"CreatedAt",          
        PropertyType:"time.Time",      
        ColPk: false,          
    }
    fields[4] = & model.FieldMeta{
        Col:"updated_at",                 
        ColType:"datetime",           
        ColLen:0,               
        ColComment:"更新时间",            
        ColNull:false ,           
        Property:"UpdatedAt",          
        PropertyType:"time.Time",   
        ColPk: false,             
    }
    var meta = make([]*model.TableMeta,1)
    meta[0] =& model.TableMeta{
        TableName: "user",
        TableComment: "用户",
        ModelName: "User",                 
        ModelNameLower: "user",    
        Fields: fields,          
    }
    return meta
}

// 初始化输出目录
func (c *GenController) initOutputDir(uid string) string {
    output := c.config.OutputDir
    log.Println("output dir:",output)
    return filepath.Join(output,uid)
}

// 返回路由根路径
func (c GenController) Route() string {
	return "/gen"
}

// 构造器
func NewGenController(config *config.GenConfig) *GenController {
    log.Println("templates dir:",config.TemplateDir)
    _engine := iris.Handlebars(config.TemplateDir,".tpl")
    err := _engine.Load(); if err != nil {
        log.Println("error load template:",err.Error())
        return nil
    }

	return &GenController {
        engine : _engine,
        config: config,
	}
}