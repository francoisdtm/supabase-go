package db

import (
	"context"

	"github.com/carlmjohnson/requests"
)

type InsertBuilder struct{ req *requests.Builder }

type InsertionResult struct {
	Data       []any  `json:"data"`
	Status     int    `json:"status"`
	StatusText string `json:"statusText"`
}

func Insert(req *requests.Builder, table string) *InsertBuilder {
	return &InsertBuilder{req.Pathf("/rest/v1/%s", table).Method("POST").Header("Prefer", "return=minimal")}
}

func Upsert(req *requests.Builder, table string) *InsertBuilder {
	return &InsertBuilder{req.Pathf("/rest/v1/%s", table).Method("POST").Header("Prefer", "return=minimal, resolution=merge-duplicates")}
}

func (b *InsertBuilder) Values(v any) *InsertBuilder {
	b.req.BodyJSON(v)
	return b
}

func (b *InsertBuilder) To(v any) *InsertBuilder {
	b.req.Header("Prefer", "return=representation").ToJSON(v)
	return b
}

func (b *InsertBuilder) Execute(ctx context.Context) error {
	return b.req.Fetch(ctx)
}
