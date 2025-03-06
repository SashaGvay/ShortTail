package service

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"

	qrcode "github.com/skip2/go-qrcode"

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

	qrCode, err := s.generateQrCodeAsBase64(fmt.Sprintf("%s/%s", s.HttpAddr, dto.Alias))
	if err != nil {
		return nil, fmt.Errorf("s.generateQrCodeAsBase64: %w", err)
	}

	if qrCode != nil {
		dto.QrCode = *qrCode
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

func (s *Service) generateQrCodeAsBase64(url string) (*string, error) {
	img, err := qrcode.Encode(url, qrcode.Medium, 256)
	if err != nil {
		return nil, fmt.Errorf("qrcode.Encode: %w", err)
	}

	asBase64 := "data:image/png;base64," + base64.StdEncoding.EncodeToString(img)

	return &asBase64, nil
}
