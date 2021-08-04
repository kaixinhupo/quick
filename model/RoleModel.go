package model

import (
	"github.com/kaixinhupo/quick/infrastructure/core"
	"github.com/kaixinhupo/quick/infrastructure/web"
)

// RoleInfoReq 创建请求参数
type RoleInfoReq struct {
	RoleName string `json:"roleName"` //角色名称

}

// RoleQueryReq 查询请求参数
type RoleQueryReq struct {
	Page     web.PageParam `json:"page"`     // 分页
	RoleName string        `json:"roleName"` //角色名称

}

// RoleDetailResp 详情响应
type RoleDetailResp struct {
	Id        int64          `json:"id"`        //ColComment
	RoleName  string         `json:"roleName"`  //ColComment
	CreatedAt core.LocalTime `json:"createdAt"` //ColComment
	UpdatedAt core.LocalTime `json:"updatedAt"` //ColComment
}
