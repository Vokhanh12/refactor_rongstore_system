package errors

import (
	domain "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/errors"
)

var RoleDbError = DBError{
	Conflict: domain.ROLE_CONFLICT,
	Invalid:  domain.ROLE_INVALID,
	NotFound: domain.ROLE_NOT_FOUND,
	Internal: domain.ROLE_PERSIST_FAILED,
}
