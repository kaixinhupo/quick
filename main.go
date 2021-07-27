package main

import (
	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"

	"github.com/kaixinhupo/quick/controller"
	_ "github.com/kaixinhupo/quick/docs"
)

// @title Swagger Quick Admin API
// @version 1.0
// @description a golang backend api for admin project.
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	app := iris.New()
	controller.Configure(app)
	configSwagger(app)
	app.Run(iris.Addr(":8080"))
}

func configSwagger(app *iris.Application) {
	config := swagger.Config{
        URL:          "http://localhost:8080/swagger/doc.json",
        DeepLinking:  true,
    }
    swaggerUI := swagger.CustomWrapHandler(&config, swaggerFiles.Handler )

    //app.Get("/swagger", swaggerUI)
    app.Get("/swagger/{any:path}", swaggerUI)
}