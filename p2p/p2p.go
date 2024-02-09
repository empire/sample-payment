package p2p

import (
	"fmt"

	"github.com/empire/sample-payment/pkg/sdk"
)

type PeerToPeerPaymentApp struct {
	paymentService sdk.PaymentService
}

func NewPeerToPeerPaymentApp(service sdk.PaymentService) *PeerToPeerPaymentApp {
	return &PeerToPeerPaymentApp{
		paymentService: service,
	}
}

func (app *PeerToPeerPaymentApp) MakePayment(senderID, receiverID int, amount int) (string, error) {
	withdrawReq := &sdk.WithdrawRequest{
		AccountID: senderID,
		Amount:    amount,
	}
	_, err := app.paymentService.Withdraw(withdrawReq)
	if err != nil {
		return "", fmt.Errorf("failed to withdraw from sender's account: %w", err)
	}

	depositReq := &sdk.DepositRequest{
		AccountID: receiverID,
		Amount:    amount,
	}
	_, err = app.paymentService.Deposit(depositReq)
	if err != nil {
		return "", fmt.Errorf("failed to deposit to receiver's account: %w", err)
	}

	return fmt.Sprintf("Payment of %d successfully transferred from user %d to user %d\n", amount, senderID, receiverID), nil
}

func (app *PeerToPeerPaymentApp) CheckBalance(userID int) (int, error) {
	balanceReq := &sdk.BalanceRequest{
		AccountID: userID,
	}
	balanceResp, err := app.paymentService.Balance(balanceReq)
	if err != nil {
		return 0, fmt.Errorf("Failed to check balance for user %d: %v", userID, err)
	}

	return balanceResp.Balance, nil
}
