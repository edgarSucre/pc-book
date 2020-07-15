package main

import (
	"context"
	"flag"
	"log"
	"pcbook/pb"
	"pcbook/sample"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	serverAddres := flag.String("address", "", "The server address")
	flag.Parse()
	log.Printf("dial server: %s", *serverAddres)

	conn, err := grpc.Dial(*serverAddres, grpc.WithInsecure())
	if err != nil {
		log.Fatal("cannot dial server:", err)
	}

	laptopClient := pb.NewLaptopServiceClient(conn)
	laptop := sample.NewLaptop()
	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	res, err := laptopClient.CreateLaptop(context.Background(), req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			log.Print("laptop already exists")
		} else {
			log.Fatal("cannot create laptop: ", err)
		}
		return
	}

	log.Printf("created laptop with id: %s", res.Id)
}
