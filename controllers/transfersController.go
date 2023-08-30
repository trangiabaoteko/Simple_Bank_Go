package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/trangiabaoteko/simplebank/initializers"
	"github.com/trangiabaoteko/simplebank/models"
)

// [GET] /
func TransferGetAll(c *gin.Context) {
	var transfers []models.Transfer

	initializers.DB.Find(&transfers)

	c.JSON(200, gin.H{
		"transfers": transfers,
	})
}

// [PUT] /:sender_id/:receiver_id
func TransferTransact(c *gin.Context) {
	var sender, receiver models.Account

	// get sender's id
	sender_id, _ := strconv.Atoi(c.Param("sender_id"))
	receiver_id, _ := strconv.Atoi(c.Param("receiver_id"))

	// find sender's account
	result := initializers.DB.First(&sender, sender_id)
	if result.Error != nil {
		c.JSON(404, gin.H{
			"msg": "sender's account not found",
		})
		return
	}

	// find receiver's account 
	result = initializers.DB.First(&receiver, receiver_id)
	if result.Error != nil {
		c.JSON(404, gin.H{
			"msg": "receiver's account not found",
		})
		return
	}

	// get req.body
	var body struct {
		Amount   int
		Currency string
		Message	 string
	}

	c.Bind(&body)

	// check currency
	if receiver.Currency != body.Currency {
		c.JSON(400, gin.H{
			"msg": "doesn't match the currency",
		})
		return
	}

	// check sender's balance
	if sender.Balance - 50000 < body.Amount {
		c.JSON(400, gin.H{
			"msg": "insufficient balance",
		})
		return
	}

	// update balance of 2 accounts
	initializers.DB.Model(&sender).Update("balance", sender.Balance - body.Amount)
	initializers.DB.Model(&receiver).Update("balance", receiver.Balance + body.Amount)

	// create new transfer and save to db
	transfer := models.Transfer{Sender: sender_id, Receiver: receiver_id, Amount: body.Amount, Message: body.Message}
	result = initializers.DB.Create(&transfer)

	// response
	if result.Error != nil {
		c.JSON(400, gin.H{
			"msg": "can't create transaction",
		})
		return
	}
	
	c.JSON(200, gin.H{
		"transfer": transfer,
	})
}