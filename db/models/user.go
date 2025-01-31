package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"library/db/types"
	"library/errs"

	pq "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

/* type UserModel interface {
	GetUsers(limit int, lastID int64) ([]*types.User, error)
	GetUser(id int64) (*types.User, error)
	CreateUser(*types.UserCreate)(int64, error)
} */

type PsqlUserModel struct {
	DB *sql.DB
}

const defaultLimit = 10

func (m PsqlUserModel) GetUsersWithTags(limit int, lastID int64) ([]*types.User, error) {
	if limit <= 0 {
		limit = defaultLimit
	}
	q := createQueryGetUsersWithTags(limit, lastID)
	rows, err := m.DB.Query(q)
	if err != nil {
		return nil, err
	}
	var users []*types.User
	for i := 0; rows.Next(); i++ {
		var u types.User

		args := []any{&u.ID, &u.FullName, &u.Email, &u.TagsJson}
		if err = rows.Scan(args...); err != nil {
			fmt.Println("Error exttracting row into user.. ", err)
			continue
		}

		users = append(users, &u)
	}

	return users, nil
}
func (m PsqlUserModel) GetUsers(limit int, lastID int64) ([]*types.User, error) {
	if limit <= 0 {
		limit = defaultLimit
	}
	q := createQueryGetUsers(limit, lastID)
	rows, err := m.DB.Query(q)
	if err != nil {
		return nil, err
	}
	var users []*types.User
	for i := 0; rows.Next(); i++ {
		var u types.User

		args := []any{&u.ID, &u.FullName, &u.Email, &u.IsActive, &u.IsVerified, &u.Role}
		if err = rows.Scan(args...); err != nil {

			return nil, err
		}

		users = append(users, &u)
	}

	return users, nil
}

func (m PsqlUserModel) GetUserWithTags(id int64) (*types.User, error) {
	var u types.User
	q := createQueryGetUserWithTags(id)
	r := m.DB.QueryRow(q)
	args := []any{&u.ID, &u.FullName, &u.Email, &u.HashPass, &u.TagsJson}
	if err := r.Scan(args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.ErrNotFound
		}
		return nil, err
	}
	return &u, nil
}
func (m PsqlUserModel) GetUser(id int64) (*types.User, error) {
	var u types.User
	q := createQueryGetUser(id)
	r := m.DB.QueryRow(q)
	args := []any{&u.ID, &u.FullName, &u.Email, &u.HashPass, &u.IsActive, &u.IsVerified, &u.Role}
	if err := r.Scan(args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.ErrNotFound
		}
		return nil, err
	}
	return &u, nil
}
func (m PsqlUserModel) GetUserByEmail(email string) (*types.User, error) {
	var u types.User
	q := createQueryGetUserByEmail(email)
	r := m.DB.QueryRow(q)
	args := []any{&u.ID, &u.FullName, &u.Email, &u.HashPass, &u.IsActive, &u.IsVerified, &u.Role}
	if err := r.Scan(args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.ErrNotFound
		}
		return nil, err
	}
	return &u, nil
}

func (m PsqlUserModel) CreateUser(u *types.UserCreate) (int64, error) {
	qUser := "INSERT INTO users (full_name, email, hash_pass, role) VALUES($1, $2, $3, $4) RETURNING id"
	tx, err := m.DB.BeginTx(context.Background(), nil)
	if err != nil {
		return -1, err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(u.PasswdPlain), bcrypt.DefaultCost)
	if err != nil {
		return -1, err
	}
	r := tx.QueryRow(qUser, u.FullName, u.Email, hash, u.Role)
	var userID int64
	if err := r.Scan(&userID); err != nil {
		tx.Rollback()
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				err = errs.ErrDuplicateEmail
			}
		}
		return -1, err
	}
	for _, tag := range u.Tags {
		if _, err := tx.Exec("INSERT INTO users_tags VALUES($1, $2)", userID, tag.ID); err != nil {
			tx.Rollback()
			return -1, err
		}
	}
	tx.Commit()
	return userID, nil
}

func (m PsqlUserModel) DeleteUser(id int64) error {

	var isActive bool
	r := m.DB.QueryRow("SELECT is_active FROM users WHERE id = $1", id)
	if err := r.Scan(&isActive); err != nil {
		return err
	}
	if !isActive {
		return fmt.Errorf("user is already inactive")
	}
	if _, err := m.DB.Exec("UPDATE users SET is_active = FALSE WHERE id = $1", id); err != nil {
		return err
	}
	return nil
}
