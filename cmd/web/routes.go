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
	mux.Use(SessionLoad)

	// mux.Get("/paystack-terminal", app.PaystackTerminal)
	//mux.Get("/cart-view", app.CartView)

	mux.Get("/api/products", app.GetAllProducts)
	mux.Get("/api/products/{id}", app.GetProductsByID)
	mux.Get("/api/products/slug/{slug}", app.GetProductsBySlug)
	mux.Post("/api/save-transaction", app.CreateTransaction)
	mux.Post("/api/save-order", app.CreateOrder)

	//paystack shit
	mux.Post("/initialize-payment", app.initializePayment)
	mux.Get("/verify-payment", app.verifyPayment)

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
