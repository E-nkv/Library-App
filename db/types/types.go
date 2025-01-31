package types

import "time"

type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserAuth struct {
	ID   int64  `json:"id"`
	Role string `json:"role"`
}

type UserCreate struct {
	FullName    string `json:"fullName"`
	PasswdPlain string `json:"password"`
	Tags        []Tag  `json:"tags"`
	Email       string `json:"email"`
	Role        string `json:"role"`
}

type UserLogin struct {
	Email       string `json:"email"`
	PlainPasswd string `json:"password"`
}
type User struct {
	ID         int64     `json:"id"`
	FullName   string    `json:"fullName"`
	Email      string    `json:"email"`
	HashPass   string    `json:"-"`
	IsVerified bool      `json:"isVerified"`
	IsActive   bool      `json:"is_active"`
	Role       string    `json:"role"`
	TagsJson   string    `json:"tags"`
	CreatedAt  time.Time `json:"created_at"`
}
