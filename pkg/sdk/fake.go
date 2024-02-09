package sdk

import (
	"errors"
	"fmt"
)

type FakePaymentService struct {
	accounts map[int]int
	calls    int
}

var _ PaymentService = (*FakePaymentService)(nil)

func NewFakePaymentService(accounts map[int]int) *FakePaymentService {
	return &FakePaymentService{
		accounts: accounts,
		calls:    0,
	}
}

func (f *FakePaymentService) Deposit(req *DepositRequest) (*DepositResponse, error) {
	f.calls++

	balance, ok := f.accounts[req.AccountID]
	if !ok {
		return nil, errors.New("account not found")
	}

	f.accounts[req.AccountID] = balance + req.Amount

	return &DepositResponse{
		Message: fmt.Sprintf("Deposit successful: new balance is %.2f", f.accounts[req.AccountID]),
	}, nil
}

func (f *FakePaymentService) Balance(req *BalanceRequest) (*BalanceResponse, error) {
	f.calls++

	balance, ok := f.accounts[req.AccountID]
	if !ok {
		return nil, errors.New("account not found")
	}

	return &BalanceResponse{
		Balance: balance,
	}, nil
}

func (f *FakePaymentService) Withdraw(req *WithdrawRequest) (*WithdrawResponse, error) {
	f.calls++

	balance, ok := f.accounts[req.AccountID]
	if !ok {
		return nil, errors.New("account not found")
	}

	if balance < req.Amount {
		return nil, errors.New("insufficient funds")
	}

	f.accounts[req.AccountID] = balance - req.Amount

	return &WithdrawResponse{
		Message: fmt.Sprintf("Withdraw successful: new balance is %.2f", f.accounts[req.AccountID]),
	}, nil
}

func (f *FakePaymentService) GetCalls() int {
	return f.calls
}

func (f *FakePaymentService) ResetCalls() {
	f.calls = 0
}
