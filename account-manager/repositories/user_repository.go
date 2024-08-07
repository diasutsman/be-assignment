package repositories

import (
	"be-test-concrete-ai/account-manager/config"
	"be-test-concrete-ai/account-manager/models"
	"be-test-concrete-ai/account-manager/prisma/db"
	"context"
	"fmt"
)

func CreateUser(user *models.User) (*db.UserModel, error) {

	return config.Client.User.CreateOne(
		db.User.Username.Set(user.Username),
		db.User.Password.Set(user.Password),
	).Exec(context.Background())
}

func GetUserByUsername(username string) (*db.UserModel, error) {

	fmt.Println("username:", username)

	user, err := config.Client.User.FindUnique(
		db.User.Username.Equals(username),
	).Exec(context.Background())

	if err != nil {
		return nil, err
	}

	return user, nil
}

func UserIsExistsByUsername(username string) (bool) {

	user, _ := GetUserByUsername(username)

	
	return user != nil
}

func GetUserByID(id int) (*db.UserModel, error) {
	fmt.Println("id:", id)

	user, err := config.Client.User.FindUnique(
		db.User.ID.Equals(id),
	).Exec(context.Background())

	if err != nil {
		return nil, err
	}

	return user, nil
}
