package entity

import (
	"github.com/jmoiron/sqlx"
	mysql_migration "github.com/mixarchitecture/arch/shared/migration/mysql"
)

type migration struct {
	database string
}

func NewMigration() mysql_migration.MigrationItem {
	return &migration{
		database: "boilerplate.users",
	}
}

func (m *migration) Up(db *sqlx.DB) error {
	query := `CREATE TABLE IF NOT EXISTS boilerplate.users (
        uuid VARCHAR(255) NOT NULL PRIMARY KEY,
		password VARCHAR(255) NOT NULL,
        email VARCHAR(255) NOT NULL,
		roles TEXT NOT NULL,
        is_active BOOLEAN NOT NULL,
        created_at DATETIME NOT NULL,
        updated_at DATETIME NOT NULL
    )`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (m *migration) Down(db *sqlx.DB) error {
	query := `DROP TABLE boilerplate.users`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (m *migration) GetDatabase() string {
	return m.database
}
