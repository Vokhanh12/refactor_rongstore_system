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
