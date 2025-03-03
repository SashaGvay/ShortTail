package repository

import (
	"context"
	"fmt"

	badger "github.com/dgraph-io/badger/v4"

	"short_tail/internal/domain/url/models"
)

func (r *Repository) Put(_ context.Context, url *models.URL) error {
	err := r.db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(url.Alias), []byte(url.Original))
		if err != nil {
			return fmt.Errorf("txn.Set: %w", err)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("r.db.Update: %w", err)
	}

	return nil
}
