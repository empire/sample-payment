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
		2: 0,
	}
	ps := fake.NewFakePaymentService(m)
	app := NewPeerToPeerPaymentApp(ps)
	msg, err := app.MakePayment(1, 2, 1000)
	require.NoError(t, err)
	require.Equal(t, "Payment of 1000 successfully transferred from user 1 to user 2\n", msg)
	// require.Equal(t, ps.Balance())
}