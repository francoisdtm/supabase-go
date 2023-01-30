package db

import (
	"context"
	"fmt"

	"github.com/carlmjohnson/requests"
)

type DeleteBuilder struct{ req *requests.Builder }

func Delete(req *requests.Builder, table string) *DeleteBuilder {
	return &DeleteBuilder{req.Pathf("/rest/v1/%s", table).Method("DELETE")}
}

func (b *DeleteBuilder) Eq(column string, value string) *DeleteBuilder {
	b.req.Param(column, fmt.Sprintf("eq.%s", value))
	return b
}

func (b *DeleteBuilder) Neq(column string, value string) *DeleteBuilder {
	b.req.Param(column, fmt.Sprintf("neq.%s", value))
	return b
}

func (b *DeleteBuilder) Gt(column string, value string) *DeleteBuilder {
	b.req.Param(column, fmt.Sprintf("gt.%s", value))
	return b
}

func (b *DeleteBuilder) Gte(column string, value string) *DeleteBuilder {
	b.req.Param(column, fmt.Sprintf("gte.%s", value))
	return b
}

func (b *DeleteBuilder) Lt(column string, value string) *DeleteBuilder {
	b.req.Param(column, fmt.Sprintf("lt.%s", value))
	return b
}

func (b *DeleteBuilder) Lte(column string, value string) *DeleteBuilder {
	b.req.Param(column, fmt.Sprintf("lte.%s", value))
	return b
}

func (b *DeleteBuilder) Execute(ctx context.Context, v any) error {
	return b.req.ToJSON(v).Fetch(ctx)
}
