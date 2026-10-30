package entity

type Credential struct {
	userID       string
	passwordHash string
	authzVersion int
}

func NewCredential(userID string, passwordHash string) Credential {
	return Credential{
		userID:       userID,
		passwordHash: passwordHash,
		authzVersion: 1, // Mặc định version quyền là 1 khi mới tạo tài khoản
	}
}

// Dùng cho trường hợp hydrate lại dữ liệu từ database (qua sqlc)
func UnmarshalCredential(userID string, passwordHash string, authzVersion int) Credential {
	return Credential{
		userID:       userID,
		passwordHash: passwordHash,
		authzVersion: authzVersion,
	}
}

func (c Credential) UserID() string {
	return c.userID
}

func (c Credential) PasswordHash() string {
	return c.passwordHash
}

func (c Credential) AuthzVersion() int {
	return c.authzVersion
}

// Hành vi đổi mật khẩu (Domain Behavior)
func (c *Credential) ChangePassword(newPasswordHash string) {
	c.passwordHash = newPasswordHash
	c.authzVersion++ // Tăng authzVersion lên để vô hiệu hóa các token cũ nếu cần
}
