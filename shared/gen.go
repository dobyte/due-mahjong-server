package main

import (
	"due-mahjong-server/shared/model/mail"
	"due-mahjong-server/shared/model/user"
	"github.com/dobyte/gen-mongo-dao"
	"log"
)

func main() {
	g := gen.NewGenerator(&gen.Options{
		OutputDir:    "./dao",
		OutputPkg:    "due-mahjong-server/shared/dao",
		EnableSubPkg: true,
	})

	g.AddModels(
		&user.User{},
		&mail.Mail{},
	)

	err := g.MakeDao()
	if err != nil {
		log.Fatalf("generate code failed: %v", err)
	}

	log.Printf("generate code success")
}
