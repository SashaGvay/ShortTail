package repository

import badger "github.com/dgraph-io/badger/v4"

type Repository struct {
	db *badger.DB
}

func New(db *badger.DB) *Repository {
	return &Repository{
		db: db,
	}
}
