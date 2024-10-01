package models

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"
)

type Token struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
	TokenHash []byte    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Expiry    time.Time `json:"expiry"`
}

// GetByToken returns a pointer to Token or error
func (t *Token) GetByToken(plainTextToken string) (*Token, error) {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	query := `SELECT id, user_id, email, token, token_hash, created_at, updated_at, expiry FROM tokens WHERE token = $1`

	var token Token

	row := db.QueryRowContext(ctx, query, plainTextToken)
	err := row.Scan(&token.ID, &token.UserID, &token.Email, &token.Token, &token.TokenHash, &token.CreatedAt, &token.UpdatedAt, &token.Expiry)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

// GetUserByToken returns a pointer to User or error
func (t *Token) GetUserByToken(token Token) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	query := `SELECT id, email, first_name, last_name, password, created_at, updated_at, user_active FROM users WHERE id = $1`

	var user User

	row := db.QueryRowContext(ctx, query, token.UserID)

	err := row.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.Active)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GenerateToken returns a pointer to Token or error
func (t *Token) GenerateToken(userID int, ttl time.Duration) (*Token, error) {
	token := &Token{
		UserID: userID,
		Expiry: time.Now().Add(ttl),
	}

	// Generate a random bytes
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}

	token.Token = base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)
	hashedToken := sha256.Sum256([]byte(token.Token))
	token.TokenHash = hashedToken[:]

	return token, nil
}

// AuthenticateToken returns a pointer to User or error
func (t *Token) AuthenticateToken(r *http.Request) (*User, error) {
	authorizationHeader := r.Header.Get("Authorization")
	if authorizationHeader == "" {
		return nil, errors.New("authorization header not found")
	}

	log.Println("authorizationHeader =", authorizationHeader)

	headerParts := strings.Split(authorizationHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return nil, errors.New("bearer not found in authorization header")
	}

	tokenStr := headerParts[1]
	log.Println("tokenStr =", tokenStr)

	if len(tokenStr) != 26 {
		return nil, errors.New("invalid token length")
	}

	token, err := t.GetByToken(tokenStr)
	if err != nil {
		return nil, errors.New("no matching token found")
	}

	if token.Expiry.Before(time.Now()) {
		return nil, errors.New("token is expired")
	}

	user, err := t.GetUserByToken(*token)
	if err != nil {
		return nil, errors.New("no matching user found")
	}

	if user.Active == 0 {
		return nil, errors.New("user not active")
	}

	return user, nil
}

// Insert new token for the given user
func (t *Token) Insert(token Token, u User) error {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	// Delete existing token
	stmt := `DELETE FROM tokens WHERE user_id = $1`
	_, err := db.ExecContext(ctx, stmt, token.UserID)
	if err != nil {
		return err
	}

	token.Email = u.Email

	stmt = `INSERT INTO tokens (user_id, email, token, token_hash, created_at, updated_at, expiry) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err = db.ExecContext(ctx, stmt, token.UserID, token.Email, token.Token, token.TokenHash, time.Now(), time.Now(), token.Expiry)
	if err != nil {
		return err
	}

	return nil
}

// DeleteByToken deletes a token
func (t *Token) DeleteByToken(plainTextToken string) error {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	stmt := `DELETE FROM tokens WHERE token = $1`
	_, err := db.ExecContext(ctx, stmt, plainTextToken)
	if err != nil {
		return err
	}

	return nil
}

// ValidToken checks the validity of the given token
func (t *Token) ValidToken(plainTextToken string) (bool, error) {
	token, err := t.GetByToken(plainTextToken)
	log.Println("token =", token)
	if err != nil {
		return false, errors.New("no matching token found")
	}

	_, err = t.GetUserByToken(*token)
	if err != nil {
		return false, errors.New("no matching user found")
	}

	if token.Expiry.Before(time.Now()) {
		return false, errors.New("token is expired")
	}

	return true, nil
}

// DeleteTokensForUser delete a token for the given user id
func (t *Token) DeleteTokensForUser(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	stmt := "delete from tokens where user_id = $1"
	_, err := db.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	return nil
}
