package services

import (
    "be-test-concrete-ai/account-manager/models"
    "be-test-concrete-ai/account-manager/repositories"
)

func GetAccounts(userID uint) ([]models.Account, error) {
    return repositories.GetAccountsByUserID(userID)
}

func GetAccountTransactions(accountID string) ([]models.PaymentHistory, error) {
    return repositories.GetTransactionsByAccountID(accountID)
}

func CreateAccount(account *models.Account) error {
	return repositories.CreateAccount(account)
}