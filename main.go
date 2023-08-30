package main

import (
	"github.com/gin-gonic/gin"
	"github.com/trangiabaoteko/simplebank/initializers"
	"github.com/trangiabaoteko/simplebank/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()

	routes.AccountRoutes(r.Group("/accounts"))
	routes.TransferRoutes(r.Group("/transfer"))

	r.Run()
}
