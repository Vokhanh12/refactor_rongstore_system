package grpc

import (
	"fmt"

	comv1 "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/common/v1"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type AuthOptions struct {
	Resource        string
	Action          string
	ResourceIDField string
	RequireTenant   bool
}

func extractAuthOptions(req proto.Message) (*AuthOptions, error) {

	md := req.ProtoReflect().Descriptor()

	opts := md.Options()

	resource := proto.GetExtension(opts, comv1.E_Resource).(string)
	action := proto.GetExtension(opts, comv1.E_Action).(string)
	resourceField := proto.GetExtension(opts, comv1.E_ResourceIdField).(string)

	return &AuthOptions{
		Resource:        resource,
		Action:          action,
		ResourceIDField: resourceField,
	}, nil
}

func extractResourceID(
	req proto.Message,
	fieldName string,
) (string, error) {

	msg := req.ProtoReflect()

	field := msg.Descriptor().Fields().
		ByName(protoreflect.Name(fieldName))

	if field == nil {
		return "", fmt.Errorf("field not found")
	}

	value := msg.Get(field)

	return value.String(), nil
}
