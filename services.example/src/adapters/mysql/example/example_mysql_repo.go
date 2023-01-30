package mysql_example

import (
	mysql_migration "github.com/mixarchitecture/arch/shared/migration/mysql"

	"github.com/mixarchitecture/arch/example/src/adapters/mysql/example/entity"
	"github.com/mixarchitecture/arch/example/src/config"
	"github.com/mixarchitecture/arch/example/src/domain/example"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type exampleRepo struct {
	db             *sqlx.DB
	exampleFactory example.Factory
	mapper         MySqlExampleMapper
}

func NewExampleRepo(db *sqlx.DB, exampleFactory example.Factory) example.Repository {
	if db == nil {
		panic("db is nil")
	}
	if exampleFactory.IsZero() {
		panic("exampleFactory is zero")
	}
	mapper := NewMySqlExampleMapper(exampleFactory)
	return &exampleRepo{
		db:             db,
		exampleFactory: exampleFactory,
		mapper:         mapper,
	}
}

func New(cnf config.MySQLExample) (*sqlx.DB, error) {
	config := mysql.NewConfig()
	config.Net = "tcp"
	config.Addr = cnf.Address
	config.User = cnf.Username
	config.Passwd = cnf.Password
	config.DBName = cnf.Database
	config.ParseTime = true

	db, err := sqlx.Connect("mysql", config.FormatDSN())
	if err != nil {
		return nil, errors.Wrap(err, "failed to open mysql connection")
	}

	migration := mysql_migration.New(db)
	migration.Add(entity.NewExampleMigration())
	err = migration.Up()
	if err != nil {
		return nil, errors.Wrap(err, "failed to migrate mysql")
	}
	return db, nil
}
