package main

import (
	"fmt"
	"log"

	"github.com/empire/sample-payment/payment"
)

func main() {
	service, err := payment.NewPaymentService("file::memory:")
	if err != nil {
		log.Fatal(err)
	}

	// Test the payment service
	err = service.Deposit(1, 100) // Deposit 100 to account 1
	if err != nil {
		log.Fatal(err)
	}
	balance, err := service.Balance(1) // Check the balance of account 1
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance)          // Should print 100
	err = service.Withdraw(1, 50) // Withdraw 50 from account 1
	if err != nil {
		log.Fatal(err)
	}
	balance, err = service.Balance(1) // Check the balance of account 1
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance) // Should print 50
}
