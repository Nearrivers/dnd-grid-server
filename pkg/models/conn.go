package models

import (
	"database/sql"
	"embed"
	"log"
	"os"

	"github.com/Nearrivers/dnd-grid-server/pkg/models/repository"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
)

//go:embed schema/*.sql
var embedMigrations embed.FS

func createDbFile() {
	_, err := os.Stat("core.db")
	if err == nil || !os.IsNotExist(err) {
		return
	}

	f, errCreate := os.Create("core.db")
	if errCreate != nil {
		log.Fatalf("Impossible de créer le fichier de base de donnée: %v\n", errCreate)
	}

	f.Close()
}

func ConnectToDb() *repository.Queries {
	createDbFile()

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("sqlite"); err != nil {
		log.Fatalf("Impossible d'initialiser le dialecte de la bdd")
	}

	db, err := sql.Open("sqlite3", "core.db")
	if err != nil {
		log.Fatalf("Impossible de se connecter à la base de donnée: %v\n", err)
	}

	if err := goose.Up(db, "schema"); err != nil {
		log.Fatalf("Impossible d'effectuer les migrations: %v\n", err)
	}

	return repository.New(db)
}
