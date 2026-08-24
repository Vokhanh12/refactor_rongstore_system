package postgres

import (
	repos "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/repositories"
	pg "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/postgres"
)

var _ repos.CredentialRepository = (*CredentialRepository)(nil)

type CredentialRepository struct {
	dba *pg.DbAdapter
}
