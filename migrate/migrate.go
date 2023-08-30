package main

import (
	"github.com/trangiabaoteko/simplebank/initializers"
	"github.com/trangiabaoteko/simplebank/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Account{})
	initializers.DB.AutoMigrate(&models.Transfer{})
}
