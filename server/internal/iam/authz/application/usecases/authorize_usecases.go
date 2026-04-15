package usecases

import (
	"context"
	"strings"
	"time"

	"github.com/google/uuid"
	domerrs "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/errors"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/apperrors"

	com "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/command"
	cs "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/caches"
	rs "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/repositories"
	vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"
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
		return deny(&domerrs.AUTHORIZATION_ROLE_REQUIRED)
	}

	if cmd.Resource == "" || cmd.Action == "" {
		return deny(&domerrs.AUTHORIZATION_RESOURCE_OR_ACTION_REQUIRED)
	}

	targetKey := cmd.Resource + ":" + cmd.Action

	rolePermKeysMap := make(map[string][]vo.ResourceAction, len(cmd.Roles))
	var roleKeysMissCache []vo.RoleRef

	for _, role := range cmd.Roles {

		parts := strings.Split(role, ":")
		roleCode := parts[0]
		scopeID, err := uuid.Parse(parts[1])
		if err != nil {
			panic(err)
		}

		roleKey, errdetails := vo.NewRoleRef(&scopeID, roleCode)
		if errdetails != nil {
			return nil, aerrs.New(domerrs.UNAUTHORIZED, aerrs.WithAppendErrorDetails(errdetails))
		}

		perms, err := u.rolePermissionCache.GetPermissions(ctx, *roleKey)
		if err != nil {
			return nil, err
		}

		if len(perms) > 0 {
			rolePermKeysMap[role] = perms
		} else {
			roleKeysMissCache = append(roleKeysMissCache, *roleKey)
		}
	}

	if len(roleKeysMissCache) > 0 {
		rolePermissions, err := u.rolePermissionRepository.FindAllByRoleRefs(ctx, roleKeysMissCache)
		if err != nil {
			return nil, err
		}

		for _, rp := range rolePermissions {
			if rp.Role.IsElevated() {
				return allow()
			}

			rolePermKeysMap[rp.Role.RoleRef().String()] = append(rolePermKeysMap[rp.Role.RoleRef().String()], rp.Permission.ResourceAction())
		}

		for _, roleRef := range roleKeysMissCache {

			perms := rolePermKeysMap[roleRef.String()]
			u.rolePermissionCache.SetPermissions(ctx, roleRef, perms, permissionCacheTTL)
		}
	}

	globalPerms := make(map[string]struct{})
	for _, perms := range rolePermKeysMap {
		for _, p := range perms {
			globalPerms[p.String()] = struct{}{}
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
