package serializer_test

import (
	"pcbook/pb"
	"pcbook/sample"
	"pcbook/serializer"
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestWriteProtobufToBinaryFile(t *testing.T) {
	binaryFile := "../tmp/laptop.bin"

	laptop1 := sample.NewLaptop()

	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)

	if err != nil {
		t.Fatal(err.Error())
	}

	laptop2 := &pb.Laptop{}
	err = serializer.ReadProtobufFromBinaryFile(binaryFile, laptop2)
	if err != nil {
		t.Fatal(err.Error())
	}

	if !proto.Equal(laptop1, laptop2) {
		t.Fatal("The laptop messages don't match")
	}
}
