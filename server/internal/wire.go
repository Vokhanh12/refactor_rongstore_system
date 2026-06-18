//go:build wireinject
// +build wireinject

package wire

import (
	"context"
	"server/internal/iam/infrastructure/cache"
	"server/internal/iam/infrastructure/client"
	"server/internal/iam/infrastructure/db/postgres"
	"server/internal/iam/infrastructure/eventbus"
	"server/internal/iam/infrastructure/storage/r2"
	"server/pkg/config"
)

func Initialize(ctx context.Context, cfg *config.Config) *RootDeps {
	infra := &Infra{
		DB:            postgres.InitPostgresDB(ctx, cfg),
		Redis:         cache.InitRedisCache(ctx, cfg),
		Keycloak:      client.InitKeycloakClient(ctx, cfg),
		ObjectStorage: r2.InitR2Storage(ctx, cfg),
		EventBus:      eventbus.InitRabbitMQEventBus(cfg),
	}

	iam := InitializeIamHandler(
		infra.DB,
		infra.Redis,
		infra.Keycloak,
		infra.ObjectStorage,
		infra.EventBus,
	)

	return &RootDeps{
		Infra: infra,
		IAM:   iam,
	}
}

func InitializeIamHandler(infra *Infra) IamDeps {
	commandRepo := gorm.NewGormStoreOwnerCommandRepository(infra.DB)
	queryRepo := sqlc.NewSqlcStoreOwnerQueryRepository(infra.DB)

	loginUC := auth.NewLoginUsecase(infra.Keycloak)
	handshakeUC := auth.NewHandshakeUsecase(infra.Redis)

	mutateUC := store_owner.NewMutateUsecase(
		commandRepo,
		infra.ObjectStorage,
	)

	viewUC := store_owner.NewViewUsecase(queryRepo)

	handler := grpc.NewIamHandler(
		loginUC,
		handshakeUC,
		mutateUC,
		viewUC,
	)

	return IamDeps{
		Handler: handler,
	}
}
