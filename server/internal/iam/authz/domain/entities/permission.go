package entities

import (
	"github.com/google/uuid"
	vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"
)

type Permission struct {
	id   uuid.UUID
	code string

	name        *string
	description *string

	resourceAction vo.ResourceAction

	isActive bool
}

func NewPermissionFromPersistence(
	id uuid.UUID,
	code string,
	name *string,
	description *string,
	resource string,
	action string,
	isActive bool,
) Permission {

	ra := vo.NewResourceActionFromPersistence(resource, action)

	return Permission{
		id:             id,
		code:           code,
		name:           name,
		description:    description,
		resourceAction: ra,
		isActive:       isActive,
	}
}

func (p Permission) ResourceAction() vo.ResourceAction {
	return p.resourceAction
}

func (p Permission) Key() string {
	return p.resourceAction.Resource() + ":" + p.resourceAction.Action()
}

func (p Permission) Match(resource, action string) bool {
	return p.resourceAction.Resource() == resource && p.resourceAction.Action() == action
}
