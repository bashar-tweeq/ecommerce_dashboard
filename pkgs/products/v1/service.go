package v1

import (
	"context"
	"ecommerce_dashboard/pkgs/product/proto"
	"ecommerce_dashboard/pkgs/products/store"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type server struct {
	store store.ProductStore
}

func (s server) CreateProduct(ctx context.Context, request *proto.CreateProductRequest) (*proto.CreateProductResponse, error) {
	id := uuid.New()

	product, err := s.store.CreateProduct(ctx, id, request.Title, request.Price, request.Qty)

	if err != nil {
		return nil, fmt.Errorf("error service CreateProduct:  %v", err)
	}

	return &proto.CreateProductResponse{Product: &proto.Product{
		Id:    product.Id.String(),
		Title: product.Title,
		Price: product.Price,
		Qty:   product.Qty,
	}}, nil
}

func (s server) GetProduct(ctx context.Context, request *proto.GetProductByIdRequest) (*proto.GetProductByIdResponse, error) {
	pid := uuid.MustParse(request.ProductId)

	product, err := s.store.GetProduct(ctx, pid)

	if err != nil {
		return nil, fmt.Errorf("error service CreateCustomer:  %v", err)
	}

	return &proto.GetProductByIdResponse{Product: &proto.Product{
		Id:    product.Id.String(),
		Title: product.Title,
		Price: product.Price,
		Qty:   product.Qty,
	}}, nil
}
func New(conn *pgx.Conn) proto.ProductServiceServer {
	return &server{
		store: store.NewProductStore(conn),
	}
}
