package main

import (
	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"
)

func configSwagger(app *iris.Application) {
	config := swagger.Config{
		URL:         "http://localhost:8080/swagger/doc.json",
		DeepLinking: true,
	}
	swaggerUI := swagger.CustomWrapHandler(&config, swaggerFiles.Handler)

	//app.Get("/swagger", swaggerUI)
	app.Get("/swagger/{any:path}", swaggerUI)
}