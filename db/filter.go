package db

import (
	"fmt"
	"strings"

	"github.com/carlmjohnson/requests"
)

type filterBuilder struct{ *builder }

func newFilterBuilder(req *requests.Builder) *filterBuilder {
	return &filterBuilder{newBuilder(req)}
}

func (b *filterBuilder) Eq(column string, value string) *filterBuilder {
	b.req.Param(column, fmt.Sprintf("eq.%s", value))
	return b
}

func (b *filterBuilder) Gt(column string, value string) *filterBuilder {
	b.req.Param(column, fmt.Sprintf("gt.%s", value))
	return b
}

func (b *filterBuilder) Gte(column string, value string) *filterBuilder {
	b.req.Param(column, fmt.Sprintf("gte.%s", value))
	return b
}

func (b *filterBuilder) Lt(column string, value string) *filterBuilder {
	b.req.Param(column, fmt.Sprintf("lt.%s", value))
	return b
}

func (b *filterBuilder) Lte(column string, value string) *filterBuilder {
	b.req.Param(column, fmt.Sprintf("lte.%s", value))
	return b
}

func (b *filterBuilder) Neq(column string, value string) *filterBuilder {
	b.req.Param(column, fmt.Sprintf("neq.%s", value))
	return b
}

func (b *filterBuilder) Like(column string, value string) *filterBuilder {
	b.req.Param(column, fmt.Sprintf("like.%s", value))
	return b
}

func (b *filterBuilder) Ilike(column string, value string) *filterBuilder {
	b.req.Param(column, fmt.Sprintf("ilike.%s", value))
	return b
}

func (b *filterBuilder) Match(column string, value string) *filterBuilder {
	b.req.Param(column, fmt.Sprintf("match.%s", value))
	return b
}

func (b *filterBuilder) Imatch(column string, value string) *filterBuilder {
	b.req.Param(column, fmt.Sprintf("imatch.%s", value))
	return b
}

func (b *filterBuilder) In(column string, values ...string) *filterBuilder {
	var v []string
	for _, value := range values {
		if strings.ContainsAny(value, "[,()] ") {
			v = append(v, fmt.Sprintf(`"%s"`, value))
		} else {
			v = append(v, value)
		}
	}
	b.req.Param(column, fmt.Sprintf("in.(%s)", strings.Join(v, ",")))
	return b
}

func (b *filterBuilder) Is(column string, value string) *filterBuilder {
	b.req.Param(column, fmt.Sprintf("is.%s", value))
	return b
}

func (b *filterBuilder) Contains(column string, value string) *filterBuilder {
	b.req.Param(column, fmt.Sprintf("cs.%s", value))
	return b
}

func (b *filterBuilder) Contained(column string, values ...string) *filterBuilder {
	b.req.Param(column, fmt.Sprintf("cd.%s", strings.Join(values, ",")))
	return b
}

func (b *filterBuilder) Overlap(column string, values ...string) *filterBuilder {
	b.req.Param(column, fmt.Sprintf("ov.%s", strings.Join(values, ",")))
	return b
}

func (b *filterBuilder) RangeLt(column string, values ...string) *filterBuilder {
	b.req.Param(column, fmt.Sprintf("sl.%s", strings.Join(values, ",")))
	return b
}

func (b *filterBuilder) RangeGt(column string, values ...string) *filterBuilder {
	b.req.Param(column, fmt.Sprintf("sr.%s", strings.Join(values, ",")))
	return b
}

func (b *filterBuilder) RangeLte(column string, values ...string) *filterBuilder {
	b.req.Param(column, fmt.Sprintf("nxr.%s", strings.Join(values, ",")))
	return b
}

func (b *filterBuilder) RangeGte(column string, values ...string) *filterBuilder {
	b.req.Param(column, fmt.Sprintf("nxl.%s", strings.Join(values, ",")))
	return b
}

func (b *filterBuilder) RangeAdjacent(column string, values ...string) *filterBuilder {
	b.req.Param(column, fmt.Sprintf("adj.%s", strings.Join(values, ",")))
	return b
}

func (b *filterBuilder) Or(filters ...string) *filterBuilder {
	b.req.Param("or", strings.Join(filters, ","))
	return b
}

func (b *filterBuilder) And(filters ...string) *filterBuilder {
	b.req.Param("and", strings.Join(filters, ","))
	return b
}

func (b *filterBuilder) Not(filter string) *filterBuilder {
	b.req.Param("not", filter)
	return b
}
