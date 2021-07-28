package contract

import "github.com/kaixinhupo/quick/model"

type UserService interface {
	CreateUser(user *model.UserInfoReq) (*model.UserDetailResp, error)
}