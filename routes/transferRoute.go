package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/trangiabaoteko/simplebank/controllers"
)

func TransferRoutes(transfer *gin.RouterGroup) {
	transfer.GET("/", controllers.TransferGetAll)
	transfer.PUT("/:sender_id/:receiver_id", controllers.TransferTransact)
}
