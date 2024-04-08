package main

import (
	"fmt"

	"github.com/Polengolas/api-rest-go/database"
	"github.com/Polengolas/api-rest-go/models"
	"github.com/Polengolas/api-rest-go/routes"
)

func main() {

	models.Personalities = []models.Personality{
		{Id: 1, Name: "Name 1", History: "History 1"},
		{Id: 2, Name: "Name 2", History: "History 2"},
	}

	database.ConectDataBase()

	fmt.Println("Iniciando o servidor Rest com GO")
	routes.HandleRequest()
}
