package storage

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

type StorageEngine interface {
	Store(vendor string, payload []byte) error
}

type sqliteEngine struct {
	db *sql.DB
}

func NewFromViperSettings() (StorageEngine, error) {
	engine := viper.GetString("storage.engine")

	switch engine {
	case "sqlite3":
		return newSqliteEngine(viper.GetString("storage.connection_string"))
	case "":
		return nil, errors.New("[storage.engine] needs to be defined in the config file, not assuming any default value")
	default:
		return nil, errors.New(fmt.Sprintf("%s is not a supported storage option", engine))
	}
}

func newSqliteEngine(file string) (sqliteEngine, error) {
	db, err := sql.Open("sqlite3", file)

	return sqliteEngine{
		db: db,
	}, err
}

func (s sqliteEngine) Store(vendor string, payload []byte) error {
	_, err := s.db.Exec(
		"INSERT INTO webhooks (vendor, raw_body) VALUES ($1, $2)",
		vendor,
		payload,
	)

	if err != nil {
		return err
	} else {
		return nil
	}
}
