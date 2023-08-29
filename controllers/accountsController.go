package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/trangiabaoteko/simplebank/initializers"
	"github.com/trangiabaoteko/simplebank/models"
)

// [GET] /
func AccountGetAll(c *gin.Context) {
	var accounts []models.Account

	initializers.DB.Find(&accounts)

	c.JSON(200, gin.H{
		"accounts": accounts,
	})
}

// [GET] /:id
func AccountGetById(c *gin.Context) {
	var account models.Account

	// get id
	id := c.Param("id")

	// find account
	result := initializers.DB.First(&account, id)

	// response
	if result.Error != nil {
		c.Status(404)
		return
	}

	c.JSON(200, gin.H{
		"account": account,
	})
}

// [POST] /
func AccountCreate(c *gin.Context) {
	// get req.body
	var body struct {
		Owner    string
		Balance  int
		Currency string
	}

	c.Bind(&body)

	// create and insert new account to database
	account := models.Account{Owner: body.Owner, Balance: body.Balance, Currency: body.Currency}

	result := initializers.DB.Create(&account)

	// response
	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"account": account,
	})
}

// [PUT] :/id
func AccountUpdate(c *gin.Context) {
	var account models.Account

	// get id
	id := c.Param("id")

	// get req.body
	var body struct {
		Owner    string
		Balance  int
		Currency string
	}

	c.Bind(&body)

	// find account
	initializers.DB.First(&account, id)

	// update account
	result := initializers.DB.Model(&account).Updates(models.Account{Owner: body.Owner, Balance: body.Balance, Currency: body.Currency})

	// response
	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"account": account,
	})
}

// [DELETE] :/id
func AccountDelete(c *gin.Context) {
	var account models.Account

	// get id
	id := c.Param("id")

	// find account
	if err := initializers.DB.First(&account, id).Error; err != nil {
		c.Status(404)
		return
	}

	// delete account
	initializers.DB.Delete(&account, id)

	c.Status(200)
}
