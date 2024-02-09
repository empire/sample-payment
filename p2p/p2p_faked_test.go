package p2p

import (
	"testing"

	"github.com/empire/sample-payment/pkg/sdk"
	"github.com/empire/sample-payment/pkg/sdk/fake"
	"github.com/stretchr/testify/require"
)

func Test_Faked_MakeRequest(t *testing.T) {
	m := map[sdk.AccountID]sdk.Amount{
		1: 1500,
		2: 200,
	}
	ps := fake.NewFakePaymentService(m)
	app := NewPeerToPeerPaymentApp(ps)
	msg, err := app.MakePayment(1, 2, 1000)
	require.NoError(t, err)
	require.Equal(t, "Payment of 1000 successfully transferred from user 1 to user 2\n", msg)

	amount, err := ps.Balance(1)
	require.NoError(t, err)
	require.EqualValues(t, 500, amount)

	amount, err = ps.Balance(2)
	require.NoError(t, err)
	require.EqualValues(t, 1200, amount)
}

func Test_Faked_CheckBalance(t *testing.T) {
	ps := fake.NewFakePaymentService(map[sdk.AccountID]sdk.Amount{1: 1500})
	app := NewPeerToPeerPaymentApp(ps)
	amount, err := app.CheckBalance(1)
	require.NoError(t, err)
	require.Equal(t, 1500, amount)
}
