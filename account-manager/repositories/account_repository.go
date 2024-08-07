package repositories

import (
    "be-test-concrete-ai/account-manager/config"
    "be-test-concrete-ai/account-manager/models"
)

func GetAccountsByUserID(userID uint) ([]models.Account, error) {
    var accounts []models.Account
    if err := config.DB.Where("user_id = ?", userID).Find(&accounts).Error; err != nil {
        return nil, err
    }
    return accounts, nil
}

func GetTransactionsByAccountID(accountID string) ([]models.PaymentHistory, error) {
    var transactions []models.PaymentHistory
    if err := config.DB.Where("account_id = ?", accountID).Find(&transactions).Error; err != nil {
        return nil, err
    }
    return transactions, nil
}

func CreateAccount(account *models.Account) error {
	return config.DB.Create(account).Error
}