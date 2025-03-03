package repository

import (
	"context"
	"errors"
	"fmt"

	"short_tail/internal/domain/url/models"

	badger "github.com/dgraph-io/badger/v4"
)

func (r *Repository) Get(_ context.Context, alias string) (*models.URL, error) {
	url := models.URL{
		Alias: alias,
	}

	err := r.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(url.Alias))
		if err != nil {
			return fmt.Errorf("txn.Get: %w", err)
		}

		return item.Value(func(val []byte) error {
			url.Original = string(val)

			return nil
		})
	})
	if err != nil {
		if errors.Is(err, badger.ErrKeyNotFound) {
			return &url, models.ErrNotFound
		}

		return nil, fmt.Errorf("r.db.View: %w", err)
	}

	return &url, nil
}
