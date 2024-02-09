package p2p

import (
	"testing"

	"github.com/empire/sample-payment/pkg/sdk"
	"github.com/empire/sample-payment/pkg/sdk/mock_sdk"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func Test_Mock_MakePayment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ps := mock_sdk.NewMockPaymentService(ctrl)
	app := NewPeerToPeerPaymentApp(ps)

	var amount sdk.Amount
	var account sdk.AccountID
	ps.EXPECT().Withdraw(gomock.Any(), gomock.Any()).DoAndReturn(func(acc sdk.AccountID, am sdk.Amount) (string, error) {
		account = acc
		amount = am
		return "sample message", nil
	})

	ps.EXPECT().Deposit(gomock.Any(), gomock.Any()).Return("not actual response", nil)

	msg, err := app.MakePayment(1, 2, 1000)
	require.NoError(t, err)
	require.Equal(t, "Payment of 1000 successfully transferred from user 1 to user 2\n", msg)
	require.EqualValues(t, 1, account)
	require.EqualValues(t, 1000, amount)
}

func Test_Mock_CheckBalance(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ps := mock_sdk.NewMockPaymentService(ctrl)
	app := NewPeerToPeerPaymentApp(ps)

	ps.EXPECT().Balance(gomock.Any()).Return(1000, nil)

	amount, err := app.CheckBalance(1)
	require.NoError(t, err)
	require.EqualValues(t, 1000, amount)
}
