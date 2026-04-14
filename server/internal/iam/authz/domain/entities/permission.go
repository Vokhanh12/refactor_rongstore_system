package entities

import (
	"github.com/google/uuid"

	vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"
)

// ============================================================
// ENTITY DOMAIN
// ============================================================

type Permission struct {
	id   uuid.UUID
	code string

	name        *string
	description *string

	resourceAction vo.ResourceAction

	isActive bool
}

func (p *Permission) validate() error {
	return nil
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

// ============================================================
// ENTITY DATABASE
// ============================================================

type NewPermissionParams struct {
	ID   uuid.UUID
	Code string

	Name        *string
	Description *string

	Resource string
	Action   string

	IsActive bool
}

func NewPermissionFromPersistence(pp NewPermissionParams, rap vo.NewResourceActionParms) Permission {
	return Permission{
		id:             pp.ID,
		code:           pp.Code,
		name:           pp.Name,
		description:    pp.Description,
		resourceAction: vo.NewResourceActionFromPersistence(rap),
		isActive:       pp.IsActive,
	}
}
