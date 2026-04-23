package valueobjects

import (
	"strconv"

	core "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/errors"
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

func NewResourceAction(it ValidatedResourceAction) *ResourceAction {
	return &ResourceAction{
		resource: it.ResourceAction.resource,
		action:   it.ResourceAction.action,
	}
}

func NewResourceActions(values []string) ([]ResourceAction, []aerrs.AppErrorDetail) {
	var (
		result  = make([]ResourceAction, 0, len(values))
		details []aerrs.AppErrorDetail
	)

	for i, v := range values {
		ra, errs := NewResourceAction(v)
		if len(errs) > 0 {
			for _, d := range errs {
				d.Field = "resourceActions[" + strconv.Itoa(i) + "]"
				details = append(details, d)
			}
			continue
		}

		result = append(result, *ra)
	}

	if len(details) > 0 {
		return nil, details
	}

	return result, nil
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
// VALIDATION
// ============================================================

func (r *ResourceAction) validate() *aerrs.AppError {
	var details []aerrs.AppErrorDetail

	if r.resource == "" {
		details = append(details, aerrs.NewDetail(
			core.REASON_VAL_REQUIRED,
			aerrs.WithField("resource"),
			aerrs.WithMessageDetail("resource is required"),
		))
	}

	if r.action == "" {
		details = append(details, aerrs.NewDetail(
			core.REASON_VAL_REQUIRED,
			aerrs.WithField("action"),
			aerrs.WithMessageDetail("action is required"),
		))
	}

	return aerrs.New(core.VALIDATION_FAILED, aerrs.WithAppendErrorDetails(details))
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
