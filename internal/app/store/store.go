package store

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Store struct {
	config          *Config
	db              *sql.DB
	trablRepository *TrablRepository
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Open ..
func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) Services() *TrablRepository {
	if s.trablRepository != nil {
		return s.trablRepository
	}

	s.trablRepository = &TrablRepository{
		store: s,
	}

	return s.trablRepository
}

// store.Service().Select()
