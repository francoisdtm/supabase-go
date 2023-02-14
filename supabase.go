package supabase

import (
	"context"
	"fmt"

	"github.com/francoisdtm/supabase-go/auth"
	"github.com/francoisdtm/supabase-go/db"

	"github.com/carlmjohnson/requests"
)

type Client struct{ url, key, accessToken, refreshToken string }

func NewClient(url string, key string) *Client {
	return &Client{url: url, key: key, accessToken: key}
}

// Get a base request builder
func (c *Client) api() *requests.Builder {
	req := requests.URL(c.url).Header("apikey", c.key)
	if c.accessToken != "" {
		req.Header("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
	}
	return req
}

// AUTH: Sign up a new user with the given credentials
func (c *Client) SignUp(ctx context.Context, email string, password string) (user *auth.User, err error) {
	return auth.SignUp(ctx, c.api(), email, password)
}

// AUTH: Sign in an existing user with the given credentials
func (c *Client) SignIn(ctx context.Context, email string, password string) (err error) {
	authDetails, err := auth.SignIn(ctx, c.api(), email, password)
	if err != nil {
		c.accessToken = authDetails.AccessToken
		c.refreshToken = authDetails.RefreshToken
	}
	return err
}

// AUTH: Sign in with a token
func (c *Client) SignInWithToken(ctx context.Context, token string) (user *auth.User, err error) {
	c.accessToken = token
	return c.GetUser(ctx)
}

// AUTH: Send a magic link to the given email address
func (c *Client) SignInWithMagicLink(ctx context.Context, email string) (err error) {
	return auth.SignInWithMagicLink(ctx, c.api(), email)
}

// AUTH: Get User details
func (c *Client) GetUser(ctx context.Context) (user *auth.User, err error) {
	return auth.GetUser(ctx, c.api())
}

// AUTH: Refresh the access token
func (c *Client) RefreshToken(ctx context.Context) (err error) {
	authDetails, err := auth.RefreshToken(ctx, c.api(), c.refreshToken)
	if err != nil {
		c.accessToken = authDetails.AccessToken
		c.refreshToken = authDetails.RefreshToken
	}
	return err
}

// AUTH: Sign out the given user
func (c *Client) SignOut(ctx context.Context) (err error) {
	err = auth.SignOut(ctx, c.api())
	if err == nil {
		c.accessToken = ""
		c.refreshToken = ""
	}
	return
}

// DB: Creates a new select query
func (c *Client) From(table string) *db.SelectBuilder {
	return db.From(c.api(), table)
}

// DB: Creates a new insert query
func (c *Client) Insert(table string) *db.InsertBuilder {
	return db.Insert(c.api(), table)
}

// DB: Creates a new insert query with the upsert option
func (c *Client) Upsert(table string) *db.InsertBuilder {
	return db.Upsert(c.api(), table)
}

// DB: Creates a new update query
func (c *Client) Update(table string) *db.UpdateBuilder {
	return db.Update(c.api(), table)
}

// DB: Creates a new delete query
func (c *Client) Delete(table string) *db.DeleteBuilder {
	return db.Delete(c.api(), table)
}
