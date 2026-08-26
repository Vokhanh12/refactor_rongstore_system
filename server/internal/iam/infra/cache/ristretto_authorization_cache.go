package cache

import (
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/caches"
)

var _ caches.AuthorizationCache = (*RistrettoAuthorizationCache)(nil)

type RistrettoAuthorizationCache struct{}
