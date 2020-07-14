package serializer_test

import (
	"pcbook/pb"
	"pcbook/sample"
	"pcbook/serializer"
	"testing"

	"github.com/golang/protobuf/proto"
)

func TestWriteProtobufToBinaryFile(t *testing.T) {
	binaryFile := "../tmp/laptop.bin"
	jsonFile := "../tmp/laptop.json"

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

	err = serializer.WriteProtobufToJSONFile(laptop2, jsonFile)
	if err != nil {
		t.Fatal(err.Error())
	}
}
