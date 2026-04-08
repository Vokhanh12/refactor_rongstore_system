package testhelpers

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	db "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/sqlc"
)

const (
	dbUser     = "admin"
	dbPassword = "rongstore@180320"
	dbName     = "rongstore"
)

const (
	iamAuthDir     = "../sql/iam/auth/schema"
	iamAuthzDir    = "../sql/iam/authz/schema"
	iamIdentityDir = "../sql/iam/identity/schema"
)

type PostgresTestContainer struct {
	Conn    *pgx.Conn
	Queries *db.Queries
}

func SetupTestDB(t *testing.T) *PostgresTestContainer {
	ctx := context.Background()

	dsn := fmt.Sprintf(
		"postgres://%s:%s@100.114.31.30:5432/%s?sslmode=disable",
		dbUser,
		dbPassword,
		dbName,
	)

	err := applySchema(ctx, dsn)
	require.NoError(t, err)

	conn, err := pgx.Connect(ctx, dsn)
	require.NoError(t, err)

	queries := db.New(conn)

	return &PostgresTestContainer{
		Conn:    conn,
		Queries: queries,
	}
}

func (p *PostgresTestContainer) Close(t *testing.T) {
	ctx := context.Background()

	if p.Conn != nil {
		err := p.Conn.Close(ctx)
		require.NoError(t, err, "Failed to close database connection")
	}
}

func (p *PostgresTestContainer) WithTx(t *testing.T, fn func(q *db.Queries)) {
	ctx := context.Background()

	tx, err := p.Conn.Begin(ctx)
	require.NoError(t, err)

	q := db.New(tx)

	fn(q)

	_ = tx.Rollback(ctx)
}

func applySchema(ctx context.Context, dsn string) error {

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return err
	}

	db := stdlib.OpenDBFromPool(pool)

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	projectRoot := cwd
	for {
		if _, err := os.Stat(filepath.Join(projectRoot, "go.mod")); err == nil {
			break
		}
		parent := filepath.Dir(projectRoot)
		if parent == projectRoot {
			return fmt.Errorf("go.mod not found")
		}
		projectRoot = parent
	}

	migrationsDir := filepath.Join(projectRoot, iamAuthzDir)

	if err := goose.Up(db, migrationsDir); err != nil {
		return err
	}

	return nil
}
