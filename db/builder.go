package db

import (
	"context"

	"github.com/carlmjohnson/requests"
)

type builder struct{ req *requests.Builder }

func newBuilder(req *requests.Builder) *builder {
	return &builder{req}
}

// Execute the request
func (b *builder) Execute(ctx context.Context) error {
	return b.req.Fetch(ctx)
}
