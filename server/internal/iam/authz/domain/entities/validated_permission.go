package entities

type ValidatedPermission struct {
	Permission
	isValidated bool
}

func (vp *ValidatedPermission) IsValid() bool {
	return vp.isValidated
}

func NewValidatedPermission(Permission *Permission) (*ValidatedPermission, error) {
	if err := Permission.validate(); err != nil {
		return nil, err
	}

	return &ValidatedPermission{
		Permission:  *Permission,
		isValidated: true,
	}, nil
}
