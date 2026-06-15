package usecases

import (
	"context"
	"time"

	merrs "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/errors"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"

	com "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/command"
	repos "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/query"
	cs "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/caches"
	vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"
)

const permissionCacheTTL = 15 * time.Minute

type AuthorizeUsecase struct {
	cache cs.AuthorizationCache
	query repos.AuthorizationQueryRepository
}

func NewAuthorizeUsecase(c cs.AuthorizationCache, q repos.AuthorizationQueryRepository) *AuthorizeUsecase {
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
	if len(cmd.RoleKeyStrs) == 0 {
		return deny(&merrs.ROLE_REQUIRED)
	}

	if cmd.Resource == "" || cmd.Action == "" {
		return deny(&merrs.RESOURCE_OR_ACTION_REQUIRED)
	}

	targetKey := cmd.Resource + ":" + cmd.Action

	rasMap := make(map[string][]vo.ResourceAction, len(cmd.RoleKeyStrs))
	var rksMissCache []vo.RoleKey

	for _, roleKeyStr := range cmd.RoleKeyStrs {

		roleKey, err := vo.ParseRoleKey(roleKeyStr)
		if err != nil {
			return nil, err
		}

		rasCache, aerr := u.cache.GetResourceActionByRoleKey(ctx, roleKey)
		if aerr != nil {
			return nil, aerr
		}

		if len(rasCache) > 0 {
			rasMap[roleKeyStr] = rasCache
			continue
		}

		rksMissCache = append(rksMissCache, roleKey)
	}

	if len(rksMissCache) > 0 {

		authzGrants, err := u.query.ListGrantsByRoleKeys(ctx, rksMissCache)
		if err != nil {
			return nil, err
		}

		for _, ag := range authzGrants {

			roKeystr := ag.RoleKey.String()

			rasMap[roKeystr] = append(
				rasMap[roKeystr],
				ag.ResourceAction,
			)
		}

		for _, RoleKey := range rksMissCache {
			ras := rasMap[RoleKey.String()]
			u.cache.SetResourceActionByRoleKey(ctx, RoleKey, ras, permissionCacheTTL)
		}
	}

	globalPerms := make(map[string]struct{})

	for _, perms := range rasMap {
		for _, p := range perms {
			globalPerms[p.String()] = struct{}{}
		}
	}

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
