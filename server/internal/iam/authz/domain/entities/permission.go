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

type NewPermissionParams struct {
	ID   uuid.UUID
	Code string

	Name        *string
	Description *string

	Resource string
	Action   string

	IsActive bool
}

func NewPermissionFromPersistence(p NewPermissionParams) Permission {
	return Permission{
		id:             p.ID,
		code:           p.Code,
		name:           p.Name,
		description:    p.Description,
		resourceAction: vo.NewResourceActionFromPersistence(p.Resource, p.Action),
		isActive:       p.IsActive,
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
