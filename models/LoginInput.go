package models

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}