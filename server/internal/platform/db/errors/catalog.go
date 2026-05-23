package errors

import (
	domain "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/errors"
)

var RoleDbError = DBError{
	Conflict: domain.ROLE_ALREADY_EXISTS,
	Invalid:  domain.ROLE_INVALID,
	NotFound: domain.ROLE_NOT_FOUND,
}
