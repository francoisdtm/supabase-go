package auth

import (
	"context"

	"github.com/carlmjohnson/requests"
)

// Sign up a new user with the given credentials
func SignUp(ctx context.Context, api *requests.Builder, email string, password string) (user *User, err error) {
	data := map[string]string{"email": email, "password": password}
	err = Run(ctx, api.Path("/auth/v1/signup").BodyJSON(&data).ToJSON(&user))
	return
}

// Sign in an existing user with the given credentials
func SignIn(ctx context.Context, api *requests.Builder, email string, password string) (details *AuthDetails, err error) {
	data := map[string]string{"email": email, "password": password}
	err = Run(ctx, api.Path("/auth/v1/token").Param("grant_type", "password").BodyJSON(&data).ToJSON(&details))
	return
}

// Send a magic link to the given email address
func SignInWithMagicLink(ctx context.Context, api *requests.Builder, email string) (err error) {
	data := map[string]string{"email": email}
	err = Run(ctx, api.Path("/auth/v1/magiclink").BodyJSON(&data))
	return
}

// Get User details from the given access token
func GetUser(ctx context.Context, api *requests.Builder) (user *User, err error) {
	err = Run(ctx, api.Path("/auth/v1/user").ToJSON(&user))
	return
}

// Refresh the access token of the given user
func RefreshToken(ctx context.Context, api *requests.Builder, refreshToken string) (details *AuthDetails, err error) {
	data := map[string]string{"refresh_token": refreshToken}
	err = Run(ctx, api.Path("/auth/v1/token").Param("grant_type", "refresh_token").BodyJSON(&data).ToJSON(&details))
	return
}

// Sign out the given user
func SignOut(ctx context.Context, api *requests.Builder) (err error) {
	err = Run(ctx, api.Path("/auth/v1/signout").Method("POST"))
	return
}
