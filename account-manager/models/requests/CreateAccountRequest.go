package requests

import "be-test-concrete-ai/account-manager/prisma/db"

type CreateAccountRequest struct {
	Type    string  `json:"type"  validate:"oneof=credit debit loan"`
	Balance float64 `json:"balance"`
}

func (a *CreateAccountRequest) GetType() db.AccountType {
	m := map[string]db.AccountType{
		"credit": db.AccountTypeCredit,
		"debit":  db.AccountTypeDebit,
		"loan":   db.AccountTypeLoan,
	}
	return m[a.Type]
}
