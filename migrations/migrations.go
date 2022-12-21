package migrations

import (
	"database/sql"
	"embed"

	migrate "github.com/rubenv/sql-migrate"
	_ "modernc.org/sqlite"
)

//go:embed *.sql
var migrationsFiles embed.FS

func Run(dbPath string) error {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return err
	}
	defer db.Close()

	migrations := &migrate.EmbedFileSystemMigrationSource{
		FileSystem: migrationsFiles,
		Root:       ".",
	}

	if _, err := migrate.Exec(db, "sqlite3", migrations, migrate.Up); err != nil {
		return err
	}

	return nil
}
