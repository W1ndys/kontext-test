package repository

import (
	"database/sql"
	"time"
)

type User struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	Nickname  string    `json:"nickname"`
	Avatar    string    `json:"avatar"`
}

func (u *User) Scan(row *sql.Row) error {
	return row.Scan(&u.ID, &u.CreatedAt, &u.UpdatedAt, &u.Username, &u.Password, &u.Nickname, &u.Avatar)
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{db: GetDB()}
}

func (r *UserRepository) Create(user *User) error {
	result, err := r.db.Exec(
		"INSERT INTO users (username, password, nickname, avatar) VALUES (?, ?, ?, ?)",
		user.Username, user.Password, user.Nickname, user.Avatar,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = uint(id)
	return nil
}

func (r *UserRepository) GetByID(id uint) (*User, error) {
	user := &User{}
	err := r.db.QueryRow(
		"SELECT id, created_at, updated_at, username, password, nickname, avatar FROM users WHERE id = ?",
		id,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt, &user.Username, &user.Password, &user.Nickname, &user.Avatar)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetByUsername(username string) (*User, error) {
	user := &User{}
	err := r.db.QueryRow(
		"SELECT id, created_at, updated_at, username, password, nickname, avatar FROM users WHERE username = ?",
		username,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt, &user.Username, &user.Password, &user.Nickname, &user.Avatar)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) Update(user *User) error {
	_, err := r.db.Exec(
		"UPDATE users SET username = ?, password = ?, nickname = ?, avatar = ? WHERE id = ?",
		user.Username, user.Password, user.Nickname, user.Avatar, user.ID,
	)
	return err
}
