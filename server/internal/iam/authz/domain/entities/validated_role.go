package entities

type ValidatedRole struct {
	Role
	isValidated bool
}

func (vp *ValidatedRole) IsValid() bool {
	return vp.isValidated
}

func NewValidatedRole(Role *Role) (*ValidatedRole, error) {
	if err := Role.validate(); err != nil {
		return nil, err
	}

	return &ValidatedRole{
		Role:        *Role,
		isValidated: true,
	}, nil
}
