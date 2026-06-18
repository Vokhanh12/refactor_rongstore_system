package wire

import (
	"context"
	"server/internal/iam/infrastructure/client"
	"server/internal/iam/infrastructure/db/postgres"
	"server/pkg/config"

	iamgrpc "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/adapter/grpc/handler"
	cache "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/cache/redis"
	db "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/postgres"
)

type Infra struct {
	Postgres db.PostgresDB
	Redis    cache.RedisCache
}

type RootDeps struct {
	Infra *Infra

	IAM IamDeps
	FNB FnBDeps
	HR  HrDeps
}

type IamDeps struct {
	AuthHandler  *iamgrpc.AuthHandler
	AuthzHandler *iamgrpc.AuthzHandler
}

type FnBDeps struct{}

type HrDeps struct{}

func Initialize(
	ctx context.Context,
	cfg *config.Config,
) *RootDeps {

	infra := &Infra{
		Postgres: postgres.InitPostgresDB(ctx, cfg),
		Redis:    cache.InitRedisCache(ctx, cfg),
	}

	return &RootDeps{
		Infra: infra,

		IAM: InitializeIAM(ctx, cfg, infra),
		FNB: InitializeFNB(ctx, cfg, infra),
		HR:  InitializeHR(ctx, cfg, infra),
	}
}

func InitializeIAM(
	ctx context.Context,
	cfg *config.Config,
	infra *Infra,
) IamDeps {

	keycloak := client.InitKeycloakClient(ctx, cfg)

	commandRepo := gorm.NewGormStoreOwnerCommandRepository(
		infra.Postgres,
	)

	queryRepo := sqlc.NewSqlcStoreOwnerQueryRepository(
		infra.Postgres,
	)

	loginUC := auth.NewLoginUsecase(
		keycloak,
	)

	handshakeUC := auth.NewHandshakeUsecase(
		infra.Redis,
	)

	authzHandler := grpc.NewAuthzHandler(
		loginUC,
		handshakeUC,
	)

	return IamDeps{
		AuthzHandler: authzHandler,
	}
}
