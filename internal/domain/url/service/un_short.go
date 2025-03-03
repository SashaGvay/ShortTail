package service

import (
	"context"
	"fmt"

	"short_tail/internal/domain/url/models"
)

func (s *Service) UnShort(ctx context.Context, alias string) (*models.URL, error) {
	dto, err := s.Repository.Get(ctx, alias)
	if err != nil {
		return nil, fmt.Errorf("s.Repository.Get: %w", err)
	}

	return dto, err
}
