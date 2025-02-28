package db

import (
	"github.com/tubfuzzy/banraiphisan-reservation/config"
)

type DB struct {
}

func NewDB(conf config.DatabaseConfig) (*DB, error) {

	return &DB{}, nil
}
