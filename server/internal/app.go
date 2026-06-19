package wire

import (
	"context"
	"net"
	"server/pkg/config"

	wire "github.com/vokhanh12/refactor-rongstore-system/server/internal"
	"google.golang.org/grpc"
)

type Application struct {
	cfg  *config.Config
	deps *wire.RootDeps
}

func New(
	cfg *config.Config,
	deps *wire.RootDeps,
) *Application {
	return &Application{
		cfg:  cfg,
		deps: deps,
	}
}

func (a *Application) Run(
	ctx context.Context,
) error {

	server := grpc.NewServer()

	pb.RegisterAuthzServiceServer(
		server,
		a.deps.IAM.AuthzHandler,
	)

	lis, err := net.Listen(
		"tcp",
		a.cfg.GRPC.Address,
	)
	if err != nil {
		return err
	}

	return server.Serve(lis)
}
