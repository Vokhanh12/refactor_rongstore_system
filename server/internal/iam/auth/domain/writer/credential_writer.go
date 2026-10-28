package writer

import "context"

type CredentialWriter interface {
	FindByIdentifier(
		ctx context.Context,
		identifier string,
	) (*Credential, error)
}
