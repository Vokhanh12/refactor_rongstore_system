package presenters

import (
	authzrs "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/iam/authz/v1/resources"
	corem "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/adapter/assemblers"
	cmd "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/command"
	"google.golang.org/protobuf/types/known/anypb"
)

type CreateRolePresenter struct {
	Result *cmd.CreateRoleCommandResult
}

func (p CreateRolePresenter) ToProto() *anypb.Any {
	pb := &authzrs.CreateResult{
		RoleResult: &authzrs.RoleResult{
			Id:          p.Result.Role.Id,
			Name:        p.Result.Role.Name,
			Description: p.Result.Role.Description,
			CreateAt:    corem.ToProtoTime(p.Result.Role.CreatedAt),
			UpdateAt:    corem.ToProtoTime(p.Result.Role.UpdatedAt),
		},
	}
	return corem.MustMarshalAny(pb)
}

type UpdateRolePresenter struct {
	Result *cmd.UpdateRoleCommandResult
}

func (p UpdateRolePresenter) ToProto() *anypb.Any {
	pb := &authzrs.UpdateResult{
		RoleResult: &authzrs.RoleResult{
			Id:          p.Result.Role.Id,
			Name:        p.Result.Role.Name,
			Description: p.Result.Role.Description,
			CreateAt:    corem.ToProtoTime(p.Result.Role.CreatedAt),
			UpdateAt:    corem.ToProtoTime(p.Result.Role.UpdatedAt),
		},
	}
	return corem.MustMarshalAny(pb)
}

type DeleteRolePresenter struct {
	Result *cmd.DeleteRoleCommandResult
}

func (p DeleteRolePresenter) ToProto() *anypb.Any {
	pb := &authzrs.DeleteResult{}
	return corem.MustMarshalAny(pb)
}
