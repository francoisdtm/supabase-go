package supabase

import (
	"context"
	"fmt"
	"time"
)

type User struct {
	ID                 string                    `json:"id"`
	Aud                string                    `json:"aud"`
	Role               string                    `json:"role"`
	Email              string                    `json:"email"`
	Phone              string                    `json:"phone"`
	InvitedAt          time.Time                 `json:"invited_at"`
	ConfirmedAt        time.Time                 `json:"confirmed_at"`
	ConfirmationSentAt time.Time                 `json:"confirmation_sent_at"`
	AppMetadata        struct{ provider string } `json:"app_metadata"`
	UserMetadata       map[string]interface{}    `json:"user_metadata"`
	CreatedAt          time.Time                 `json:"created_at"`
	UpdatedAt          time.Time                 `json:"updated_at"`
}

type AuthDetails struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	User         User   `json:"user"`
}

// Error returned by the auth API
type authError struct {
	ErrorName        string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

// Create a new error from the error name and description
func (e *authError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorName, e.ErrorDescription)
}

// Manually set the access key to use for authentication
func (c *Client) SetAccessToken(accessToken string) {
	c.accessToken = accessToken
}

// Check if the user is signed in
func (c *Client) IsSignedIn() bool {
	return c.accessToken != ""
}

// Sign up a new user with the given credentials
func (c *Client) SignUp(ctx context.Context, email string, password string) (user *User, err error) {
	var errJSON *authError
	data := map[string]string{"email": email, "password": password}
	err = c.api().Path("/auth/v1/signup").BodyJSON(&data).ToJSON(&user).ErrorJSON(&errJSON).Fetch(ctx)
	if errJSON != nil {
		err = errJSON
	}
	return
}

// Sign in an existing user with the given credentials
func (c *Client) SignIn(ctx context.Context, email string, password string) (user *User, err error) {
	var errJSON *authError
	var details *AuthDetails
	data := map[string]string{"email": email, "password": password}
	err = c.api().Path("/auth/v1/token").Param("grant_type", "password").BodyJSON(&data).ToJSON(&details).ErrorJSON(&errJSON).Fetch(ctx)
	if errJSON != nil {
		err = errJSON
	}
	if err == nil {
		c.accessToken = details.AccessToken
		c.refreshToken = details.RefreshToken
		user = &details.User
	}
	return
}

// Send a magic link to the given email address
func (c *Client) SignInWithMagicLink(ctx context.Context, email string) (err error) {
	var errJSON *authError
	data := map[string]string{"email": email}
	err = c.api().Path("/auth/v1/magiclink").BodyJSON(&data).ErrorJSON(&errJSON).Fetch(ctx)
	if errJSON != nil {
		err = errJSON
	}
	return
}

// Get User details from the given access token
func (c *Client) GetUser(ctx context.Context) (user *User, err error) {
	if !c.IsSignedIn() {
		return nil, fmt.Errorf("you must sign in before getting user details")
	}

	var errJSON *authError
	err = c.api().Path("/auth/v1/user").ToJSON(&user).ErrorJSON(&errJSON).Fetch(ctx)
	if errJSON != nil {
		err = errJSON
	}
	return
}

// Refresh the access token of the given user
func (c *Client) RefreshToken(ctx context.Context) (err error) {
	if !c.IsSignedIn() {
		return fmt.Errorf("you must sign in before refreshing the token")
	}

	var errJSON *authError
	var details *AuthDetails
	data := map[string]string{"refresh_token": c.refreshToken}
	err = c.api().Path("/auth/v1/token").Param("grant_type", "refresh_token").BodyJSON(&data).ToJSON(&details).ErrorJSON(&errJSON).Fetch(ctx)
	if errJSON != nil {
		err = errJSON
	}
	if err == nil {
		c.accessToken = details.AccessToken
		c.refreshToken = details.RefreshToken
	}
	return
}

// Refresh the access token of the given user at the given interval
func (c *Client) RefreshTokenAtInterval(ctx context.Context, interval time.Duration) (err error) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(interval):
			if err = c.RefreshToken(ctx); err != nil {
				return
			}
		}
	}
}

// Sign out the given user
func (c *Client) SignOut(ctx context.Context) (err error) {
	if !c.IsSignedIn() {
		return fmt.Errorf("you must sign in before signing out")
	}

	var errJSON *authError
	err = c.api().Path("/auth/v1/signout").BodyJSON(struct{}{}).ErrorJSON(&errJSON).Fetch(ctx)
	if errJSON != nil {
		err = errJSON
	}
	return
}
