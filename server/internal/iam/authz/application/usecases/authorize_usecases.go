package usecases

import (
	"context"
	"time"

	domerrs "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/errors"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/apperrors"

	com "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/command"
	cs "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/caches"
	rs "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/repositories"
)

const permissionCacheTTL = 15 * time.Minute

type AuthorizeUsecase struct {
	rolePermissionCache      cs.RolePermissionCache
	rolePermissionRepository rs.RolePermissionRepository
}

func NewAuthorizeUsecase(
	rpCache cs.RolePermissionCache,
	rpRepository rs.RolePermissionRepository,
) *AuthorizeUsecase {
	return &AuthorizeUsecase{
		rolePermissionCache:      rpCache,
		rolePermissionRepository: rpRepository,
	}
}

func (u *AuthorizeUsecase) Execute(
	ctx context.Context,
	cmd com.AuthorizeCommand,
) (*com.AuthorizeCommandResult, *aerrs.AppError) {

	if len(cmd.Roles) == 0 {
		return deny(aerrs.New(
			domerrs.AUTHORIZATION_REQUIRED, aerrs.WithAppendErrorDetail(*aerrs.NewDetail(
				domerrs.REASON_REQUIRED, aerrs.WithField("role"),
				aerrs.WithMessageDetail("role is required")))))
	}

	if cmd.Resource == "" || cmd.Action == "" {
		return deny(aerrs.New(
			domerrs.AUTHORIZATION_REQUIRED,
			aerrs.WithAppendErrorDetail(*aerrs.NewDetail(domerrs.REASON_REQUIRED, aerrs.WithField("resouce"), aerrs.WithMessageDetail("resouce is required"))),
			aerrs.WithAppendErrorDetail(*aerrs.NewDetail(domerrs.REASON_REQUIRED, aerrs.WithField("action"), aerrs.WithMessageDetail("action is required"))),
		))
	}

	targetKey := cmd.Resource + ":" + cmd.Action

	rolePermsMap := make(map[string][]string, len(cmd.Roles))
	var rolesMissCache []string

	for _, role := range cmd.Roles {
		perms, err := u.rolePermissionCache.GetPermissions(ctx, role)
		if err != nil {
			return nil, err
		}

		if len(perms) > 0 {
			rolePermsMap[role] = perms
		} else {
			rolesMissCache = append(rolesMissCache, role)
		}
	}

	if len(rolesMissCache) > 0 {
		rolePermissions, err := u.rolePermissionRepository.FindAllByRoles(ctx, rolesMissCache)
		if err != nil {
			return nil, err
		}

		for _, rp := range rolePermissions {
			if rp.Role.IsElevated() {
				return allow()
			}
			rolePermsMap[rp.Role.Code] = append(rolePermsMap[rp.Role.Code], rp.Permission.Key())
		}

		for _, role := range rolesMissCache {
			perms := rolePermsMap[role]
			u.rolePermissionCache.SetPermissions(ctx, role, perms, permissionCacheTTL)
		}
	}

	globalPerms := make(map[string]struct{})
	for _, perms := range rolePermsMap {
		for _, p := range perms {
			globalPerms[p] = struct{}{}
		}
	}

	if _, ok := globalPerms[targetKey]; ok {
		return allow()
	}

	return &com.AuthorizeCommandResult{Allowed: false}, nil
}

func allow() (*com.AuthorizeCommandResult, *aerrs.AppError) {
	return &com.AuthorizeCommandResult{Allowed: true}, nil
}

func deny(err *aerrs.AppError) (*com.AuthorizeCommandResult, *aerrs.AppError) {
	return &com.AuthorizeCommandResult{Allowed: false}, err
}
