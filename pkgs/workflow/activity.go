package workflow

import (
	"context"
	customerspb "ecommerce_dashboard/pkgs/customers/proto"
	productspb "ecommerce_dashboard/pkgs/product/proto"
	transactionpb "ecommerce_dashboard/pkgs/transaction/proto"
)

type Activity struct {
	transactionService transactionpb.TransactionServiceServer
	customerService    customerspb.CustomerServiceServer
	productService     productspb.ProductServiceServer
}

func (a Activity) GetCustomerById(ctx context.Context, id string) (*customerspb.GetCustomerByIdResponse, error) {
	cu, err := a.customerService.GetCustomerById(
		ctx, &customerspb.GetCustomerByIdRequest{Id: id},
	)

	if err != nil {
		return nil, err
	}

	return &customerspb.GetCustomerByIdResponse{
		Customer: &customerspb.Customer{
			CustomerId: cu.Customer.CustomerId,
			Email:      cu.Customer.Email,
			FirstName:  cu.Customer.FirstName,
			LastName:   cu.Customer.LastName,
			Address:    cu.Customer.Address,
		},
	}, nil
}

func (a Activity) GetProduct(ctx context.Context, id string) (*productspb.GetProductByIdResponse, error) {
	product, err := a.productService.GetProduct(ctx, &productspb.GetProductByIdRequest{ProductId: id})

	if err != nil {
		return nil, err
	}
	return &productspb.GetProductByIdResponse{Product: &productspb.Product{
		Id:    product.Product.Id,
		Title: product.Product.Title,
		Price: product.Product.Price,
		Qty:   product.Product.Qty,
	}}, nil
}

func (a Activity) CreateTransaction(ctx context.Context, transaction transactionpb.CreateTransactionRequest) (*transactionpb.CreateTransactionResponse, error) {
	tx, err := a.transactionService.CreateTransaction(ctx, &transactionpb.CreateTransactionRequest{
		CustomerId: transaction.CustomerId,
		ProductId:  transaction.ProductId,
		Qty:        transaction.Qty,
		TotalPrice: transaction.TotalPrice,
	})

	if err != nil {
		return nil, err
	}

	return &transactionpb.CreateTransactionResponse{Transaction: &transactionpb.Transaction{
		Id:         tx.Transaction.Id,
		CustomerId: tx.Transaction.CustomerId,
		ProductId:  tx.Transaction.ProductId,
		Qty:        tx.Transaction.Qty,
		TotalPrice: tx.Transaction.TotalPrice,
	}}, nil
}
