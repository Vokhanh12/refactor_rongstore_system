package valueobject

type TokenType string

const (
	TypeAccess  TokenType = "ACCESS"
	TypeRefresh TokenType = "REFRESH"
)

func (t TokenType) IsValid() bool {
	switch t {
	case TypeAccess, TypeRefresh:
		return true
	}
	return false
}
