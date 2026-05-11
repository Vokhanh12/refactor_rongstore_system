package entities

import (
	"github.com/google/uuid"

	cren "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/entities"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/core/validator"
	vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

// ============================================================
// ENTITY
// ============================================================

type Permission struct {
	cren.BaseEntity

	id uuid.UUID

	code string

	name        *string
	description *string

	resourceAction vo.ResourceAction

	isActive bool
}

// ============================================================
// PAYLOAD
// ============================================================

type PermissionPayload struct {
	Code string

	Name        *string
	Description *string

	Resource string
	Action   string

	IsActive bool
}

// ============================================================
// CONSTRUCTOR
// ============================================================

func NewPermission(
	payload PermissionPayload,
) (*Permission, *aerrs.AppError) {

	if err := validatePermissionPayload(payload); err != nil {
		return nil, err
	}

	resourceAction, err := vo.NewResourceAction(
		payload.Resource,
		payload.Action,
	)

	if err != nil {
		return nil, err
	}

	permission := &Permission{
		id: uuid.Must(uuid.NewV7()),

		code: payload.Code,

		name:        payload.Name,
		description: payload.Description,

		resourceAction: resourceAction,

		isActive: payload.IsActive,
	}

	return permission, nil
}

// ============================================================
// RESTORE
// ============================================================

func RestorePermission(
	id uuid.UUID,
	payload PermissionPayload,
) *Permission {

	resourceAction := vo.RestoreResourceAction(
		payload.Resource,
		payload.Action,
	)

	return &Permission{
		id: id,

		code: payload.Code,

		name:        payload.Name,
		description: payload.Description,

		resourceAction: resourceAction,

		isActive: payload.IsActive,
	}
}

// ============================================================
// VALIDATION
// ============================================================

func validatePermissionPayload(
	payload PermissionPayload,
) *aerrs.AppError {

	v := validator.New()

	return v.
		Required("code", payload.Code).
		MaxLen("code", payload.Code, 100).
		Required("resource", payload.Resource).
		Required("action", payload.Action).
		Err()
}

// ============================================================
// DOMAIN METHODS
// ============================================================

func (p *Permission) Activate() {
	p.isActive = true
}

func (p *Permission) Deactivate() {
	p.isActive = false
}

func (p Permission) Key() string {
	return p.resourceAction.Resource() +
		":" +
		p.resourceAction.Action()
}

func (p Permission) Match(
	resource string,
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
