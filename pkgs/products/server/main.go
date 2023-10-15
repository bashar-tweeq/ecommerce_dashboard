package main

import (
	"context"
	"ecommerce_dashboard/pkgs/product/proto"
	productsv1 "ecommerce_dashboard/pkgs/products/v1"
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

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))

	if err != nil {
		log.Println(fmt.Sprintf("cannot listen on :%s"), port)
		log.Fatal(err.Error())
	}

	s := grpc.NewServer()
	proto.RegisterProductServiceServer(s, productsv1.New(conn))
	log.Printf("server is now listening at %v", lis.Addr())
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
