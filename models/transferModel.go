package models

import (
	"gorm.io/gorm"
)

type Transfer struct {
	gorm.Model
	ID       int
	Sender   int
	Receiver int
	Amount   int
	Message  string
}
