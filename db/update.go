package db

import (
	"context"
	"fmt"

	"github.com/carlmjohnson/requests"
)

type UpdateBuilder struct{ req *requests.Builder }

func Update(req *requests.Builder, table string) *UpdateBuilder {
	return &UpdateBuilder{req.Pathf("/rest/v1/%s", table).Method("PATCH")}
}

func (b *UpdateBuilder) Eq(column string, value string) *UpdateBuilder {
	b.req.Param(column, fmt.Sprintf("eq.%s", value))
	return b
}

func (b *UpdateBuilder) Neq(column string, value string) *UpdateBuilder {
	b.req.Param(column, fmt.Sprintf("neq.%s", value))
	return b
}

func (b *UpdateBuilder) Gt(column string, value string) *UpdateBuilder {
	b.req.Param(column, fmt.Sprintf("gt.%s", value))
	return b
}

func (b *UpdateBuilder) Gte(column string, value string) *UpdateBuilder {
	b.req.Param(column, fmt.Sprintf("gte.%s", value))
	return b
}

func (b *UpdateBuilder) Lt(column string, value string) *UpdateBuilder {
	b.req.Param(column, fmt.Sprintf("lt.%s", value))
	return b
}

func (b *UpdateBuilder) Lte(column string, value string) *UpdateBuilder {
	b.req.Param(column, fmt.Sprintf("lte.%s", value))
	return b
}

func (b *UpdateBuilder) Execute(ctx context.Context, v any) error {
	return b.req.ToJSON(v).Fetch(ctx)
}
