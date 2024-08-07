package services

import (
	"be-test-concrete-ai/account-manager/models"
	"be-test-concrete-ai/account-manager/prisma/db"
	"be-test-concrete-ai/account-manager/repositories"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user *models.User) (*db.UserModel, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
        return nil, err
    }
    user.Password = string(hashedPassword)
	createdUser, errCreateUser := repositories.CreateUser(user)

	if errCreateUser != nil {
		println("errCreateUser:", errCreateUser)
		return nil, errCreateUser
	}
	return createdUser, nil
}

func UserIsExistsByUsername(username string) (bool) {
	return repositories.UserIsExistsByUsername(username)	
}

func LoginUser(username, password string) (string, error) {
	user, err := repositories.GetUserByUsername(username)
	if err == db.ErrNotFound {
		return "", errors.New("invalid credentials")
	}
	if err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})
	return token.SignedString([]byte("secret"))
}
