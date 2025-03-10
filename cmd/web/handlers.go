package main

import "net/http"

func (app *application) HomePage(w http.ResponseWriter, r *http.Request) {
	if err := app.renderTemplates(w, r, "home", nil); err != nil {
		app.errorLog.Println(err)
	}
}

func (app *application) PaystackTerminal(w http.ResponseWriter, r *http.Request) {
	if err := app.renderTemplates(w, r, "terminal", nil); err != nil {
		app.errorLog.Println(err)
	}
}

func (app *application) CartView(w http.ResponseWriter, r *http.Request) {
	if err := app.renderTemplates(w, r, "cart", nil); err != nil {
		app.errorLog.Println(err)
	}
}
