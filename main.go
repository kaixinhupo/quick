package main

import (
	"github.com/kataras/iris/v12"
	"log"

	_ "github.com/kaixinhupo/quick/docs"
)

// @title Quick Admin API
// @version 1.0
// @description a golang backend api for admin project.
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	app := iris.New()
	ConfigureRouter(app)
	configSwagger(app)
	err := app.Run(iris.Addr(":8080"))
	if err != nil {
		log.Println("start fail:", err.Error())
		return
	}
}
