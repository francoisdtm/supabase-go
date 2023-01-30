package supabase

import (
	"context"
	"fmt"
	"strings"

	"github.com/carlmjohnson/requests"
)

type Builder struct{ req *requests.Builder }

func (c *Client) From(table string) *Builder {
	return &Builder{c.api().Pathf("/rest/v1/%s", table)}
}

func (b *Builder) Select(columns ...string) *Builder {
	b.req.Param("select", strings.Join(columns, ","))
	return b
}

func (b *Builder) Delete() *Builder {
	b.req.Method("DELETE")
	return b
}

func (b *Builder) Eq(column string, value string) *Builder {
	b.req.Param(column, fmt.Sprintf("eq.%s", value))
	return b
}

func (b *Builder) Neq(column string, value string) *Builder {
	b.req.Param(column, fmt.Sprintf("neq.%s", value))
	return b
}

func (b *Builder) Gt(column string, value string) *Builder {
	b.req.Param(column, fmt.Sprintf("gt.%s", value))
	return b
}

func (b *Builder) Gte(column string, value string) *Builder {
	b.req.Param(column, fmt.Sprintf("gte.%s", value))
	return b
}

func (b *Builder) Lt(column string, value string) *Builder {
	b.req.Param(column, fmt.Sprintf("lt.%s", value))
	return b
}

func (b *Builder) Lte(column string, value string) *Builder {
	b.req.Param(column, fmt.Sprintf("lte.%s", value))
	return b
}

func (b *Builder) ToJSON(v any) *Builder {
	b.req.ToJSON(v)
	return b
}

func (b *Builder) Execute(ctx context.Context) error {
	return b.req.Fetch(ctx)
}
