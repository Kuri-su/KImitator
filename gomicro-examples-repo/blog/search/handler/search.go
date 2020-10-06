package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	search "github.com/micro/examples/blog/search/proto/search"
)

type Search struct{}

func (e *Search) Index(ctx context.Context, req *search.IndexRequest, rsp *search.IndexResponse) error {
	log.Info("Received Search.Index request")
	return nil
}

func (e *Search) Search(ctx context.Context, req *search.SearchRequest, rsp *search.SearchResponse) error {
	log.Info("Received Search.Search request")
	return nil
}
