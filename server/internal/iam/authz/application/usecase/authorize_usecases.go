package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/errors"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"

	com "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/command"
	ce "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/cache"
	re "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/repositories"
)

type AuthorizeUsecase struct {
	rolePermissionCache      ce.RolePermissionCache
	rolePermissionRepository re.RolePermissionRepository
}

func NewAuthorizeUsecase(rpCache ce.RolePermissionCache, rpRepository re.RolePermissionRepository) *AuthorizeUsecase {
	return &AuthorizeUsecase{
		rolePermissionCache:      rpCache,
		rolePermissionRepository: rpRepository,
	}
}

func (u *AuthorizeUsecase) Execute(ctx context.Context, cmd com.AuthorizeCommand) (*com.AuthorizeCommandResult, *aerrs.AppError) {

	if len(cmd.RoleCodes) == 0 {
		return &com.AuthorizeCommandResult{Allowed: false},
			errors.New(iamErrors.AUTHORIZATION_ROLE_REQUIRED)
	}

	if cmd.ResourceCheck == "" || cmd.ActionCheck == "" {
		return &com.AuthorizeCommandResult{Allowed: false},
			errors.New(iamErrors.AUTHORIZATION_RESOURCE_OR_ACTION_REQUIRED)
	}

	permKey := fmt.Sprintf("%s:%s", cmd.ResourceCheck, cmd.ActionCheck)

	for _, code := range cmd.RoleCodes {
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
