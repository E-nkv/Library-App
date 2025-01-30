package models

import (
	"database/sql"
	"library/db/types"

	_ "github.com/lib/pq"
)

type Models struct {
	Users UserModel
	Books BookModel
}
type UserModel interface {
	GetUsersWithTags(limit int, lastID int64) ([]*types.User, error)
	GetUserWithTags(id int64) (*types.User, error)
	GetUsers(limit int, lastID int64) ([]*types.User, error)
	GetUser(id int64) (*types.User, error)
	CreateUser(*types.UserCreate) (int64, error)
	DeleteUser(id int64) error
}
type BookModel interface {
	GetBooks()
	GetBook()
}

type VoidUserModel struct {
	DB *sql.DB
	UserModel
}
type VoidBookModel struct {
	DB *sql.DB
	BookModel
}

func InitDB() (*sql.DB, error) {
	connStr := "postgres://postgres:admin@localhost:5432/libraryapp?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
func NewModels(db *sql.DB) *Models {
	return &Models{
		Users: PsqlUserModel{DB: db},
		Books: VoidBookModel{DB: db},
	}
}
