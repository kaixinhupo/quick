package controller

import (
	"log"
	"time"

	"github.com/kaixinhupo/quick/model"
	"github.com/kaixinhupo/quick/service/contract"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type GenController struct {
    s contract.GencodeService
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
    meta := c.prepareMeta()
    err := c.s.Generate(meta); if err !=nil {
        return mvc.Response{
            Text: "fail",
            Err: err,
        }
    }
    return mvc.Response{
        Text: "ok",
    }
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



// 返回路由根路径
func (c GenController) Route() string {
	return "/gen"
}

// 构造器
func NewGenController(service contract.GencodeService) *GenController {
	return &GenController {
        s:service,
	}
}