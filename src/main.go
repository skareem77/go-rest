package main

import (
	"rest-api/src/app"
	"rest-api/src/config"
)

//https://github.com/mingrammer/go-todo-rest-api-example
// https://github.com/diegothucao/rest-api-golang
func main() {

	config := config.GetDBConfigurations()

	app := &app.App{}
	app.Init(config)
	app.Run(":8080")
}
