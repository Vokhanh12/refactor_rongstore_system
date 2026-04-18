package mappers

import (
	authzrs "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/iam/authz/v1/resources"
	authzuc "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/usecases"
)

// ============================================================
// PROTO → COMMAND / QUERY
// ============================================================

func PermissionMutateRequestToBatch(req *authzrs.PermissionMutateRequest) authzuc.PermissionMutationBatch {
	return authzuc.PermissionMutationBatch{}
}

func RolePermissionMutateRequestToBatch(req *authzrs.RolePermissionMutateRequest) authzuc.RolePermissionMutationBatch {
	return authzuc.RolePermissionMutationBatch{}
}

func RoleViewRequestToBatch(req *authzrs.RoleViewRequest) authzuc.RoleViewBatch {
	return authzuc.RoleViewBatch{}
}

func PermissionViewRequestToBatch(req *authzrs.PermissionViewRequest) authzuc.PermissionViewBatch {
	return authzuc.PermissionViewBatch{}
}

func RolePermissionViewRequestToBatch(req *authzrs.RolePermissionViewRequest) authzuc.RolePermissionViewBatch {
	return authzuc.RolePermissionViewBatch{}
}
