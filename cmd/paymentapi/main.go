package main

import (
	"PaymentAPI/pkg/handler"
	"PaymentAPI/pkg/repository"
	"PaymentAPI/pkg/service"
	"fmt"

	_ "PaymentAPI/docs"
)

// @title           PaymentAPI
// @version         1.0
// @description     This is a payment transaction processing system
// @contact.name    log1c0
// @contact.email   log1c0@protonmail.com
// @host            localhost:8080
// @BasePath        /api/
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
