package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

/*
func (app *application) HomePage(w http.ResponseWriter, r *http.Request) {

	productmodel, err := app.DB.GetAllProducts()
	if err != nil {
		fmt.Println(err)
	}

	data := make(map[string]interface{})
	data["productmodel"] = productmodel

	if err != nil {
		http.Error(w, "unable to fetch products", http.StatusInternalServerError)
	}

	err = app.renderTemplates(w, r, "home", &templateData{
		Data: data,
	})
	if err != nil {
		app.errorLog.Println("couldnt render template")
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
*/

func (app *application) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := app.DB.GetAllProducts()
	if err != nil {
		app.errorLog.Println(err)
	}
	jsonOut, err := json.Marshal(products)
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonOut)
}

func (app *application) GetProductsByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	productId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		app.errorLog.Println(err)
	}

	product, err := app.DB.GetProduct(productId)
	if err != nil {
		app.errorLog.Println(err)
	}

	jsonOut, err := json.Marshal(product)
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonOut)
}

func (app *application) GetProductsBySlug(w http.ResponseWriter, r *http.Request) {
	slug := strings.TrimPrefix(r.URL.Path, "/api/products/slug/")

	product, err := app.DB.GetProductBySlug(slug)

	if err != nil {
		app.errorLog.Println(err)
	}

	jsonOut, err := json.Marshal(product)
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonOut)
}

func (app *application) CreateTransaction(w http.ResponseWriter, r *http.Request) {

}

func (app *application) CreateOrder(w http.ResponseWriter, r *http.Request) {

}
