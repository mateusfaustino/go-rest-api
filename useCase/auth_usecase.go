// useCase/auth_usecase.go
package usecase

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/mateusfaustino/go-rest-api/repository"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthUseCase struct {
	UserRepo repository.UserRepository
	Secret   string
}

func NewAuthUseCase(userRepo repository.UserRepository, secret string) AuthUseCase {
	return AuthUseCase{
		UserRepo: userRepo,
		Secret:   secret,
	}
}

func (au *AuthUseCase) Login(username, password string) (string, error) {
	user, err := au.UserRepo.GetUserByUsername(username)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := au.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (au *AuthUseCase) GenerateToken(userID int) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // 1 dia de expiração

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(au.Secret))
}
