package models

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrDuplicateEmail = errors.New("duplicate email address")
	ErrNotFound       = errors.New("not found")
)

type User struct {
	ID       int64  `db:"id" json:"id"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"-"`
	Token    string `db:"token" json:"token"`
}

func (u *User) String() string {
	return fmt.Sprintf("&User{id: %d, email: %s ", u.ID, u.Email)
}

// hash password
func (u *User) GenerateHash(password string) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	u.Password = string(hash)
}

// Check password
func (u *User) Authenticate(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return false
	}
	return true
}

// created user
func CreateUser(user *User) error {
	result, err := db.Exec("INSERT INTO admin_users (email, password, token) VALUES (?,?,?)", user.Email, user.Password, user.Token)
	if err != nil {
		if strings.Contains(err.Error(), "1062") {
			return ErrDuplicateEmail
		}
		return err
	}
	user.ID, _ = result.LastInsertId()
	return nil
}

// found user
func GetUserByEmail(email string) (*User, error) {
	user := new(User)
	err := db.Get(user, "select id, email, password, token from admin_users where email = ?", email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return user, nil
}

func GetUserByToken(token string) (*User, error) {
	user := new(User)
	err := db.Get(user, "select id, email, password, token from admin_users where token = ?", token)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return user, nil
}
