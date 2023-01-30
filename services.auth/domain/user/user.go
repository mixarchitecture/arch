package user

import "github.com/mixarchitecture/arch/shared/jwt"

type User struct {
	UUID      string
	Email     string
	Roles     []string
	Password  []byte
	IsActive  bool
	CreatedAt string
	UpdatedAt string
}

func (u *User) SetPassword(password []byte) {
	u.Password = password
}

func (u *User) CleanPassword() {
	u.Password = nil
}

func (u *User) ToJwtClaims() *jwt.UserClaim {
	return &jwt.UserClaim{
		UUID:      u.UUID,
		Email:     u.Email,
		Roles:     u.Roles,
		ExpiresIn: 3600,
	}
}
