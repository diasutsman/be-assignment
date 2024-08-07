package services

import (
	"be-test-concrete-ai/account-manager/models/requests"
	"be-test-concrete-ai/account-manager/prisma/db"
	"be-test-concrete-ai/account-manager/repositories"
)

func GetAccounts(userID int) ([]db.AccountModel, error) {
    return repositories.GetAccountsByUserID(userID)
}

func GetAccountTransactions(accountID int) ([]db.PaymentHistoryModel, error) {
    return repositories.GetTransactionsByAccountID(accountID)
}

func CreateAccount(account *requests.CreateAccountRequest, userID int) (*db.AccountModel, error) {
	return repositories.CreateAccount(account, userID)
}