package services

import (
    "be-test-concrete-ai/account-manager/models"
    "be-test-concrete-ai/account-manager/repositories"
    "errors"
    "golang.org/x/crypto/bcrypt"
    "github.com/dgrijalva/jwt-go"
    "time"
)

func CreateUser(user *models.User) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user.Password = string(hashedPassword)
    return repositories.CreateUser(user)
}

func LoginUser(email, password string) (string, error) {
    user, err := repositories.GetUserByEmail(email)
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
