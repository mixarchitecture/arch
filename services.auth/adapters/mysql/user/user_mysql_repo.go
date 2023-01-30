package user

import (
	"github.com/jmoiron/sqlx"
	"github.com/mixarchitecture/arch/auth/domain/user"
)

type repo struct {
	db          *sqlx.DB
	userFactory user.Factory
	mapper      Mapper
}

func New(userFactory user.Factory, db *sqlx.DB) user.Repository {
	validate(userFactory, db)
	return &repo{
		userFactory: userFactory,
		db:          db,
		mapper:      NewMapper(),
	}
}

func validate(userFactory user.Factory, db *sqlx.DB) {
	if userFactory.IsZero() {
		panic("exampleFactory is zero")
	}
	if db == nil {
		panic("db is nil")
	}
}
