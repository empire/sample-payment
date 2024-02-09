package http

import (
	"encoding/json"
	"net/http"

	"github.com/empire/sample-payment/internal/service"
)

func deposit(service *service.PaymentService) http.HandlerFunc {
	type request struct {
		AccountID int     `json:"accountID"` // The ID of the account to deposit to
		Amount    float64 `json:"amount"`    // The amount to deposit
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var req request
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Call the Deposit method of the payment service
		err = service.Deposit(req.AccountID, req.Amount)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Deposit successful"))
	}
}

func balance(service *service.PaymentService) http.HandlerFunc {
	type request struct {
		AccountID int `json:"accountID"` // The ID of the account to check the balance
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var req request
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		balance, err := service.Balance(req.AccountID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]float64{"balance": balance})
	}
}

func withdraw(service *service.PaymentService) http.HandlerFunc {
	type request struct {
		AccountID int     `json:"accountID"` // The ID of the account to withdraw from
		Amount    float64 `json:"amount"`    // The amount to withdraw
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var req request
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = service.Withdraw(req.AccountID, req.Amount)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Withdraw successful"})
	}
}
