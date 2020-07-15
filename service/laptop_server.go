package service

import (
	"context"
	"errors"
	"log"
	"pcbook/pb"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//LaptopServer provides laptop services
type LaptopServer struct {
	Store LaptopStore
}

//NewLaptopServer returns a new LaptopServer
func NewLaptopServer(store LaptopStore) *LaptopServer {
	return &LaptopServer{store}
}

//CreateLaptop creates a new Laptop
func (server *LaptopServer) CreateLaptop(
	ctx context.Context,
	req *pb.CreateLaptopRequest,
) (*pb.CreateLaptopResponse, error) {

	laptop := req.GetLaptop()
	log.Printf("received a create laptop request with id: %s", laptop.Id)

	if len(laptop.Id) > 0 {
		_, err := uuid.Parse(laptop.Id)
		if err != nil {
			//using grpc status, remember this is not HTTP :)
			return nil, status.Errorf(codes.InvalidArgument, "laptop ID is not a valid UUID: %v", err)
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "cannot generate a new laptop ID: %v", err)
		}

		laptop.Id = id.String()
	}

	//save laptop to the store
	err := server.Store.Save(laptop)
	if err != nil {
		code := codes.Internal
		if errors.Is(err, ErrAlreadyExists) {
			code = codes.AlreadyExists
		}
		return nil, status.Errorf(code, "cannot store a new laptop: %v", err)
	}

	resp := &pb.CreateLaptopResponse{
		Id: laptop.Id,
	}

	log.Printf("saved laptop with id: %s", laptop.Id)

	return resp, nil
}
