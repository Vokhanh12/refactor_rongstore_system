package valueobject

type TokenType string

const (
	TypeAccess  TokenType = "access"
	TypeRefresh TokenType = "refresh"
)

func (t TokenType) IsValid() bool {
	switch t {
	case TypeAccess, TypeRefresh:
		return true
	}
	return false
}
