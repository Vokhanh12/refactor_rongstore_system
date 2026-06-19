package wire

import (
	"context"
	"server/internal/iam/infrastructure/client"
	"server/pkg/config"

	iamgrpc "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/adapter/grpc/handler"
)

type IamDeps struct {
	AuthHandler  *iamgrpc.AuthHandler
	AuthzHandler *iamgrpc.AuthzHandler
}

type FnBDeps struct{}

type HrDeps struct{}

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

func InitializeHR(
	ctx context.Context,
	cfg *config.Config,
	infra *Infra,
) HrDeps {
	panic("none")
}

func InitializeFNB(
	ctx context.Context,
	cfg *config.Config,
	infra *Infra,
) FnBDeps {
	panic("none")
}
