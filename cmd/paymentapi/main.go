package main

import (
	"PaymentAPI/pkg/handler"
	"PaymentAPI/pkg/repository"
	"PaymentAPI/pkg/service"
	"fmt"
)

// @title           Your API Title
// @version         1.0
// @description     This is a sample API.
// @termsOfService  http://swagger.io/terms/
// @contact.name    API Support
// @contact.email   log1c05678@gmail.com
// @host            localhost:8080
// @BasePath        /api/v1
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
