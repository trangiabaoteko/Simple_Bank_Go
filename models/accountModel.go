package models

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	id         int
	owner      string
	balance    int
	currency   string
	created_at time.Time
}
