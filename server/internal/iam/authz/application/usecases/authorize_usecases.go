package usecases

import (
	"context"
	"strings"
	"time"

	"github.com/google/uuid"
	core "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/errors"
	merrs "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/errors"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"

	com "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/command"
	q "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/query"
	cs "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/caches"
	vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"
)

const permissionCacheTTL = 15 * time.Minute

type AuthorizeUsecase struct {
	cache cs.AuthorizationCache
	query q.AuthorizationQuery
}

func NewAuthorizeUsecase(c cs.AuthorizationCache, q q.AuthorizationQuery) *AuthorizeUsecase {
	return &AuthorizeUsecase{
		cache: c,
		query: q,
	}
}

func (u *AuthorizeUsecase) Execute(
	ctx context.Context,
	cmd com.AuthorizeCommand,
) (*com.AuthorizeCommandResult, *aerrs.AppError) {

	// ---------- Validate input ----------
	if len(cmd.Roles) == 0 {
		return deny(&merrs.ROLE_REQUIRED)
	}

	if cmd.Resource == "" || cmd.Action == "" {
		return deny(&merrs.RESOURCE_OR_ACTION_REQUIRED)
	}

	targetKey := cmd.Resource + ":" + cmd.Action

	rolePermKeysMap := make(map[string][]vo.ResourceAction, len(cmd.Roles))
	var RoleKeysMissCache []vo.RoleKey

	// ---------- Parse + cache check ----------
	for _, role := range cmd.Roles {

		parts := strings.Split(role, ":")
		if len(parts) != 2 {
			return nil, aerrs.New(core.STRING_SPLIT_INVALID)
		}

		roleCode := parts[0]

		scopeID, err := uuid.Parse(parts[1])
		if err != nil {
			return nil, aerrs.New(core.UUID_INVALID, aerrs.WithCauseDetail(err))
		}

		roleKey, aerr := vo.NewRoleKey(&scopeID, roleCode)
		if aerr != nil {
			return nil, aerr
		}

		// cache
		perms, aerr := u.cache.GetResourceActionByRoleKey(ctx, roleKey)
		if err != nil {
			return nil, aerr
		}

		if len(perms) > 0 {
			rolePermKeysMap[role] = perms
			continue
		}

		RoleKeysMissCache = append(RoleKeysMissCache, roleKey)
	}

	// ---------- Load from DB ----------
	if len(RoleKeysMissCache) > 0 {

		authzGrants, err := u.query.ListGrantsByRoleKeys(ctx, RoleKeysMissCache)
		if err != nil {
			return nil, err
		}

		for _, ag := range authzGrants {

			if ag.Match(ag.Resource, ag.Action) {
				return allow()
			}

			key := ag.RoleKey.String()

			rolePermKeysMap[key] = append(
				rolePermKeysMap[key],
				ag.Permission.ResourceAction(),
			)
		}

		// cache back
		for _, RoleKey := range RoleKeysMissCache {
			perms := rolePermKeysMap[RoleKey.String()]
			u.rolePermissionCache.SetPermissions(ctx, RoleKey, perms, permissionCacheTTL)
		}
	}

	// ---------- Flatten permissions ----------
	globalPerms := make(map[string]struct{})

	for _, perms := range rolePermKeysMap {
		for _, p := range perms {
			globalPerms[p.String()] = struct{}{}
		}
	}

	// ---------- Check permission ----------
	if _, ok := globalPerms[targetKey]; ok {
		return allow()
	}

	return deny(nil)
}
func allow() (*com.AuthorizeCommandResult, *aerrs.AppError) {
	return &com.AuthorizeCommandResult{Allowed: true}, nil
}

func deny(err *aerrs.AppError) (*com.AuthorizeCommandResult, *aerrs.AppError) {
	return &com.AuthorizeCommandResult{Allowed: false}, err
}
