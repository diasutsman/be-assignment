package repositories

import (
	"be-test-concrete-ai/account-manager/config"
	"be-test-concrete-ai/account-manager/models/requests"
	"be-test-concrete-ai/account-manager/prisma/db"
	"context"
)

func GetAccountsByUserID(userID int) ([]db.AccountModel, error) {
	println("userID", userID)
	accounts, err := config.Client.Account.FindMany(
		db.Account.UserID.Equals(userID),
	).With(
		db.Account.History.Fetch(),
	).Exec(context.Background())

	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func GetTransactionsByAccountID(accountID int) ([]db.PaymentHistoryModel, error) {
	transactions, err := config.Client.PaymentHistory.FindMany(
		db.PaymentHistory.ID.Equals(accountID),
	).Exec(context.Background())

	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func CreateAccount(account *requests.CreateAccountRequest, UserID int) (*db.AccountModel, error) {
	return config.Client.Account.CreateOne(
		db.Account.Type.Set(account.GetType()),
		db.Account.User.Link(
			db.User.ID.Equals(UserID),
		),
		db.Account.Balance.Set(account.Balance),
	).Exec(context.Background())
}
