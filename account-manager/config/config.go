package config

import (
	"be-test-concrete-ai/account-manager/prisma/db"
	"fmt"
)

var (
	Client *db.PrismaClient
)

func Init() error {
	Client = db.NewClient()
	println("Client:", Client)
	if err := Client.Prisma.Connect(); err != nil {
		fmt.Println("Init:", err)
		return err
	}

	return nil
}

func Disconnect() {

	if err := Client.Prisma.Disconnect(); err != nil {
		panic(fmt.Errorf("could not disconnect: %w", err))
	}

}
