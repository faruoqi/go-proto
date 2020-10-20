package main

import (
	"context"
	"go-proto/customer"
	"google.golang.org/grpc"
	"io"
	"log"
)

const (
	ADDRESS = "localhost:50051"
)

func CreateCustomer(client customer.CustomerClient, customer *customer.CustomerRequest) {

	resp, err := client.CreateCustomer(context.Background(), customer)
	if err != nil {
		log.Fatalf("couldn't create customer : %v", err)
	}
	if resp.Success {
		log.Printf("A new costumer has been added with id : %v", resp.Id)
	}

}

func GetCustomers(client customer.CustomerClient, filter *customer.CustomerFilter) {

	stream, err := client.GetCustomers(context.Background(), filter)
	if err != nil {
		log.Fatalf("error on get customers : %v", err)
	}

	for {
		customer, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetCustomers(_) =  _, %v", client, err)
		}
		log.Printf("Customer : %v", customer)
	}

}

func main() {

	conn, err := grpc.Dial(ADDRESS, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("didn't connect to server %v", err)
	}
	defer conn.Close()
	client := customer.NewCustomerClient(conn)

	/*
		customer1 := &customer.CustomerRequest{
			Id:        101,
			Name:      "Oqi Faruoqi",
			Email:     "oqi@gmail.com",
			Phone:     "3430340",
			Addresses: []*customer.CustomerRequest_Address{
				&customer.CustomerRequest_Address{
					Street:     "Jl Parkit",
					City:       "Tangsel",
					State:      "Banten",
					Zip:        "43122",
					IsShipping: false,
				},
				&customer.CustomerRequest_Address{
					Street:     "Jl Fatmawati",
					City:       "Jaksel",
					State:      "DKI jakarta",
					Zip:        "42212",
					IsShipping: true,
				},
			},
		}


		CreateCustomer(client,customer1)

		customer2 := &customer.CustomerRequest{
			Id:        101,
			Name:      "Robbi",
			Email:     "robbi@gmail.com",
			Phone:     "2242232",
			Addresses: []*customer.CustomerRequest_Address{
				&customer.CustomerRequest_Address{
					Street:     "Jl sulaiman",
					City:       "Jaksel",
					State:      "DKI Jakarta",
					Zip:        "22121",
					IsShipping: false,
				},
				&customer.CustomerRequest_Address{
					Street:     "Buncit",
					City:       "Jaksel",
					State:      "DKI jakarta",
					Zip:        "22111",
					IsShipping: true,
				},
			},
		}

		CreateCustomer(client,customer2)

	*/

	filter := &customer.CustomerFilter{Keyword: "Robbi"}
	GetCustomers(client, filter)

}
