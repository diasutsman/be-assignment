package models

import "github.com/jinzhu/gorm"

type PaymentHistory struct {
    gorm.Model
    AccountID uint `json:"account_id"`
    Amount    float64 `json:"amount"`
    Timestamp string `json:"timestamp"`
    ToAddress string `json:"to_address"`
    Status    string `json:"status"`
}
