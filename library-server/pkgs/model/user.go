package model

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           string
	Username     string
	Email        string
	PasswordHash string
}

// setPassword will hash the provided password and add it to the UserModel
//		err := userModel.setPassword("password123")
func (u *User) setPassword(password string) error {
	if len(password) == 0 {
		return errors.New("password cannot be empty")
	}
	bytePassword := []byte(password)
	passwordHash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(passwordHash)
	return nil
}

// checkPassword compares provided password with the PasswordHash on the UserModel
// 		err := userModel.checkPassword("password1234")
func (u *User) checkPassword(password string) error {
	byteHashedPassword := []byte(u.PasswordHash)
	bytePassword := []byte(password)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}
