package main

import (
	"PaymentAPI/pkg/handler"
	"PaymentAPI/pkg/repository"
	"PaymentAPI/pkg/service"
	"fmt"
)

func main() {
	postgres, err := repository.NewPostgresdb()
	if err != nil {
		fmt.Println("Connection to db failed")
		return
	}
	defer postgres.Close()

	repo := repository.NewRepository(postgres)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	err = handlers.InitRoutes().Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
