package root

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"short_tail/config"
	"short_tail/internal/domain/url"
	urlRepository "short_tail/internal/domain/url/repository"
	urlService "short_tail/internal/domain/url/service"

	badger "github.com/dgraph-io/badger/v4"
	jsonrpc "github.com/osamingo/jsonrpc/v2"
)

type Root struct {
	Cfg *config.Conf

	Adapters struct {
		JsonRpc *jsonrpc.MethodRepository
	}

	Infrastructure struct {
		DbInstance *badger.DB
	}

	Entity struct {
		Url struct {
			Service    url.Service
			Repository url.Repository
		}
	}
}

func New(ctx context.Context, cfg *config.Conf) (*Root, error) {
	r := &Root{
		Cfg: cfg,
	}

	err := r.initInfrastructure(ctx)
	if err != nil {
		return nil, fmt.Errorf("r.initInfrastructure: %w", err)
	}

	err = r.initEntities(ctx)
	if err != nil {
		return nil, fmt.Errorf("r.initEntities: %w", err)
	}

	err = r.initJsonRpc(ctx)
	if err != nil {
		return nil, fmt.Errorf("r.initJsonRpc: %w", err)
	}

	r.initHttp(ctx)

	return r, nil
}

func (r *Root) initInfrastructure(_ context.Context) error {
	var err error

	r.Infrastructure.DbInstance, err = badger.Open(badger.DefaultOptions(r.Cfg.BadgetDataDir))
	if err != nil {
		return fmt.Errorf("badger.Open: %w", err)
	}

	return nil
}

func (r *Root) initEntities(_ context.Context) error {
	r.Entity.Url.Repository = urlRepository.New(r.Infrastructure.DbInstance)
	r.Entity.Url.Service = urlService.New(r.Entity.Url.Repository)

	return nil
}

func (r *Root) Run(_ context.Context) error {
	log.Printf("Listening on %s", r.Cfg.HttpAddr)

	if err := http.ListenAndServe(r.Cfg.HttpAddr, http.DefaultServeMux); err != nil {
		return fmt.Errorf("http.ListenAndServe: %w", err)
	}

	return nil
}
