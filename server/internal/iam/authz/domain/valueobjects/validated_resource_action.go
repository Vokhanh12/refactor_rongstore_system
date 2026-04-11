package valueobjects

type ValidatedResourceAction struct {
	ResourceAction
	isValidated bool
}

func (vp *ValidatedResourceAction) IsValid() bool {
	return vp.isValidated
}

func NewValidatedResourceAction(ResourceAction *ResourceAction) (*ValidatedResourceAction, error) {
	if err := ResourceAction.validate(); err != nil {
		return nil, err
	}

	return &ValidatedResourceAction{
		ResourceAction: *ResourceAction,
		isValidated:    true,
	}, nil
}
