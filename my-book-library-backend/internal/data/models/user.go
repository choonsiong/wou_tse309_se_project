package models

import (
	"context"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Active    int       `json:"active"`
	Token     Token     `json:"token"`
}

// GetAll returns a slice of pointer to User or error
func (u *User) GetAll() ([]*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	//query := `SELECT id, email, first_name, last_name, password, created_at, updated_at, user_active case when (select count(id) from tokens t where user_id = users.id and t.expiry > now()) > 0 then 1 else 0 end as has_token FROM users ORDER BY last_name`
	query := `SELECT id, email, first_name, last_name, password, created_at, updated_at, user_active FROM users ORDER BY last_name`
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer CloseRows(rows)

	var users []*User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.Active)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

// GetByEmail returns a pointer to User or error
func (u *User) GetByEmail(email string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	query := `SELECT id, email, first_name, last_name, password, created_at, updated_at, user_active FROM users WHERE email = $1`

	var user User

	row := db.QueryRowContext(ctx, query, email)

	err := row.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.Active)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetById returns a pointer to User or error
func (u *User) GetById(id int) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	query := `SELECT id, email, first_name, last_name, password, created_at, updated_at, user_active FROM users WHERE id = $1`

	var user User

	row := db.QueryRowContext(ctx, query, id)

	err := row.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.Active)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Update updates a user
func (u *User) Update() error {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	stmt := `UPDATE users SET email = $1, first_name = $2, last_name = $3, updated_at = $4, user_active = $5 WHERE id = $6`

	_, err := db.ExecContext(ctx, stmt, u.Email, u.FirstName, u.LastName, time.Now(), u.Active, u.ID)
	if err != nil {
		return err
	}

	return nil
}

// Delete deletes a user
func (u *User) Delete() error {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	stmt := `DELETE FROM users WHERE id = $1`

	_, err := db.ExecContext(ctx, stmt, u.ID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteById deletes a user by id
func (u *User) DeleteById(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), databaseTimeout)
	defer cancel()

	stmt := `DELETE FROM users WHERE id = $1`

	_, err := db.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	return nil
}
