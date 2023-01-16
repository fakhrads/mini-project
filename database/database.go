package database

import (
	"database/sql"
	"fmt"

	"github.com/gobuffalo/packr/v2"
	migrate "github.com/rubenv/sql-migrate"
)

var (
	DbConntection *sql.DB
)

func DbMigrate(dbParams *sql.DB) {
	migrations := &migrate.PackrMigrationSource{
		Box: packr.New("migrations", "./sql_migrations"),
	}

	n, errs := migrate.Exec(dbParams, "postgres", migrations, migrate.Up)
	DbConntection = dbParams
	if errs != nil {
		panic(errs)
	}

	fmt.Println("Applied ", n, "migrations!")
}
