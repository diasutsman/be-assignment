package config

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
	"be-test-concrete-ai/account-manager/models"
)

var (
    DB *gorm.DB
)

func Init() {
    var err error
    dsn := "host=localhost user=postgres password=12345 dbname=be_test_concrete_ai port=5432 sslmode=disable"
    DB, err = gorm.Open("postgres", dsn)
    if err != nil {
        panic("failed to connect to database")
    }
    DB.AutoMigrate(&models.User{}, &models.Account{}, &models.PaymentHistory{})
}
