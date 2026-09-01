package repository

import "../repository/context"

type CredentialRepository interface {
	FindByIdentifier(
		ctx context.Context,
		identifier string,
	) (*Credential, error)
}
