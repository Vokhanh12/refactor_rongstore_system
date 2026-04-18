package mappers

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func MarshalAny(msg proto.Message) (*anypb.Any, error) {
	return anypb.New(msg)
}

func Must(msg string) {
	panic("mapper fatal: " + msg)
}
