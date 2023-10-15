package store

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type Transaction struct {
	Id         uuid.UUID
	CustomerId string
	ProductId  string
	Qty        uint32
	TotalPrice float64
}

type ProductBySales struct {
	Id    uuid.UUID
	Total float64
}

type TransactionStore interface {
	CreateTransaction(ctx context.Context, id uuid.UUID, customerId, productId string, qty uint32, totalPrice float64) (*Transaction, error)
	GetTransaction(ctx context.Context, id uuid.UUID) (*Transaction, error)
	GetTotalSales(ctx context.Context) (float64, error)
	GetTotalSalesByProduct(ctx context.Context) ([]ProductBySales, error)
}

type store struct {
	session *pgx.Conn
}

func (s store) CreateTransaction(ctx context.Context, id uuid.UUID, customerId, productId string, qty uint32, totalPrice float64) (*Transaction, error) {
	insertSql := `INSERT INTO transactions(
        			id,
                	customer_id,
                    product_id,
                	qty,
            		total_price)
				VALUES ($1, $2, $3, $4, $5)`

	commandTag, err := s.session.Exec(ctx, insertSql, id, customerId, productId, qty, totalPrice)

	if err != nil {
		return nil, fmt.Errorf("error in creating transaction : %v", err)
	}

	if commandTag.RowsAffected() != 1 {
		return nil, errors.New(commandTag.String())
	}

	return &Transaction{
		Id:         id,
		CustomerId: customerId,
		ProductId:  productId,
		Qty:        qty,
		TotalPrice: totalPrice,
	}, nil
}

func (s store) GetTransaction(ctx context.Context, id uuid.UUID) (*Transaction, error) {
	transaction := &Transaction{}

	err := s.session.QueryRow(ctx, `SELECT * FROM transactions WHERE id=$1`, id).Scan(
		&transaction.Id,
		&transaction.CustomerId,
		&transaction.ProductId,
		&transaction.Qty,
		&transaction.TotalPrice)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("error in store.Transaction: %v", err)
	}

	return transaction, nil
}

func (s store) GetTotalSales(ctx context.Context) (float64, error) {
	totalSales := float64(0)

	err := s.session.QueryRow(ctx, "select sum(total_price) as total from transactions").Scan(&totalSales)

	if err != nil {
		return 0, fmt.Errorf("error in store.GetTotalSales %w", err)
	}

	return totalSales, nil
}

func (s store) GetTotalSalesByProduct(ctx context.Context) ([]ProductBySales, error) {

	rows, err := s.session.Query(ctx, "select product_id, sum(total_price) as total_sales from transactions group by product_id")

	if err != nil {
		return nil, fmt.Errorf("error in store.GetSalesByProduct.\n%w", err)
	}

	productSalesList := []ProductBySales{}
	for rows.Next() {
		productSales := ProductBySales{}
		err := rows.Scan(&productSales.Id, &productSales.Total)
		if err != nil {
			return nil, err
		}
		productSalesList = append(productSalesList, productSales)
	}

	return productSalesList, nil
}

func NewTransactionStore(session *pgx.Conn) TransactionStore {
	return &store{session: session}
}
