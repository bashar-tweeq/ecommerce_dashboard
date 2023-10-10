package v1

import (
	"context"
	"ecommerce_dashboard/genproto/customers/proto"
	"ecommerce_dashboard/pkgs/customers/store"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"log"
)

type server struct {
	store store.CustomerStore
}

func (s server) CreateCustomer(ctx context.Context, request *proto.CreateCustomerRequest) (*proto.CreateCustomerResponse, error) {
	id := uuid.New()
	cu, err := s.store.CreateCustomer(ctx, id, request.Email, request.FirstName, request.LastName, request.Address)

	if err != nil {
		return nil, fmt.Errorf("error in service CreateCustomer:  %v", err)
	}

	return &proto.CreateCustomerResponse{
		Customer: &proto.Customer{
			CustomerId: cu.Id.String(),
			Email:      cu.Email,
			FirstName:  cu.FirstName,
			LastName:   cu.LastName,
			Address:    cu.Address,
		},
	}, nil
}

func (s server) GetCustomerByEmail(ctx context.Context, request *proto.GetCustomerByEmailRequest) (*proto.GetCustomerByEmailResponse, error) {
	log.Printf("called GetCustomerByEmail: %v\n", request)
	
	cu, err := s.store.GetCustomerByEmail(ctx, request.Email)

	if err != nil {
		return nil, fmt.Errorf("error in service GetCustomerByEmail:  %v", err)
	}

	return &proto.GetCustomerByEmailResponse{
		Customer: &proto.Customer{
			CustomerId: cu.Id.String(),
			Email:      cu.Email,
			FirstName:  cu.FirstName,
			LastName:   cu.LastName,
			Address:    cu.Address,
		},
	}, nil
}

func (s server) GetTopCustomers(ctx context.Context, request *proto.GetTopCustomersRequest) (*proto.GetTopCustomersResponse, error) {
	topCustomers, err := s.store.GetTopCustomers(ctx)

	if err != nil {
		return nil, fmt.Errorf("error in service GetTopCustomers:  %v", err)

	}

	customersPb := []*proto.Customer{}
	for _, cu := range topCustomers {
		customersPb = append(customersPb, &proto.Customer{
			CustomerId: cu.Id.String(),
			Email:      cu.Email,
			FirstName:  cu.FirstName,
			LastName:   cu.LastName,
			Address:    cu.Address,
		})
	}

	return &proto.GetTopCustomersResponse{TopCustomers: customersPb}, nil
}

func New(conn *pgx.Conn) proto.CustomerServiceServer {
	return &server{
		store: store.NewCustomerStore(conn),
	}
}
