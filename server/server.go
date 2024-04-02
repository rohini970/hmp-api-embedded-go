package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "hmp-api-embedded-go/proto"
	"hmp-api-embedded-go/services/utils"

	"google.golang.org/grpc"
)

type homeGateway struct {
	pb.UnimplementedHomeGatewayServer
}

var (
	ip   = flag.String("ip", "localhost", "The server IP address")
	port = flag.Int("port", 50051, "The server port")
)

func (s *homeGateway) GetDeviceInfo(ctx context.Context, req *pb.EmptyRequest) (*pb.DeviceInfo, error) {
	// Implement the logic to handle GetDeviceInfo RPC
	fmt.Println("Got a request:", req)

	// Fetch device information using utility function
	manufacturerOUI, _ := utils.GetDataModelValue("Device.DeviceInfo.ManufacturerOUI")
	manufacturer, _ := utils.GetDataModelValue("Device.DeviceInfo.Manufacturer")
	modelName, _ := utils.GetDataModelValue("Device.DeviceInfo.ModelName")
	serialNumber, _ := utils.GetDataModelValue("Device.DeviceInfo.SerialNumber")
	id := fmt.Sprintf("urn:dev:os:%s-%s", manufacturerOUI, serialNumber)
	productClass, _ := utils.GetDataModelValue("Device.DeviceInfo.ProductClass")
	softwareVersion, _ := utils.GetDataModelValue("Device.DeviceInfo.SoftwareVersion")
	hardwareVersion, _ := utils.GetDataModelValue("Device.DeviceInfo.HardwareVersion")
	statusLastChange, _ := utils.GetDataModelValue("Device.DeviceInfo.UpTime")

	// Construct DeviceInfo response
	reply := &pb.DeviceInfo{
		Type:             "HOMEGATEWAY",
		Id:               id,
		ManufacturerOUI:  manufacturerOUI,
		Manufacturer:     manufacturer,
		ModelName:        modelName,
		SerialNumber:     serialNumber,
		ProductClass:     productClass,
		SoftwareVersion:  softwareVersion,
		HardwareVersion:  hardwareVersion,
		StatusLastChange: statusLastChange,
	}

	return reply, nil
}

func main() {
	flag.Parse()

	// Create a listener using the specified IP address and port
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *ip, *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a new gRPC server
	s := grpc.NewServer()

	// Register the HomeGateway service implementation
	pb.RegisterHomeGatewayServer(s, &homeGateway{})

	log.Printf("server listening at %s:%d", *ip, *port)

	// Start serving gRPC requests
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
