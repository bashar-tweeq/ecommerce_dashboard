package v1

import (
	"context"
	"ecommerce_dashboard/pkgs/transaction/proto"
	"ecommerce_dashboard/pkgs/transactions/store"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type server struct {
	store                  store.TransactionStore
	transactionsStreamChan chan store.Transaction
}

func (s server) CreateTransaction(ctx context.Context, request *proto.CreateTransactionRequest) (*proto.CreateTransactionResponse, error) {
	id := uuid.New()

	_, err := s.store.CreateTransaction(ctx, id, request.CustomerId, request.ProductId, request.Qty, request.TotalPrice)

	if err != nil {
		return nil, fmt.Errorf("error service CreateProduct:  %v", err)
	}
	return nil, nil
}

func (s server) GetTransaction(ctx context.Context, request *proto.GetTransactionRequest) (*proto.GetTransactionResponse, error) {
	id := uuid.MustParse(request.Id)

	_, err := s.store.GetTransaction(ctx, id)

	if err != nil {
		return nil, fmt.Errorf("error service CreateCustomer:  %v", err)
	}
	return nil, nil
}

func (s server) GetTotalSales(ctx context.Context, request *proto.GetTotalSalesRequest) (*proto.GetTotalSalesResponse, error) {
	_, err := s.store.GetTotalSales(ctx)
	if err != nil {
		return nil, fmt.Errorf("error service GetTotalSalesRequest:  %v", err)
	}
	return nil, nil
}
func (s server) StreamTransaction(request *proto.StreamTransactionRequest, stream proto.TransactionService_StreamTransactionServer) error {

	for {
		transaction := <-s.transactionsStreamChan
		err := stream.Send(&proto.Transaction{
			Id:         transaction.Id.String(),
			CustomerId: transaction.CustomerId,
			ProductId:  transaction.ProductId,
			Qty:        transaction.Qty,
			TotalPrice: transaction.TotalPrice,
		})
		if err != nil {
			return fmt.Errorf("error in StreamTransaction: \n%w", err)
		}
	}
}

func (s server) GetTotalSalesByProduct(ctx context.Context, request *proto.GetTotalSalesByProductRequest) (*proto.GetTotalSalesByProductResponse, error) {
	_, err := s.store.GetTotalSalesByProduct(ctx)
	if err != nil {
		return nil, fmt.Errorf("error service GetTotalSalesByProduct:  %v", err)
	}
	return nil, nil
}

func New(conn *pgx.Conn) proto.TransactionServiceServer {
	return &server{
		store: store.NewTransactionStore(conn),
	}
}
