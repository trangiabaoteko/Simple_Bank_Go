package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/trangiabaoteko/simplebank/controllers"
)

func AccountRoutes(accounts *gin.RouterGroup) {
	accounts.GET("/", controllers.AccountGetAll)
	accounts.GET("/:id", controllers.AccountGetById)
	accounts.POST("/", controllers.AccountCreate)
	// accounts.PUT("/:id", controllers.AccountUpdate)
	accounts.DELETE("/:id", controllers.AccountDelete)
}
