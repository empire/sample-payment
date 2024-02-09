package fake

import (
	"errors"
	"fmt"

	"github.com/empire/sample-payment/pkg/sdk"
)

type FakePaymentService struct {
	accounts map[sdk.AccountID]sdk.Amount
}

var _ sdk.PaymentService = (*FakePaymentService)(nil)

func NewFakePaymentService(accounts map[sdk.AccountID]sdk.Amount) *FakePaymentService {
	return &FakePaymentService{
		accounts: accounts,
	}
}

func (f *FakePaymentService) Deposit(accountID sdk.AccountID, amount sdk.Amount) (string, error) {
	balance, ok := f.accounts[accountID]
	if !ok {
		return "", errors.New("account not found")
	}

	f.accounts[accountID] = balance + amount

	return fmt.Sprintf("Deposit successful: new balance is %d", f.accounts[accountID]), nil
}

func (f *FakePaymentService) Balance(accountID sdk.AccountID) (sdk.Amount, error) {
	balance, ok := f.accounts[accountID]
	if !ok {
		return 0, errors.New("account not found")
	}

	return balance, nil
}

func (f *FakePaymentService) Withdraw(accountID sdk.AccountID, amount sdk.Amount) (string, error) {
	balance, ok := f.accounts[accountID]
	if !ok {
		return "", errors.New("account not found")
	}

	if balance < amount {
		return "", errors.New("insufficient funds")
	}

	f.accounts[accountID] = balance - amount

	return fmt.Sprintf("Withdraw successful: new balance is %d", f.accounts[accountID]), nil
}
