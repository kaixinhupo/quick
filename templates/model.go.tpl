package model
{{#with Config}}
import (
	"{{Module}}/infrastructure/core"
	"{{Module}}/infrastructure/web"
)
{{/with}}
{{#with Meta}}
// {{ModelName}}InfoReq 创建请求参数
type {{ModelName}}InfoReq struct {
	{{#each Fields}}
	{{#unless ColSystem}}
    {{Property}} {{PropertyType}} `json:"{{PropertyCamel}}"` //{{ColComment}}
	{{/unless}}
	{{/each}}
}

// {{ModelName}}QueryReq 查询请求参数
type {{ModelName}}QueryReq struct {
    Page web.PageParam `json:"page"` // 分页
	{{#each Fields}}
	{{#unless ColSystem}}
    {{Property}} {{PropertyType}} `json:"{{PropertyCamel}}"` //{{ColComment}}
	{{/unless}}
	{{/each}}
}

// {{ModelName}}DetailResp 详情响应
type {{ModelName}}DetailResp struct {
	{{#each Fields}}
    {{Property}} {{PropertyType}} `json:"{{PropertyCamel}}"` //ColComment
	{{/each}}
}
{{/with}}