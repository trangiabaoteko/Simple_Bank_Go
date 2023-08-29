package main

import (
	"github.com/gin-gonic/gin"
	"github.com/trangiabaoteko/simplebank/controllers"
	"github.com/trangiabaoteko/simplebank/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()

	r.GET("/accounts/", controllers.AccountGetAll)
	r.GET("/accounts/:id", controllers.AccountGetById)
	r.POST("/accounts/", controllers.AccountCreate)
	r.PUT("/accounts/:id", controllers.AccountUpdate)
	r.DELETE("/accounts/:id", controllers.AccountDelete)

	r.Run()
}
