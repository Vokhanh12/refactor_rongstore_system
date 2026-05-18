package services

type TokenSigner interface {
	Sign(userID string, orgID string, unitID string, roles []string) (string, error)
}
