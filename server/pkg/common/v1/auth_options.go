package commonv1

type ResourceOptions struct {
	Resource           string
	Action             string
	ResourceIDField    string
	CheckOwnership     bool
	RequireTenant      bool
	SkipAuth           bool
	RequiredPermission string
}

const (
	OptionResource           = "resource"
	OptionAction             = "action"
	OptionResourceIDField    = "resource_id_field"
	OptionCheckOwnership     = "check_ownership"
	OptionRequireTenant      = "require_tenant"
	OptionSkipAuth           = "skip_auth"
	OptionRequiredPermission = "required_permission"
)
