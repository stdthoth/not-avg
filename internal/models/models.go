package models

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

// DBMOdels is the type for database connection values
type DBModels struct {
	DB *sql.DB
}

// wrapper for all models
type Models struct {
	DB DBModels
}

func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModels{DB: db},
	}
}

// Products is the model type for all products
type Product struct {
	ID             int64     `json:"id"`
	Slug           string    `json:"slug"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	InventoryLevel int       `json:"inventory_level"`
	Price          int       `json:"price"`
	CreatedAt      time.Time `json:"-"`
	UpdatedAt      time.Time `json:"-"`
}

type Orders struct {
	ID            int       `json:"id"`
	ProductID     string    `json:"product_id"`
	TransactionID int       `json:"transaction_id"`
	StatusID      string    `json:"status_id"`
	CustomerName  string    `json:"customer_name"`
	CustomerEmail string    `json:"customer_email"`
	Quantity      int       `json:"quantity"`
	Address       string    `json:"address"`
	Amount        int       `json:"amount"`
	CreatedAt     time.Time `json:"-"`
	UpdatedAt     time.Time `json:"-"`
}

type Status struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	Delivered bool      `json:"delivered"`
}

type Transaction struct {
	ID                   int       `json:"id"`
	CustomerName         string    `json:"customer_name"`
	CustomerEmail        string    `json:"email"`
	Amount               int       `json:"amount"`
	TransactionReference string    `json:"reference"`
	CreatedAt            time.Time `json:"-"`
}

type TransactionStat struct {
	ID      string `json:"id"`
	Amount  int    `json:"amount"`
	Verfied bool   `json:"verified"`
}

func (m *DBModels) GetProduct(id int64) (Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var product Product

	row := m.DB.QueryRowContext(ctx, `select id, slug, name, description, inventory_level, price from products where id = ?`, id)
	err := row.Scan(&product.ID, &product.Slug, &product.Name, &product.Description, &product.InventoryLevel, &product.Price)
	if err != nil {
		return product, err
	}

	return product, nil

}

// GetAllProducts retrieves all products from the database
func (m *DBModels) GetAllProducts() ([]Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var products []Product

	query := `
        select id, slug, name, description, inventory_level, price, created_at, updated_at
        from products 
    `
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p Product
		err := rows.Scan(
			&p.ID,
			&p.Slug,
			&p.Name,
			&p.Description,
			&p.InventoryLevel,
			&p.Price,
			&p.CreatedAt,
			&p.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func (m *DBModels) GetProductBySlug(slug string) (Product, error) {
	var product Product

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
        SELECT 
            id, slug, name, description, inventory_level, price, created_at, updated_at 
        FROM 
            products 
        WHERE 
            slug = ?
    `

	row := m.DB.QueryRowContext(ctx, query, slug)
	err := row.Scan(
		&product.ID,
		&product.Slug,
		&product.Name,
		&product.Description,
		&product.InventoryLevel,
		&product.Price,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	switch {
	case err == sql.ErrNoRows:
		return Product{}, fmt.Errorf("product with slug '%s' not found", slug)
	case err != nil:
		return Product{}, fmt.Errorf("error querying product: %v", err)
	default:
		return product, nil
	}
}

// inserts a transasction into database
func (m *DBModels) InsertTransaction(txn Transaction) (int, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		insert into transactions
			(customer_name, amount,
			reference, created_at)
		values (?, ?, ?, ?)
		`
	result, err := m.DB.ExecContext(ctx, query,
		txn.CustomerName,
		txn.Amount,
		txn.TransactionReference,
		time.Now(),
	)

	if err != nil {
		return 0, nil
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *DBModels) InsertOrder(ord Orders) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		insert into orders
			(product_id, transaction_id, status_id, customer_name, customer_email,
			quantity, address, amount, created_at, updated_at)
		values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`
	result, err := m.DB.ExecContext(ctx, query,
		ord.ProductID,
		ord.TransactionID,
		ord.StatusID,
		ord.CustomerName,
		ord.CustomerEmail,
		ord.Quantity,
		ord.Address,
		ord.Amount,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return 0, nil
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
