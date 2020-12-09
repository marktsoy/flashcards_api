package apiserver

import (
	"database/sql"
	"net/http"

	"github.com/marktsoy/flashcards_api/internal/app/store/sqlstore"
)

// Start the api
func Start(conf *Config) error {

	db, err := initDB(conf.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	store := sqlstore.New(db)

	srv := newServer(conf, store)

	return http.ListenAndServe(conf.BindAddr, srv)
}

func initDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
