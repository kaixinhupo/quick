package model

import (
	"github.com/kaixinhupo/quick/infrastruture/core"
	"github.com/kaixinhupo/quick/infrastruture/web"
)

// 创建请求参数
type UserInfoReq struct {
	Username string `json:"username"` //用户名
	Password string `json:"password"` //密码
}

// 查询请求参数
type UserQueryReq struct {
	Page web.PageParam `json:"page"` // 分页
	Username  string `json:"username"` // 用户名
}

// 详情响应
type UserDetailResp struct {
	Id        int64  `json:"id"` //ID
	Username  string `json:"username"` //用户名
	Password  string `json:"password"` //密码
	CreatedAt core.LocalTime `json:"createdAt"` //创建时间
	UpdatedAt core.LocalTime `json:"updatedAt"` //修改时间
}