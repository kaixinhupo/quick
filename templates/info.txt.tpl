// copy to wire.go
{{#with Meta}}
// repository.{{ModelName}}Repository
var {{ModelNameFirstLower}}RepositorySet = wire.NewSet(
  database.New{{ModelName}}Repository, 
  wire.Bind(new (repository.{{ModelName}}Repository),new (*database.{{ModelName}}RepositoryImpl)),
)

// contract.{{ModelName}}Service
var {{ModelNameFirstLower}}ServiceSet = wire.NewSet(
  biz.New{{ModelName}}Service, 
  wire.Bind(new (contract.{{ModelName}}Service),new (*biz.{{ModelName}}ServiceImpl)),
)


func Init{{ModelName}}Controller() *controller.{{ModelName}}Controller {
	wire.Build(xormEngineSet,{{ModelNameFirstLower}}RepositorySet,{{ModelNameFirstLower}}ServiceSet,controller.New{{ModelName}}Controller)
	return &controller.{{ModelName}}Controller{}
}

================================================================================================
// copy to router.go

Init{{ModelName}}Controller(),

{{/with}}