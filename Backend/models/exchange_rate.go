package models

import "time"

type ExchangeRates struct {
	ID           uint      `gorm:"primary_key" json:"_id"`
	FromCurrency string    `json:"fromCurrency" binding:"required"`
	ToCurrency   string    `json:"toCurrency" binding:"required"`
	Rate         float64   `json:"rate" binding:"required"`
	Date         time.Time `json:"date"`
}
