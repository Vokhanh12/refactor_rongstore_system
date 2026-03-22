package commonv1

// ResourceOptions chứa các annotation/proto extension
type ResourceOptions struct {
	Resource           string
	Action             string
	ResourceIDField    string
	CheckOwnership     bool
	RequireTenant      bool
	SkipAuth           bool
	RequiredPermission string
}

// Key constants cho annotation nếu muốn attach vào struct gRPC message
const (
	OptionResource           = "resource"
	OptionAction             = "action"
	OptionResourceIDField    = "resource_id_field"
	OptionCheckOwnership     = "check_ownership"
	OptionRequireTenant      = "require_tenant"
	OptionSkipAuth           = "skip_auth"
	OptionRequiredPermission = "required_permission"
)
