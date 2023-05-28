package domain

type UserDecodeJWT struct {
	ID     string
	Name   string
	Email  string
	Method string
	Role   string
}