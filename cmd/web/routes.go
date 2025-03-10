package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", app.HomePage)
	mux.Get("/paystack-terminal", app.PaystackTerminal)
	mux.Get("/cart-view", app.CartView)

	return mux
}
