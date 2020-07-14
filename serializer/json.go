package serializer

import (
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

//ProtobufToJSON transforms protobuf message to JSON string
func ProtobufToJSON(message proto.Message) (string, error) {
	marshaller := jsonpb.Marshaler{
		EmitDefaults: true,
		EnumsAsInts:  false,
		Indent:       "  ",
		OrigName:     false,
	}

	return marshaller.MarshalToString(message)
}
