package entity

type User struct {
	ID       int
	Email    string
	Password string
}

func CreateUser(email, password string) *User {
	return &User{
		Email:    email,
		Password: password,
		ID:       -1,
	}
}
