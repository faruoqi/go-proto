package main

import (
	"context"
	"go-proto/customer"
	"google.golang.org/grpc"
	"log"
	"net"
	"strings"
)

const (
	PORT = ":50051"
)

type server struct {
	savedCustomers []*customer.CustomerRequest
}

func (s *server) CreateCustomer(ctx context.Context, in *customer.CustomerRequest) (*customer.CustomerResponse, error) {
	s.savedCustomers = append(s.savedCustomers, in)
	return &customer.CustomerResponse{
		Id:      in.Id,
		Success: true,
	}, nil
}

func (s *server) GetCustomers(filter *customer.CustomerFilter, stream customer.Customer_GetCustomersServer) error {
	for _, cust := range s.savedCustomers {
		if filter.Keyword != "" {
			if !strings.Contains(cust.Name, filter.Keyword) {
				continue
			}
		}

		if err := stream.Send(cust); err != nil {
			return err
		}
	}

	return nil

}

func main() {

	listen, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}
	s := grpc.NewServer()
	customer.RegisterCustomerServer(s, &server{})
	s.Serve(listen)
}
