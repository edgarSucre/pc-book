package serializer_test

import (
	"pcbook/sample"
	"pcbook/serializer"
	"testing"
)

func TestWriteProtobufToBinaryFile(t *testing.T) {
	binaryFile := "../tmp/laptop.bin"

	laptop1 := sample.NewLaptop()

	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)

	if err != nil {
		t.Fatal(err.Error())
	}
}
