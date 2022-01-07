package domain

import (
	"github.com/nicholasanthonys/hexagonal-banking/dto"
	"github.com/nicholasanthonys/hexagonal-banking/errs"
)

// Domain object
type Account struct {
	AccountId   string
	CustomerId  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{AccountId: a.AccountId}
}

// repository
type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}
