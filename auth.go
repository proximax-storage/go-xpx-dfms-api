package api

import (
	"context"
	"time"
)

type AccessToken []byte

type Auth interface {

	// GenerateToken generates a new token for some time with id and allowed routes
	GenerateToken(context.Context, string, time.Duration, ...string) (AccessToken, error)

	// VerifyAccess verifies token access
	VerifyAccessTo(context.Context, string, AccessToken) error

	// List returns list of ids of all active tokens from store
	List(context.Context) ([]string, error)

	// RevokeToken revokes a token by ID and delete from store
	RevokeToken(context.Context, string) error
}
