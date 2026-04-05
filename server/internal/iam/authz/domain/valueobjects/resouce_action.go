package valueobjects

import (
	"strconv"
	"strings"

	domain "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/errors"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/apperrors"
)

type ResourceAction struct {
	resource string
	action   string
}

func NewResourceActionFromPersistence(resource, action string) ResourceAction {
	return ResourceAction{
		resource: resource,
		action:   action,
	}
}

func NewResourceAction(value string) (*ResourceAction, []aerrs.AppErrorDetail) {
	var details []aerrs.AppErrorDetail

	if strings.TrimSpace(value) == "" {
		details = append(details, *aerrs.NewDetail(
			domain.REASON_REQUIRED,
			aerrs.WithField("ResourceAction"),
			aerrs.WithMessageDetail("ResourceAction is required"),
		))
		return nil, details
	}

	parts := strings.Split(value, ":")
	if len(parts) != 2 {
		details = append(details, *aerrs.NewDetail(
			domain.REASON_INVALID_FORMAT,
			aerrs.WithField("ResourceAction"),
			aerrs.WithMessageDetail("ResourceAction must be in format resource:action"),
		))
		return nil, details
	}

	resource := strings.TrimSpace(parts[0])
	action := strings.TrimSpace(parts[1])

	if resource == "" || action == "" {
		details = append(details, *aerrs.NewDetail(
			domain.REASON_INVALID_FORMAT,
			aerrs.WithField("permission"),
			aerrs.WithMessageDetail("resource or action is empty"),
		))
		return nil, details
	}

	return &ResourceAction{
		resource: resource,
		action:   action,
	}, nil
}

func NewResourceActions(values []string) ([]ResourceAction, []aerrs.AppErrorDetail) {
	var (
		result  = make([]ResourceAction, 0, len(values))
		details []aerrs.AppErrorDetail
	)

	for i, v := range values {
		pk, errs := NewResourceAction(v)
		if len(errs) > 0 {
			for _, d := range errs {
				d.Field = "permissions[" + strconv.Itoa(i) + "]"
				details = append(details, d)
			}
			continue
		}

		result = append(result, *pk)
	}

	if len(details) > 0 {
		return nil, details
	}

	return result, nil
}

func (p ResourceAction) Action() string {
	return p.action
}

func (p ResourceAction) Resource() string {
	return p.resource
}

func (p ResourceAction) String() string {
	return p.resource + ":" + p.action
}
