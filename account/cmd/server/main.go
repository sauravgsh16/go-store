package main

import (
	"context"
	"log"

	"github.com/sauravgsh16/go-store/account/app"
	"github.com/sauravgsh16/go-store/account/datastore/postgres"
	"github.com/sauravgsh16/go-store/account/domain/account"
	"github.com/sauravgsh16/go-store/account/service"
)

// RunServer runs server
func RunServer() {
	ctx := context.Background()

	db := postgres.GetDBConn()
	repo := account.NewAccountRepo(db)
	defer repo.Close()

	s := service.NewAccService(repo)

	log.Println("Listening on port: 8080....")
	log.Fatal(app.RunGRPC(ctx, s, "8080"))
}

func main() {
	RunServer()
}
