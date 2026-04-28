package mappers

import (
	"time"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func MustMarshalAny(msg proto.Message) *anypb.Any {
	if msg == nil {
		panic("MustMarshalAny: msg is nil")
	}
	a, err := anypb.New(msg)
	if err != nil {
		panic(err)
	}
	return a
}

func ToProtoTime(t time.Time) *timestamppb.Timestamp {
	if t.IsZero() {
		return nil
	}
	return timestamppb.New(t)
}

func FromProtoTime(t *timestamppb.Timestamp) time.Time {
	if t == nil {
		return time.Time{}
	}
	return t.AsTime()
}

func Must(msg string) {
	panic("mapper fatal: " + msg)
}
