package model

import "time"

type Activity struct {
	Id          string
	Email       string `validate:"required,email"`
	Environment string `validate:"required"`
	Component   string `validate:"required"`
	Message     string `validate:"required"`
	Data        string
	CreatedAt   time.Time
}
