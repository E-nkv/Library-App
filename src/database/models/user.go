package database

import (
	"database/sql"
	"fmt"
)

type UserModel struct {
	DB *sql.DB
}

func (u *UserModel) GetUser() {
	fmt.Println("Getting user from model")
}

func NewUserModel(DB *sql.DB) *UserModel {
	return &UserModel{DB}
}
