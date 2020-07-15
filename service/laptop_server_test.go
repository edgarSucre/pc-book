package service_test

import (
	"context"
	"pcbook/pb"
	"pcbook/sample"
	"pcbook/service"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestServerCreateLaptop(t *testing.T) {
	t.Parallel()

	laptopNoID := sample.NewLaptop()
	laptopNoID.Id = ""

	laptopInvalidID := sample.NewLaptop()
	laptopInvalidID.Id = "invalid-uuid"

	laptopDuplicatedID := sample.NewLaptop()
	storeDuplicatedID := service.NewInMemoryLaptopStore()
	err := storeDuplicatedID.Save(laptopDuplicatedID)
	if err != nil {
		t.Fail()
	}

	testCases := []struct {
		name   string
		laptop *pb.Laptop
		store  service.LaptopStore
		code   codes.Code
	}{
		{
			name:   "success_with_id",
			laptop: sample.NewLaptop(),
			store:  service.NewInMemoryLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "success_no_id",
			laptop: laptopNoID,
			store:  service.NewInMemoryLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "failure_invalid_id",
			laptop: laptopInvalidID,
			store:  service.NewInMemoryLaptopStore(),
			code:   codes.InvalidArgument,
		},
		{
			name:   "failure_duplicated_id",
			laptop: laptopDuplicatedID,
			store:  storeDuplicatedID,
			code:   codes.AlreadyExists,
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			req := &pb.CreateLaptopRequest{
				Laptop: tc.laptop,
			}

			server := service.NewLaptopServer(tc.store)
			resp, err := server.CreateLaptop(context.Background(), req)
			if tc.code == codes.OK {
				if err != nil && resp == nil && resp.Id == tc.laptop.Id {
					t.Fail()
				}
			} else {
				if err == nil && resp != nil {
					t.Fail()
				}

				status, ok := status.FromError(err)
				if !ok && status.Code() != tc.code {
					t.Fail()
				}
			}
		})
	}
}
