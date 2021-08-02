package model
{{#with Config}}
import (
	"{{Module}}/infrastruture/core"
	"{{Module}}/infrastruture/web"
)
{{/with}}
{{#with Meta}}
// 创建请求参数
type {{ModelName}}InfoReq struct {
	{{#each Fields}}
	{{#if ColNested}}
	{{else}}
	{{Property}} {{PropertyType}} `json:"{{PropertyCamel}}"` //{{ColComment}}
	{{/if}}
	{{/each}}
}

// 查询请求参数
type {{ModelName}}QueryReq struct {
	Page web.PageParam `json:"page"` // 分页
	{{#each Fields}}
	{{#if ColNested}}
	{{else}}
	{{Property}} {{PropertyType}} `json:"{{PropertyCamel}}"` //{{ColComment}}
	{{/if}}
	{{/each}}
}

// 详情响应
type {{ModelName}}DetailResp struct {
	{{#each Fields}}
	{{Property}} {{PropertyType}} `json:"{{PropertyCamel}}"` //ColComment
	{{/each}}
}
{{/with}}