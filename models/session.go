package models

import "database/sql"

type Session struct {
	ID        int
	UserID    int
	// Token is only set when creating a new session.  When look up a session
	// this will be left empty, as we only store the hash of a session token
	// in our database ans we cannot reverse it into a raw token.
	Token string
	TokenHash string
}

type SessionService struct {
	DB *sql.DB
}

// 1. Query a Session via raw token, then query the user separately via
func (ss *SessionService) Create(userID int) (*Session ,error){
	// TODO : create the session token
	// TODO : Implement SessionService.Create
	return nil, nil 
}

// 2. Query a user via raw token using the SessionService
func (ss *SessionService) User(token string) (*User, error){
	// TODO : implement sessionService.User
	return nil, nil
}