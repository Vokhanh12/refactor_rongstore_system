package usecases

import (
	"context"
	"fmt"
	"time"

	domerrs "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/errors"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"

	com "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/command"
	cs "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/caches"
	rs "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/repositories"
)

type AuthorizeUsecase struct {
	rolePermissionCache      cs.RolePermissionCache
	rolePermissionRepository rs.RolePermissionRepository
}

func NewAuthorizeUsecase(rpCache cs.RolePermissionCache, rpRepository rs.RolePermissionRepository) *AuthorizeUsecase {
	return &AuthorizeUsecase{
		rolePermissionCache:      rpCache,
		rolePermissionRepository: rpRepository,
	}
}

func (u *AuthorizeUsecase) Execute(ctx context.Context, cmd com.AuthorizeCommand) (*com.AuthorizeCommandResult, *aerrs.AppError) {

	if len(cmd.Roles) == 0 {
		return &com.AuthorizeCommandResult{Allowed: false},
			aerrs.New(domerrs.AUTHORIZATION_ROLE_REQUIRED)
	}

	if cmd.Resource == "" || cmd.Action == "" {
		return &com.AuthorizeCommandResult{Allowed: false},
			aerrs.New(domerrs.AUTHORIZATION_RESOURCE_OR_ACTION_REQUIRED)
	}

	permKey := fmt.Sprintf("%s:%s", cmd.Resource, cmd.Action)

	for _, code := range cmd.Roles {
		perms, found, err := u.rolePermissionCache.GetPermissions(ctx, code)
		if err != nil {
			return nil, err
		}

		if !found {
			rolePermissions, err := u.rolePermissionRepository.FindAllByRoleCode(code)
			if err != nil {
				return nil, err
			}

			perms = make([]string, 0, len(rolePermissions))
			for _, rp := range rolePermissions {

				if rp.Role.IsSuper {
					return &com.AuthorizeCommandResult{Allowed: true}, nil
				}

				key := fmt.Sprintf("%s:%s", rp.Permission.Resource, rp.Permission.Action)
				perms = append(perms, key)
			}

			u.rolePermissionCache.SetPermissions(ctx, code, perms, 1000*time.Second)
		}

		for _, p := range perms {
			if p == permKey {
				return &com.AuthorizeCommandResult{Allowed: true}, nil
			}
		}
	}

	return &com.AuthorizeCommandResult{Allowed: false}, nil
}
