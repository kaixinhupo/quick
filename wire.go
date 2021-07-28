//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/kaixinhupo/quick/controller"
	"github.com/kaixinhupo/quick/dao/repository"
	"github.com/kaixinhupo/quick/dao/database"
	"github.com/kaixinhupo/quick/service/biz"
	"github.com/kaixinhupo/quick/service/contract"
)

var userRepositorySet = wire.NewSet(
    database.NewUserRepository, 
    wire.Bind(new (repository.UserRepository),new (*database.UserRepositoryImpl)),
)
var userServiceSet = wire.NewSet(
    biz.NewUserService, 
    wire.Bind(new (contract.UserService),new (*biz.UserServiceImpl)),
)
func InitUserController() *controller.UserController {
	wire.Build(userRepositorySet,userServiceSet,controller.NewUserController)
	return &controller.UserController{}
}