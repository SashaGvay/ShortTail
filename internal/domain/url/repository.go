package url

import (
	"context"
	"short_tail/internal/domain/url/models"
)

type Repository interface {
	Put(ctx context.Context, url *models.URL) error
	Get(ctx context.Context, alias string) (*models.URL, error)
}
