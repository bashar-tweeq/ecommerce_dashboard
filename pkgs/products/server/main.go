package main

import (
	"context"
	"ecommerce_dashboard/genproto/products/proto"
	productsv1 "ecommerce_dashboard/pkgs/products/v1"
	"fmt"
	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	port := "8082"

	conn, err := pgx.Connect(context.Background(), "postgresql://root@localhost:26257/defaultdb?sslmode=disable")

	if err != nil {
		log.Println("cannot connect to database")
		log.Fatal(err.Error())
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))

	if err != nil {
		log.Println(fmt.Sprintf("cannot listen on :%s"), port)
		log.Fatal(err.Error())
	}

	s := grpc.NewServer()
	proto.RegisterProductServiceServer(s, productsv1.New(conn))
	log.Printf("server lis at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}