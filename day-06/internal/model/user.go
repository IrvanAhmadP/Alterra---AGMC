package model

import (
	"os"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"size:150;not null"`
	Username string `json:"username" gorm:"not null; unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password,omitempty" gorm:"not null"`
}

// HashPassword is a method for struct User for Hashing password
func (u *User) HashPassword() {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(bytes)
}

// GenerateToken is a method for struct User for creating new jwt token
func (u *User) GenerateToken() (string, error) {
	var (
		JWT_SECRET_KEY = os.Getenv("JWT_SECRET_KEY")
	)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       u.ID,
		"name":     u.Name,
		"username": u.Username,
		"email":    u.Email,
	})

	tokenString, err := token.SignedString([]byte(JWT_SECRET_KEY))
	return tokenString, err
}
