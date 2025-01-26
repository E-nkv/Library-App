package database

import (
	"database/sql"
	"fmt"
)

type IUserModel interface {
	GetUser()
}

type Models struct {
	DB   *sql.DB
	User IUserModel
}

func (m *Models) MultipleResourceOperation() {
	fmt.Println("Multiple resource op")
}
func NewModels(DB *sql.DB) *Models {
	return &Models{DB: DB, User: NewUserModel(DB)}
}
