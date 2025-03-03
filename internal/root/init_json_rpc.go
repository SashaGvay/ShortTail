package root

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	urlModels "short_tail/internal/domain/url/models"

	json "github.com/goccy/go-json"
	jsonrpc "github.com/osamingo/jsonrpc/v2"
)

func (r *Root) initJsonRpc(_ context.Context) error {
	r.Adapters.JsonRpc = jsonrpc.NewMethodRepository()

	shortHandler := jsonrpc.HandlerFunc(
		func(ctx context.Context, params *json.RawMessage) (resp any, jsonRpcErr *jsonrpc.Error) {
			if params == nil {
				return nil, jsonrpc.ErrInvalidParams()
			}

			dto := &urlModels.URL{}

			err := json.Unmarshal(*params, dto)
			if err != nil {
				return nil, jsonrpc.ErrParse()
			}

			dto, err = r.Entity.Url.Service.Short(ctx, dto.Original)
			if err != nil {
				log.Printf("r.Entity.Url.Service.Short: %v", err)
				return nil, jsonrpc.ErrInternal()
			}

			return dto, nil
		},
	)

	unShortHandler := jsonrpc.HandlerFunc(
		func(ctx context.Context, params *json.RawMessage) (resp any, jsonRpcErr *jsonrpc.Error) {
			if params == nil {
				return nil, jsonrpc.ErrInvalidParams()
			}

			dto := &urlModels.URL{}

			err := json.Unmarshal(*params, dto)
			if err != nil {
				return nil, jsonrpc.ErrParse()
			}

			dto, err = r.Entity.Url.Service.UnShort(ctx, dto.Alias)
			if err != nil {
				if errors.Is(err, urlModels.ErrNotFound) {
					return nil, jsonrpc.ErrInvalidParams()
				}

				log.Printf("r.Entity.Url.Service.UnShort: %v", err)
				return nil, jsonrpc.ErrInternal()
			}

			return dto, nil
		},
	)

	err := r.Adapters.JsonRpc.RegisterMethod("Short", shortHandler, urlModels.URL{}, urlModels.URL{})
	if err != nil {
		return fmt.Errorf("r.Adapters.JsonRpc.RegisterMethod: %w", err)
	}

	err = r.Adapters.JsonRpc.RegisterMethod("UnShort", unShortHandler, urlModels.URL{}, urlModels.URL{})
	if err != nil {
		return fmt.Errorf("r.Adapters.JsonRpc.RegisterMethod: %w", err)
	}

	http.Handle("/jrpc", r.Adapters.JsonRpc)

	if r.Cfg.ENV == "DEV" {
		http.HandleFunc("/jrpc/debug", r.Adapters.JsonRpc.ServeDebug)
	}

	return nil
}
