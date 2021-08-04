//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/kaixinhupo/quick/controller"
	"github.com/kaixinhupo/quick/dao/database"
	"github.com/kaixinhupo/quick/dao/repository"
	"github.com/kaixinhupo/quick/infrastructure/config"
	"github.com/kaixinhupo/quick/infrastructure/db"
	"github.com/kaixinhupo/quick/service/biz"
	"github.com/kaixinhupo/quick/service/contract"
)

// xorm engine
var xormEngineSet = wire.NewSet(
	config.DatasourceConfig,
	db.DefaultEngine,
)

// repository.UserRepository
var userRepositorySet = wire.NewSet(
	database.NewUserRepository,
	wire.Bind(new(repository.UserRepository), new(*database.UserRepositoryImpl)),
)

// contract.UserService
var userServiceSet = wire.NewSet(
	biz.NewUserService,
	wire.Bind(new(contract.UserService), new(*biz.UserServiceImpl)),
)

var gencodeServiceSet = wire.NewSet(
	biz.NewGencodeService,
	wire.Bind(new(contract.GencodeService), new(*biz.GencodeServiceImpl)),
)

var databaseMetaServiceSet = wire.NewSet(
	biz.NewDatabaseMetaService,
	wire.Bind(new(contract.DatabaseMetaService), new(*biz.DatabaseMetaServiceImpl)),
)

func InitUserController() *controller.UserController {
	wire.Build(xormEngineSet, userRepositorySet, userServiceSet, controller.NewUserController)
	return &controller.UserController{}
}

// repository.RoleRepository
var roleRepositorySet = wire.NewSet(
	database.NewRoleRepository,
	wire.Bind(new(repository.RoleRepository), new(*database.RoleRepositoryImpl)),
)

// contract.RoleService
var roleServiceSet = wire.NewSet(
	biz.NewRoleService,
	wire.Bind(new(contract.RoleService), new(*biz.RoleServiceImpl)),
)

func InitRoleController() *controller.RoleController {
	wire.Build(xormEngineSet, roleRepositorySet, roleServiceSet, controller.NewRoleController)
	return &controller.RoleController{}
}

func InitGenController() *controller.GenController {
	wire.Build(xormEngineSet, config.GenerateConfig, databaseMetaServiceSet, gencodeServiceSet, controller.NewGenController)
	return &controller.GenController{}
}
