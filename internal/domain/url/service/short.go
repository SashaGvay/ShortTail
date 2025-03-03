package service

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"

	"short_tail/internal/domain/url/models"
)

func (s *Service) Short(ctx context.Context, url string) (*models.URL, error) {
	dto := &models.URL{
		Original: s.removeScheme(url),
		Alias:    s.generateAlias(url),
	}

	err := s.Repository.Put(ctx, dto)
	if err != nil {
		return nil, fmt.Errorf("s.Repository.Put: %w", err)
	}

	return dto, err
}

func (s *Service) removeScheme(url string) string {
	url = strings.TrimPrefix(url, "http://")
	url = strings.TrimPrefix(url, "https://")
	return url
}

func (s *Service) generateAlias(url string) string {
	hash := sha256.Sum256([]byte(url))
	alias := base64.URLEncoding.EncodeToString(hash[:])[:8]

	return alias
}
