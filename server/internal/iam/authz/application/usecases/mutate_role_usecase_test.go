package usecases

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	core "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/usecase"
	c "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/command"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/entities"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/repositories"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/apperrors"
)

var _ repositories.RoleRepository = (*MockRoleRepository)(nil)

type MockRoleRepository struct {
	CreateFn  func(ctx context.Context, role *entities.Role) (*entities.Role, *aerrs.AppError)
	UpdateFn  func(ctx context.Context, role *entities.Role) (*entities.Role, *aerrs.AppError)
	DeleteFn  func(ctx context.Context, id string) *aerrs.AppError
	GetByIDnF func(ctx context.Context, id string) (*entities.Role, *aerrs.AppError)
}

func (m *MockRoleRepository) Create(ctx context.Context, role *entities.Role) (*entities.Role, *aerrs.AppError) {
	if m.CreateFn != nil {
		return m.CreateFn(ctx, role)
	}
	return role, nil
}

func (m *MockRoleRepository) Delete(ctx context.Context, id uuid.UUID) *aerrs.AppError {
	if m.DeleteFn != nil {
		return m.DeleteFn(ctx, id)
	}
	return nil
}

func (m *MockRoleRepository) Update(ctx context.Context, role *entities.Role) (*entities.Role, *aerrs.AppError) {
	if m.UpdateFn != nil {
		return m.UpdateFn(ctx, role)
	}
	return role, nil
}

func TestMutateRoleUsecase_Create_Success(t *testing.T) {
	ctx := context.Background()

	mockRepo := &MockRoleRepository{
		CreateFn: func(ctx context.Context, role *entities.Role) (*entities.Role, *aerrs.AppError) {
			return role, nil
		},
	}

	uc := NewMutateRoleUsecase(mockRepo)

	batch := RoleMutationBatch{
		Items: []core.Operation[RoleMutation]{
			{
				OpID: "1",
				Payload: RoleMutation{
					Create: &c.CreateRoleCommand{
						Code: "ADMIN",
						Name: "Admin",
					},
				},
			},
		},
	}

	result := uc.Execute(ctx, batch)

	require.Len(t, result.Items, 1)
	assert.Nil(t, result.Items[0].Error)
	assert.Equal(t, "1", result.Items[0].OpID)
}
