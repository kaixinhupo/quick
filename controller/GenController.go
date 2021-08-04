package controller

import (
	"log"
	"strings"
	"time"

	"github.com/kaixinhupo/quick/model"
	"github.com/kaixinhupo/quick/service/contract"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type GenController struct {
	gencodeService contract.GencodeService
	metaService    contract.DatabaseMetaService
}

func (c *GenController) GetAsync(ctx iris.Context) mvc.Result {
	ctxCopy := ctx.Clone()
	go func(ctx iris.Context) {
		time.Sleep(5 * time.Second)
		log.Printf("Done! in path: %s", ctx.Path())
	}(ctxCopy)

	return mvc.Response{
		Text: "ok",
	}
}

func (c *GenController) GetTables(ctx iris.Context) mvc.Result {

	tableQuery := ctx.URLParam("tables")
	if tableQuery == "" {
		return mvc.Response{
			Text: "query param [tables] not occours",
		}
	}
	tables := strings.Split(tableQuery, ",")

	meta, err := c.prepareMeta(tables)
	if err != nil {
		return mvc.Response{
			Text: "fail",
			Err:  err,
		}
	}
	err = c.gencodeService.Generate(meta)
	if err != nil {
		return mvc.Response{
			Text: "fail",
			Err:  err,
		}
	}
	return mvc.Response{
		Text: "ok",
	}
}

// 准备元数据
func (c *GenController) prepareMeta(tables []string) ([]*model.TableMeta, error) {
	return c.metaService.ReadMeta(tables)
}

// Route 返回路由根路径
func (c GenController) Route() string {
	return "/gen"
}

// NewGenController 构造器
func NewGenController(service contract.GencodeService, metaService contract.DatabaseMetaService) *GenController {
	return &GenController{
		gencodeService: service,
		metaService:    metaService,
	}
}
