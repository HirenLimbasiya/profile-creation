package types

type UserRequest struct {
	Name   string `json:"name"`
	Pan    string `json:"pan" validate:"required,pan"`
	Mobile int64  `json:"mobile" validate:"required"`
	Email  string `json:"email" validate:"required,email"`
}
