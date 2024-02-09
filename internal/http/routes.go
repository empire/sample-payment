package http

import (
	"net/http"

	"github.com/empire/sample-payment/internal/service"
)

func RegisterRoutes(service *service.PaymentService) {
	http.HandleFunc("/deposit", deposit(service))
	http.HandleFunc("/balance", balance(service))
	http.HandleFunc("/withdraw", withdraw(service))
}
