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
