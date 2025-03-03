package service

import "short_tail/internal/domain/url"

type Service struct {
	Repository url.Repository
}

func New(repository url.Repository) *Service {
	return &Service{
		Repository: repository,
	}
}
