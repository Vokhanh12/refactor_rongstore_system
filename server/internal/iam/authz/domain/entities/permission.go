package entities

import (
	"github.com/google/uuid"

	"github.com/vokhanh12/refactor-rongstore-system/server/internal/core/validator"
	vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

// ============================================================
// ENTITY
// ============================================================

type Permission struct {
	id   uuid.UUID
	code string

	name        *string
	description *string

	resourceAction vo.ResourceAction

	isActive bool
}

// ============================================================
// PAYLOADS
// ============================================================

type PermissionPayload struct {
	Code string

	Name        *string
	Description *string

	Resource string
	Action   string

	IsActive bool
}

type NewPermissionParams struct {
	PermissionPayload
}

type RestorePermissionParams struct {
	ID uuid.UUID
	PermissionPayload
}

// ============================================================
// CONSTRUCTOR (domain - validate)
// ============================================================

func NewPermission(
	it NewPermissionParams,
) (*Permission, *aerrs.AppError) {

	v := validator.New()

	err := v.
		Required("Code", it.Code).
		MaxLen("Code", it.Code, 100).
		Required("Resource", it.Resource).
		Required("Action", it.Action).
		Err()

	if err != nil {
		return nil, err
	}

	resourceAction, appErr := vo.NewResourceAction(
		it.Resource,
		it.Action,
	)

	if appErr != nil {
		return nil, appErr
	}

	permission := newPermissionFromPayload(
		uuid.Must(uuid.NewV7()),
		it.PermissionPayload,
		resourceAction,
	)

	return &permission, nil
}

// ============================================================
// RESTORE (persistence - trust data)
// ============================================================

func RestorePermission(it RestorePermissionParams) Permission {

	resourceAction := vo.RestoreResourceAction(
		it.Resource,
		it.Action,
	)

	return newPermissionFromPayload(
		it.ID,
		it.PermissionPayload,
		resourceAction,
	)
}

// ============================================================
// PRIVATE FACTORY
// ============================================================

func newPermissionFromPayload(
	id uuid.UUID,
	payload PermissionPayload,
	resourceAction vo.ResourceAction,
) Permission {

	return Permission{
		id:             id,
		code:           payload.Code,
		name:           payload.Name,
		description:    payload.Description,
		resourceAction: resourceAction,
		isActive:       payload.IsActive,
	}
}

// ============================================================
// DOMAIN METHODS
// ============================================================

func (p Permission) Key() string {
	return p.resourceAction.Resource() +
		":" +
		p.resourceAction.Action()
}

func (p Permission) Match(
	resource,
	action string,
) bool {
	return p.resourceAction.Resource() == resource &&
		p.resourceAction.Action() == action
}

// ============================================================
// GETTERS
// ============================================================

func (p Permission) ID() uuid.UUID {
	return p.id
}

func (p Permission) Code() string {
	return p.code
}

func (p Permission) Name() *string {
	return p.name
}

func (p Permission) Description() *string {
	return p.description
}

func (p Permission) ResourceAction() vo.ResourceAction {
	return p.resourceAction
}

func (p Permission) IsActive() bool {
	return p.isActive
}
