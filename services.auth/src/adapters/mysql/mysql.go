package mysql

import (
	"github.com/jmoiron/sqlx"
	mysql_user "github.com/mixarchitecture/arch/auth/src/adapters/mysql/user"
	"github.com/mixarchitecture/arch/auth/src/domain/user"
)

type MySQL interface {
	NewUser(userFactory user.Factory, sql *sqlx.DB) user.Repository
}

type mySql struct{}

func New() MySQL {
	return &mySql{}
}

func (m *mySql) NewUser(userFactory user.Factory, sql *sqlx.DB) user.Repository {
	return mysql_user.New(userFactory, sql)
}
