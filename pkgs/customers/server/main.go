package main

import (
	"context"
	pb "ecommerce_dashboard/pkgs/customers/proto"
	customersv1 "ecommerce_dashboard/pkgs/customers/v1"
	"fmt"
	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

func main() {
	port := "8080"
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbUser := os.Getenv("DATABASE_USER")

	connStr := fmt.Sprintf("postgresql://%s@%s:%s/defaultdb?sslmode=disable", dbUser, dbHost, dbPort)
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Println("cannot connect to database")
		log.Fatal(err.Error())
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", port))

	if err != nil {
		log.Println(fmt.Sprintf("cannot listen on 0.0.0.0:%s"), port)
		log.Fatal(err.Error())
	}

	s := grpc.NewServer()
	pb.RegisterCustomerServiceServer(s, customersv1.New(conn))
	reflection.Register(s)

	log.Printf("server listen at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
