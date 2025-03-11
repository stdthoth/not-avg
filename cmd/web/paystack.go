package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-resty/resty/v2"
)

// Initialize Payment Handler (Backend)
func (app *application) initializePayment(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		Email  string `json:"email"`
		Amount int64  `json:"amount"`
	}

	var req Request
	json.NewDecoder(r.Body).Decode(&req)

	client := resty.New()
	resp, err := client.R().
		SetAuthToken(app.config.paystack.secret).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"email":  req.Email,
			"amount": req.Amount * 100, // Convert to kobo
		}).
		Post("https://api.paystack.co/transaction/initialize")

	if err != nil {
		http.Error(w, "Payment initialization failed", http.StatusInternalServerError)
		return
	} else {
		app.infoLog.Println("elon wheere my rocket ship")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp.Body())
}

// Verify Payment Handler (Backend)
func (app *application) verifyPayment(w http.ResponseWriter, r *http.Request) {
	reference := r.URL.Query().Get("reference")
	client := resty.New()
	resp, err := client.R().SetAuthToken(app.config.paystack.secret).Get("https://api.paystack.co/transaction/verify/" + reference)
	if err != nil {
		app.errorLog.Println("could not verify your payment")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp.Body())
}
