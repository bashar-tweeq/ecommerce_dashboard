package store

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type Product struct {
	Id    uuid.UUID
	Title string
	Price float64
	Qty   uint32
}

type ProductStore interface {
	CreateProduct(ctx context.Context, id uuid.UUID, title string, price float64, qty uint32) (*Product, error)
	GetProduct(ctx context.Context, id uuid.UUID) (*Product, error)
}

type store struct {
	session *pgx.Conn
}

func (s store) CreateProduct(ctx context.Context, id uuid.UUID, title string, price float64, qty uint32) (*Product, error) {

	insertSql := `INSERT INTO products(
        			id,
                	title,
                	price,
            		qty)
				VALUES ($1, $2, $3, $4)`

	err := s.session.QueryRow(ctx, insertSql, id, title, price, qty)

	if err != nil {
		return nil, fmt.Errorf("error in creating product: %v", err)
	}

	return &Product{
		Id:    id,
		Title: title,
		Price: price,
		Qty:   qty,
	}, nil
}

func (s store) GetProduct(ctx context.Context, id uuid.UUID) (*Product, error) {
	product := &Product{}

	err := s.session.QueryRow(ctx, `SELECT * FROM products WHERE id=$1`, id).Scan(
		&product.Id,
		&product.Title,
		&product.Price,
		&product.Qty)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("error in store.GetProduct: %v", err)
	}

	return product, nil
}

func NewProductStore(session *pgx.Conn) ProductStore {
	return &store{session: session}
}
