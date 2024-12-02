package types

type UserRequest struct {
	Name   string `json:"name"`
	Pan    string `json:"pan"`
	Mobile string `json:"mobile"`
	Email  string `json:"email"`
}