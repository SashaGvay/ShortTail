package service

import "short_tail/internal/domain/url"

type Service struct {
	Repository url.Repository
	HttpAddr   string
}

func New(repository url.Repository, httpAddr string) *Service {
	return &Service{
		Repository: repository,
		HttpAddr:   httpAddr,
	}
}
