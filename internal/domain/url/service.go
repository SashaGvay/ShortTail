package url

import (
	"context"
	"short_tail/internal/domain/url/models"
)

type Service interface {
	Short(ctx context.Context, url string) (*models.URL, error)
	UnShort(ctx context.Context, alias string) (*models.URL, error)
}
