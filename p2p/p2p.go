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
	_, err := app.paymentService.Withdraw(sdk.AccountID(senderID), sdk.Amount(amount))
	if err != nil {
		return "", fmt.Errorf("failed to withdraw from sender's account: %w", err)
	}

	_, err = app.paymentService.Deposit(sdk.AccountID(receiverID), sdk.Amount(amount))
	if err != nil {
		return "", fmt.Errorf("failed to deposit to receiver's account: %w", err)
	}

	return fmt.Sprintf("Payment of %d successfully transferred from user %d to user %d\n", amount, senderID, receiverID), nil
}

func (app *PeerToPeerPaymentApp) CheckBalance(userID int) (int, error) {
	balance, err := app.paymentService.Balance(sdk.AccountID(userID))
	if err != nil {
		return 0, fmt.Errorf("failed to check balance for user %d: %v", userID, err)
	}

	return int(balance), nil
}
