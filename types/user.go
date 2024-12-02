package types

type UserRequest struct {
	Name   string `json:"name"`
	Pan    string `json:"pan" validate:"required"`
	Mobile string `json:"mobile" validate:"required,len=10,numeric"`
	Email  string `json:"email" validate:"required,email"`
}