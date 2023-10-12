package workflow

import (
	"context"
	"errors"
	"time"

	transactionpb "ecommerce_dashboard/genproto/transactions/proto"
	"ecommerce_dashboard/pkgs/transaction/store"
	"github.com/jackc/pgx/v5"
	"go.temporal.io/sdk/workflow"
)

type TransactionWorkflow struct {
	activities *Activity
}

func (w *TransactionWorkflow) CreateTransactionWorkflow(ctx workflow.Context, request *pb.CreateTransactionRequest) (*pb.Transaction, error) {

	opts := workflow.ActivityOptions{StartToCloseTimeout: time.Second * 5}
	ctx = workflow.WithActivityOptions(ctx, opts)

	customer, err := w.activities.GetCustomerById(request.CustomerId)
	if err != nil {
		return nil, err
	}

	product, err := w.activities.GetProduct(request.ProductId)
	if err != nil {
		return nil, err
	}

	if product.Qty != request.Qty {
		return nil, errors.New("insufficient product qty")
	}

	transaction, err := w.activities.CreateTransaction(context.Background(), transactionpb.CreateTransactionRequest{
		CustomerId: customer.Id,
		ProductId: product.Id
		Qty: request.Qty,
		TotalPrice:  float64(request.Qty) * product.Price},
	)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}
