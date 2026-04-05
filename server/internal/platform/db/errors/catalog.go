package errors

import (
	"server/internal/iam/domain"
)

var StoreOwnerDbError = DBError{
	ConflictCatalog: domain.STORE_OWNER_CONFLICT,
	InvalidCatalog:  domain.STORE_OWNER_INVALID,
	NotfoundCatalog: domain.STORE_OWNER_NOT_FOUND,
	PersistCatalog:  domain.STORE_OWNER_PERSIST_FAILED,
}

var RoleDbError = DBError{
	ConflictCatalog: domain.ROLE_CONFLICT,
	InvalidCatalog:  domain.ROLE_INVALID,
	NotfoundCatalog: domain.ROLE_NOT_FOUND,
	PersistCatalog:  domain.ROLE_PERSIST_FAILED,
}

var PermissionDbError = DBError{
	ConflictCatalog: domain.PERMISSION_CONFLICT,
	InvalidCatalog:  domain.PERMISSION_INVALID,
	NotfoundCatalog: domain.PERMISSION_NOT_FOUND,
	PersistCatalog:  domain.PERMISSION_PERSIST_FAILED,
}
