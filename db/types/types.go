package types

type User struct {
	ID         int64    `json:"id"`
	FullName   string   `json:"fullName"`
	Email      string   `json:"email"`
	HashPass   string   `json:"hashPass"`
	IsVerified bool     `json:"isVerified"`
	Tags       []string `json:"tags"`
}
