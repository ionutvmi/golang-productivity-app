package migrations

import (
	"database/sql"
	"embed"
	"log"

	migrate "github.com/rubenv/sql-migrate"
	_ "modernc.org/sqlite"
)

//go:embed *.sql
var migrationsFiles embed.FS

func MustRun(dbPath string) {
	db, err := sql.Open("sqlite", dbPath)

	if err != nil {
		log.Fatalf("Failed to open database connection: %s", err.Error())
	}
	defer db.Close()

	migrations := &migrate.EmbedFileSystemMigrationSource{
		FileSystem: migrationsFiles,
		Root:       ".",
	}

	if _, err := migrate.Exec(db, "sqlite3", migrations, migrate.Up); err != nil {
		log.Fatalf("Failed to execute migrations: %s", err.Error())
	}
}
