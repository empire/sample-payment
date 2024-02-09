package sdk

type PaymentService interface {
	Deposit(req *DepositRequest) (*DepositResponse, error)
	Balance(req *BalanceRequest) (*BalanceResponse, error)
	Withdraw(req *WithdrawRequest) (*WithdrawResponse, error)
}
