package models

import (
	"database/sql"
	"fmt"
	"time"
)

const (
	DefaultResetDuration = 1 * time.Hour
)

type PasswordReset struct {
	ID    int
	UseID int
	// Token is only set when a passwordReset is being created.
	Token     string
	TokenHash string
	ExpiresAt time.Time
	sql.NullTime
}

type passwordResetService struct {
	DB *sql.DB
	// Token is only set when creating a new session.  When look up a session
	// this will be left empty, as we only store the hash of a session token
	// in our database ans we cannot reverse it into a raw token.
	BytesPerToken int
	// Duration is the amount of time that a PasswordReset is valid for.
	//Defautl to DefaultResetDuration
	Duration time.Duration
}

func (service *passwordResetService) Create(email string) (*PasswordReset, error) {
	return nil, fmt.Errorf("TODO: implement PasswordResetService.Create")
}
 func (service *passwordResetService) Consume(token string) (*User, error) {
	return nil, fmt.Errorf("TODO: Implement PAswordResetService.Consume")
 }