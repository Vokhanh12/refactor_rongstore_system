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

	resourceVal := proto.GetExtension(opts, comv1.E_Resource)
	resource, ok := resourceVal.(string)
	if !ok {
		return nil, fmt.Errorf("missing or invalid E_Resource extension")
	}

	actionVal := proto.GetExtension(opts, comv1.E_Action)
	action, ok := actionVal.(string)
	if !ok {
		return nil, fmt.Errorf("missing or invalid E_Action extension")
	}

	resourceFieldVal := proto.GetExtension(opts, comv1.E_ResourceIdField)
	resourceField, _ := resourceFieldVal.(string)

	requireTenantVal := proto.GetExtension(opts, comv1.E_RequireTenant)
	requireTenant, _ := requireTenantVal.(bool)

	return &AuthOptions{
		Resource:        resource,
		Action:          action,
		ResourceIDField: resourceField,
		RequireTenant:   requireTenant,
	}, nil
}

func extractResourceID(req proto.Message, fieldName string) string {
	if fieldName == "" {
		return ""
	}

	msg := req.ProtoReflect()
	field := msg.Descriptor().Fields().ByName(protoreflect.Name(fieldName))
	if field == nil {
		return ""
	}

	value := msg.Get(field)

	switch value.Interface().(type) {
	case string:
		return value.String()
	case protoreflect.EnumNumber:
		return fmt.Sprintf("%d", value.Enum())
	case int32, int64, uint32, uint64:
		return fmt.Sprintf("%v", value.Interface())
	default:
		return value.String()
	}
}
