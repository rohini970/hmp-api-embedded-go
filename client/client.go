package main

import (
	"context"
	"flag"
	"log"

	pb "hmp-api-embedded-go/proto"

	"google.golang.org/grpc"
)

var (
	serverAddress = flag.String("server_address", "localhost:50051", "The server address in the format host:port")
)

func main() {
	flag.Parse()

	// Set up a connection to the server using the user-defined address.
	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a client instance.
	c := pb.NewHomeGatewayClient(conn)

	// Contact the server and print out its response.
	response, err := c.GetDeviceInfo(context.Background(), &pb.EmptyRequest{})
	if err != nil {
		log.Fatalf("could not get device info: %v", err)
	}

	deviceInfo := response
	log.Println("Device Info:")
	log.Printf("Type: %s", deviceInfo.Type)
	log.Printf("ID: %s", deviceInfo.Id)
	log.Printf("Manufacturer OUI: %s", deviceInfo.ManufacturerOUI)
	log.Printf("Manufacturer: %s", deviceInfo.Manufacturer)
	log.Printf("Model Name: %s", deviceInfo.ModelName)
	log.Printf("Serial Number: %s", deviceInfo.SerialNumber)
	log.Printf("Product Class: %s", deviceInfo.ProductClass)
	log.Printf("Software Version: %s", deviceInfo.SoftwareVersion)
	log.Printf("Hardware Version: %s", deviceInfo.HardwareVersion)
	log.Printf("Status Last Change: %s", deviceInfo.StatusLastChange)
}
