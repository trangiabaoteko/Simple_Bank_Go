package controllers

import (
	"strconv"

	"gorm.io/gorm"
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
	// get sender's id
	sender_id, _ := strconv.Atoi(c.Param("sender_id"))
	receiver_id, _ := strconv.Atoi(c.Param("receiver_id"))

	// get req.body
	var body struct {
		Amount   int
		Currency string
		Message	 string
	}

	c.Bind(&body)

	// handle transaction
	result := TransferPerformTransaction(initializers.DB, sender_id, receiver_id, body.Amount, body.Currency, body.Message)
	if result != nil {
		c.JSON(400, gin.H{
			"msg": "can't create transaction",
		})
		return
	}
	
	c.Status(200)
}





// other support functions
func TransferPerformTransaction(db *gorm.DB, sender_id int, receiver_id int, amount int, currency string, message string) error {
	// start new transaction
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}


	// init variables
	var sender, receiver models.Account

	// find sender's account
	result := tx.First(&sender, sender_id)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	// find receiver's account 
	result = tx.First(&receiver, receiver_id)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	// check currency
	if receiver.Currency != currency {
		tx.Rollback()
		return result.Error
	}

	// check sender's balance
	if sender.Balance - 50000 < amount {
		tx.Rollback()
		return result.Error
	}

	// update balance of 2 accounts
	if err := tx.Model(&sender).Update("balance", sender.Balance - amount).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(&receiver).Update("balance", receiver.Balance + amount).Error; err != nil {
		tx.Rollback()
		return err
	}

	// create new transfer and save to db
	transfer := models.Transfer{Sender: sender_id, Receiver: receiver_id, Amount: amount, Message: message}
	if err := tx.Create(&transfer).Error; err != nil {
		tx.Rollback()
		return err
	}


	// commit transaction
	tx.Commit()
	return nil
}