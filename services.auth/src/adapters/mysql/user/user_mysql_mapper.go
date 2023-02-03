package user

import (
	"strings"

	"github.com/mixarchitecture/arch/auth/src/adapters/mysql/user/entity"
	"github.com/mixarchitecture/arch/auth/src/domain/user"
	sqb_go "gitlab.com/ssibrahimbas/sqb.go"
)

type Mapper interface {
	ToEntityMap(u *user.User) *sqb_go.M
}

type mapper struct{}

func NewMapper() Mapper {
	return &mapper{}
}

func (m *mapper) ToEntityMap(u *user.User) *sqb_go.M {
	return &sqb_go.M{
		entity.Fields.UUID:      u.UUID,
		entity.Fields.Email:     u.Email,
		entity.Fields.Roles:     strings.Join(u.Roles, ","),
		entity.Fields.Password:  u.Password,
		entity.Fields.IsActive:  u.IsActive,
		entity.Fields.CreatedAt: u.CreatedAt,
		entity.Fields.UpdatedAt: u.UpdatedAt,
	}
}
