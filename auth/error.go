package auth

import (
	"context"
	"fmt"

	"github.com/carlmjohnson/requests"
)

// Error returned by the auth API
type authError struct {
	ErrorName        string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

// Create a new error from the error name and description
func (e *authError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorName, e.ErrorDescription)
}

// Run the request and return an error if the auth API returned an error
func Run(ctx context.Context, req *requests.Builder) error {
	var authErr *authError
	if err := req.ErrorJSON(&authErr).Fetch(ctx); err != nil {
		return err
	}
	if authErr != nil {
		return authErr
	}
	return nil
}
