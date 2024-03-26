package domain

type UserID string

type User struct {
	ID     UserID `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Phone  int64  `json:"phone"`
	Avatar string `json:"avatar"`
}
