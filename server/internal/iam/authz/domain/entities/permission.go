package entities

import vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"

type Permission struct {
	id   string
	code string

	name        *string
	description *string

	resourceAction vo.ResourceAction

	isActive bool
}

func (p Permission) Key() string {
	return p.resourceAction.Resource() + ":" + p.resourceAction.Action()
}

func (p Permission) Match(resource, action string) bool {
	return p.resourceAction.Resource() == resource && p.resourceAction.Action() == action
}
