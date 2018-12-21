package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/amiralii/shipping/consignment-service/proto/consignment"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
)

const (
	address         = "127.0.0.1:8787"
	defualtFilename = "consignment.json"
)

func ParseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, nil
}

func main() {

	cmd.Init()

	// // Set up a connection to the server.
	// conn, err := grpc.Dial(address, grpc.WithInsecure())
	// if err != nil {
	// 	log.Fatalf("Did not connect: %v", err)
	// }
	// defer conn.Close()
	client := pb.NewShippingServiceClient("go.micro.srv.consignment", microclient.DefaultClient)

	// Contact the server and print out its response.
	file := defualtFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := ParseFile(file)
	if err != nil {
		log.Fatalf("Could not parse file:  %v", err)
	}

	r, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("Could not greet:  %v", err)
	}
	log.Printf("Created: %t", r.Created)

	getAll, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}
	for _, v := range getAll.Consignments {
		log.Println(v)
	}
}
