package sdk

type (
	AccountID int
	Amount    int
)

type PaymentService interface {
	Deposit(accountID AccountID, amount Amount) (string, error)
	Balance(accountID AccountID) (Amount, error)
	Withdraw(accountID AccountID, amount Amount) (string, error)
}
