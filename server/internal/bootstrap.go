package wire

import (
	"context"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/postgres"
	cfg "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/config"

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

func Initialize(
	ctx context.Context,
	cfg *cfg.Config,
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
