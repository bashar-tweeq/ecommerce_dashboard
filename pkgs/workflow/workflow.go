package workflow

import (
	"context"
	"errors"
	"time"

	transactionpb "ecommerce_dashboard/pkgs/transaction/proto"
	"ecommerce_dashboard/pkgs/transactions/store"
	"github.com/jackc/pgx/v5"
	"go.temporal.io/sdk/workflow"
)

type TransactionWorkflow struct {
	activities *Activity
}

func (w *TransactionWorkflow) CreateTransactionWorkflow(ctx workflow.Context, request transactionpb.CreateTransactionRequest) (*transactionpb.Transaction, error) {

	opts := workflow.ActivityOptions{StartToCloseTimeout: time.Second * 5}
	ctx = workflow.WithActivityOptions(ctx, opts)

	customer, err := w.activities.GetCustomerById(ctx, request.CustomerId)
	if err != nil {
		return nil, err
	}

	product, err := w.activities.GetProduct(ctx, request.ProductId)
	if err != nil {
		return nil, err
	}

	if product.Product.Qty != request.Qty {
		return nil, errors.New("insufficient product qty")
	}

	transaction, err := w.activities.CreateTransaction(context.Background(), transactionpb.CreateTransactionRequest{
		CustomerId: customer.Customer.CustomerId,
		ProductId:  product.Product.Id,
		Qty:        request.Qty,
		TotalPrice: float64(request.Qty) * product.Product.Price},
	)

	if err != nil {
		return nil, err
	}

	return transaction.Transaction, nil
}
