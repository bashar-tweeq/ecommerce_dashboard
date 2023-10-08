//package workflow
//
//import (
//	"context"
//	"errors"
//	"fmt"
//	"time"
//
//	pbC "ecommerce_dashboard/genproto/customers/proto"
//	pbP "ecommerce_dashboard/genproto/products/proto"
//	pbT "ecommerce_dashboard/genproto/transaction/proto"
//	"ecommerce_dashboard/pkgs/transaction/store"
//	"github.com/jackc/pgx/v5"
//	"go.temporal.io/sdk/workflow"
//)
//
////
////type TransactionWorkflow struct {
////	activities *WorkflowActivity
////}
////
////func (w *TransactionWorkflow) CreateTransactionWorkflow(ctx workflow.Context, request *pb.CreateTransactionRequest) (*pb.Transaction, error) {
////
////	opts := workflow.ActivityOptions{StartToCloseTimeout: time.Second * 3}
////	ctx = workflow.WithActivityOptions(ctx, opts)
////
////	customer, err := w.activities.GetCustomerById(request.CustomerId)
////	if err != nil {
////		return nil, err
////	}
////
////	product, err := w.activities.GetProductById(request.ProductId)
////	if err != nil {
////		return nil, err
////	}
////
////	if product.AvailableQuantity != request.Quantity {
////		return nil, errors.New("insufficient product quantity")
////	}
////
////	transaction := transactions.CreateTransactionRequestToModel(request)
////	transaction.TotalPrice = float64(request.Quantity) * product.Price
////
////	createdTransaction, err := w.activities.CreateTransaction(context.Background(), transaction)
////	if err != nil {
////		return nil, err
////	}
////
////	return transactions.TransactionToProto(createdTransaction), nil
////
////}
