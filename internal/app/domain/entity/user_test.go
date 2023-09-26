package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	// Stub
	email := "email@test.com"
	password := "test_password"

	// Act
	got := CreateUser(email, password)
	want := User{
		ID:       -1,
		Email:    email,
		Password: password,
	}

	// Assert
	assert.Equal(t, *got, want)
}
