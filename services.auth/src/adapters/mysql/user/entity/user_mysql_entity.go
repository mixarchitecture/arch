package entity

import (
	"strings"

	"github.com/mixarchitecture/arch/auth/src/domain/user"
)

type MySQLUser struct {
	UUID      string `db:"uuid"`
	Email     string `db:"email"`
	Roles     string `db:"roles"`
	Password  []byte `db:"password"`
	IsActive  bool   `db:"is_active"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

type fields struct {
	UUID      string
	Email     string
	Roles     string
	Password  string
	IsActive  string
	CreatedAt string
	UpdatedAt string

	Table string
}

var Fields = fields{
	UUID:      "uuid",
	Password:  "password",
	Roles:     "roles",
	Email:     "email",
	IsActive:  "is_active",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	Table:     "users",
}

func (m *MySQLUser) ToUser() *user.User {
	return &user.User{
		UUID:     m.UUID,
		Email:    m.Email,
		Password: m.Password,
		Roles:    strings.Split(m.Roles, ","),
		IsActive: m.IsActive,
	}
}
