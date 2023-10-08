package store

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type Customer struct {
	Id        uuid.UUID
	FirstName string
	LastName  string
	Email     string
	Address   string
}

type CustomerStore interface {
	CreateCustomer(ctx context.Context, id uuid.UUID, email, firstName, lastName, address string) (*Customer, error)
	GetCustomerByEmail(ctx context.Context, email string) (*Customer, error)
	GetTopCustomers(ctx context.Context) ([]Customer, error)
}

type store struct {
	session *pgx.Conn
}

func (s store) CreateCustomer(ctx context.Context, id uuid.UUID, email, firstName, lastName, address string) (*Customer, error) {
	insertSql := `INSERT INTO customers(
       			id,
               	email,
                first_name,
                last_name,
           		address)
				VALUES ($1, $2, $3, $4, $5)`

	err := s.session.QueryRow(ctx, insertSql, id, email, firstName, lastName, address)

	if err != nil {
		return nil, fmt.Errorf("error in creating customer: %v", err)
	}

	return &Customer{
		Id:        id,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Address:   address,
	}, nil
}

func (s store) GetCustomerByEmail(ctx context.Context, email string) (*Customer, error) {
	customer := &Customer{}
	err := s.session.QueryRow(ctx, `SELECT * FROM customer WHERE email=$1`, email).Scan(
		&customer.Id,
		&customer.FirstName,
		&customer.LastName,
		&customer.Email,
		&customer.Address)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("error in store.GetCustomer: %v", err)
	}

	return customer, nil
}

func (s store) GetTopCustomers(ctx context.Context) ([]Customer, error) {
	customers := []Customer{}
	rows, err := s.session.Query(ctx, "select id, email, first_name, last_name, address from customers order by (select sum(total_price) from transactions where customer_id=id) limit $1", 5)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("cannot scan customer.\n\n%w", err)
	}

	for rows.Next() {
		customer := Customer{}
		err := rows.Scan(
			&customer.Id,
			&customer.Email,
			&customer.FirstName,
			&customer.LastName,
			&customer.Address,
		)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}
	return customers, nil
}
func NewCustomerStore(session *pgx.Conn) CustomerStore {
	return &store{session: session}
}
