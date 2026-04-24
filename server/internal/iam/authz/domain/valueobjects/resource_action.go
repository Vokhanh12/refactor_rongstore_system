package valueobjects

import (
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/core/validator"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

// ============================================================
// VALUE OBJECT
// ============================================================

type ResourceAction struct {
	resource string
	action   string
}

// ============================================================
// CONSTRUCTOR (domain - có validate)
// ============================================================

func NewResourceAction(resource string, action string) (*ResourceAction, *aerrs.AppError) {

	v := validator.New().
		Required("resource", resource).
		Required("action", action)

	if err := v.Err(); err != nil {
		return nil, err
	}

	return &ResourceAction{
		resource: resource,
		action:   action,
	}, nil
}

// ============================================================
// RESTORE (persistence - trust data)
// ============================================================

func RestoreResourceAction(resource, action string) ResourceAction {
	return ResourceAction{
		resource: resource,
		action:   action,
	}
}

// ============================================================
// GETTERS
// ============================================================

func (r ResourceAction) Resource() string { return r.resource }
func (r ResourceAction) Action() string   { return r.action }

// ============================================================
// UTILS
// ============================================================

func (r ResourceAction) String() string {
	return r.resource + ":" + r.action
}
