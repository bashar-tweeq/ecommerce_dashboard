package workflow

import (
	"context"
	customerspb "ecommerce_dashboard/genproto/customers/proto"
	productspb "ecommerce_dashboard/genproto/products/proto"
	transactionpb "ecommerce_dashboard/genproto/transactions/proto"
	customersv1 "ecommerce_dashboard/pkgs/customers/v1"
	productsv1 "ecommerce_dashboard/pkgs/products/v1"
	transactionv1 "ecommerce_dashboard/pkgs/transaction/v1"
)

type Activity struct {
	transactionService transactionv1.server
	customerService    customersv1.server
	productService     productsv1.server
}

func (a Activity) GetCustomerById(ctx context.Context, id string) (customerspb.Customer, error) {
	return a.customerService.GetCustomerByEmail(ctx, customerspb.GetCustomerByIdRequest{Id: id})
}

func (a Activity) GetProduct(ctx context.Context, id string) (productspb.Product, error) {
	return a.productService.GetProduct(ctx, productspb.GetProductByIdRequest{ProductId: id})
}

func (a Activity) CreateTransaction(ctx context.Context, transaction transactionpb.CreateTransactionRequest) (transactionpb.CreateTransactionResponse, error) {
	return a.transactionService.CreateTransaction(ctx, transactionpb.CreateTransactionRequest{
		CustomerId: transaction.customerId,
		ProductId:  transaction.ProductId,
		Qty:        transaction.Qty,
		TotalPrice: transaction.TotalPrice,
	})
}
