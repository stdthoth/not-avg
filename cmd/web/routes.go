package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
		MaxAge:           500,
	}))

	mux.Get("/", app.HomePage)
	mux.Get("/paystack-terminal", app.PaystackTerminal)
	mux.Post("/initialize-payment", app.initializePayment)
	mux.Get("/verify-payment", app.verifyPayment)
	mux.Get("/cart-view", app.CartView)

	return mux
}
