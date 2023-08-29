package models

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	ID       int
	Owner    string
	Balance  int
	Currency string
}
