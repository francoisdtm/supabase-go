package db

import (
	"context"
	"fmt"
	"strings"

	"github.com/carlmjohnson/requests"
)

type SelectBuilder struct{ req *requests.Builder }

func From(req *requests.Builder, table string) *SelectBuilder {
	return &SelectBuilder{req.Pathf("/rest/v1/%s", table)}
}

func (b *SelectBuilder) Select(columns ...string) *SelectBuilder {
	b.req.Param("select", strings.Join(columns, ","))
	return b
}

func (b *SelectBuilder) Eq(column string, value string) *SelectBuilder {
	b.req.Param(column, fmt.Sprintf("eq.%s", value))
	return b
}

func (b *SelectBuilder) Neq(column string, value string) *SelectBuilder {
	b.req.Param(column, fmt.Sprintf("neq.%s", value))
	return b
}

func (b *SelectBuilder) Gt(column string, value string) *SelectBuilder {
	b.req.Param(column, fmt.Sprintf("gt.%s", value))
	return b
}

func (b *SelectBuilder) Gte(column string, value string) *SelectBuilder {
	b.req.Param(column, fmt.Sprintf("gte.%s", value))
	return b
}

func (b *SelectBuilder) Lt(column string, value string) *SelectBuilder {
	b.req.Param(column, fmt.Sprintf("lt.%s", value))
	return b
}

func (b *SelectBuilder) Lte(column string, value string) *SelectBuilder {
	b.req.Param(column, fmt.Sprintf("lte.%s", value))
	return b
}

func (b *SelectBuilder) Fetch(ctx context.Context, v any) error {
	return b.req.ToJSON(v).Fetch(ctx)
}
