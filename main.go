package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	handlers "github.com/empire/sample-payment/internal/http"
	"github.com/empire/sample-payment/internal/service"
	"github.com/empire/sample-payment/p2p"
	"github.com/empire/sample-payment/pkg/sdk"
)

var serverFlag bool

func init() {
	flag.BoolVar(&serverFlag, "server", true, "run as a server")
}

func main() {
	flag.Parse()

	if serverFlag {

		log.Println("Running as a server")
		runServer()
	} else {

		log.Println("Running as a client")
		runClient()
	}
}

func runServer() {
	ps, err := service.NewPaymentService("file::memory:")
	if err != nil {
		log.Fatal(err)
	}

	handlers.RegisterRoutes(ps)

	fmt.Println("Starting web server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func runClient() {
	client := sdk.NewClient("http://localhost:8080")
	client.Deposit(&sdk.DepositRequest{AccountID: 1, Amount: 230})

	app := p2p.NewPeerToPeerPaymentApp(client)

	balance, err := app.CheckBalance(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Current balance is: %d\n", balance)

	if _, err := app.MakePayment(1, 2, 70); err != nil {
		log.Fatal(err)
	}

	balance, err = app.CheckBalance(1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Current balance is: %d\n", balance)
}
