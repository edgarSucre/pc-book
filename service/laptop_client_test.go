package service_test

import (
	"context"
	"net"
	"pcbook/pb"
	"pcbook/sample"
	"pcbook/service"
	"testing"

	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
)

func TestClientCreateLaptop(t *testing.T) {
	t.Parallel()

	laptopServer, serverAddress := startTestLaptopServer(t)
	laptopClient := newTestLaptopClient(t, serverAddress)

	laptop := sample.NewLaptop()
	expectedId := laptop.Id

	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	resp, err := laptopClient.CreateLaptop(context.Background(), req)
	if err != nil || resp.Id != expectedId {
		t.Fail()
	}

	storedLaptop, err := laptopServer.Store.Find(resp.Id)
	if err != nil || storedLaptop == nil {
		t.Fail()
	}

	if !proto.Equal(laptop, storedLaptop) {
		t.Fail()
	}

}

func startTestLaptopServer(t *testing.T) (*service.LaptopServer, string) {
	laptopServer := service.NewLaptopServer(service.NewInMemoryLaptopStore())

	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	//listen on any random port available
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fail()
	}

	//blocking code, needs to run on a separate goroutine
	go grpcServer.Serve(listener)

	//getting the port from the listener
	return laptopServer, listener.Addr().String()
}

func newTestLaptopClient(t *testing.T, serverAddress string) pb.LaptopServiceClient {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		t.Fail()
	}
	return pb.NewLaptopServiceClient(conn)
}
