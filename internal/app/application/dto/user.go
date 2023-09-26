package dto

type SignUp struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"min=4,max=10" json:"password"`
}
type SignIn struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"min=4,max=10" json:"password"`
}
