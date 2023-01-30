package mysql

import (
	mysql_example "github.com/mixarchitecture/arch/example/src/adapters/mysql/example"
	"github.com/mixarchitecture/arch/example/src/domain/example"

	"github.com/jmoiron/sqlx"
)

type MySQL interface {
	NewExample(db *sqlx.DB, exampleFactory example.Factory) example.Repository
}

type mysql struct{}

func New() MySQL {
	return &mysql{}
}

func (m *mysql) NewExample(db *sqlx.DB, exampleFactory example.Factory) example.Repository {
	return mysql_example.NewExampleRepo(db, exampleFactory)
}
