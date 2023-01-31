package db

import (
	"strings"

	"github.com/carlmjohnson/requests"
)

type SelectBuilder struct{ *filterBuilder }

// Create a new builder for selecting from a table
func From(req *requests.Builder, table string) *SelectBuilder {
	return &SelectBuilder{newFilterBuilder(req.Pathf("/rest/v1/%s", table).Method("GET"))}
}

// Select columns to return
func (b *SelectBuilder) Select(columns ...string) *SelectBuilder {
	b.req.Param("select", strings.Join(columns, ","))
	return b
}

// Set the object to decode the response into
func (b *SelectBuilder) To(v any) *SelectBuilder {
	b.req.ToJSON(v)
	return b
}

type InsertBuilder struct{ *builder }

// Create a new builder for inserting into a table
func Insert(req *requests.Builder, table string) *InsertBuilder {
	return &InsertBuilder{newBuilder(req.Pathf("/rest/v1/%s", table).Method("POST").Header("Prefer", "return=minimal"))}
}

// Create a new builder for upserting into a table
func Upsert(req *requests.Builder, table string) *InsertBuilder {
	return &InsertBuilder{newBuilder(req.Pathf("/rest/v1/%s", table).Method("POST").Header("Prefer", "return=minimal, resolution=merge-duplicates"))}
}

// Add values to the request body
func (b *InsertBuilder) Values(v any) *InsertBuilder {
	b.req.BodyJSON(v)
	return b
}

// Set the object to decode the response into
func (b *InsertBuilder) To(v any) *InsertBuilder {
	b.req.ToJSON(v)
	return b
}

type UpdateBuilder struct{ *filterBuilder }

// Create a new builder for updating a table
func Update(req *requests.Builder, table string) *UpdateBuilder {
	return &UpdateBuilder{newFilterBuilder(req.Pathf("/rest/v1/%s", table).Method("PATCH"))}
}

// Add values to the request body
func (b *UpdateBuilder) Values(v any) *UpdateBuilder {
	b.req.BodyJSON(v)
	return b
}

// Set the object to decode the response into
func (b *UpdateBuilder) To(v any) *UpdateBuilder {
	b.req.ToJSON(v)
	return b
}

type DeleteBuilder struct{ *filterBuilder }

// Create a new builder for deleting from a table
func Delete(req *requests.Builder, table string) *DeleteBuilder {
	return &DeleteBuilder{newFilterBuilder(req.Pathf("/rest/v1/%s", table).Method("DELETE"))}
}
