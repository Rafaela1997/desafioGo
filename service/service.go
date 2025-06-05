package service

import (
	"desafioGo/database"
	"desafioGo/model"
	"github.com/go-resty/resty/v2"
	"log"
)

// Estrutura usada para mapear os dados da API externa
type apiFruit struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Nutritions struct {
		Sugar float64 `json:"sugar"`
	} `json:"nutritions"`
}

func CrawlFruits() {
	log.Println("Iniciando busca de dados da API externa...")

	client := resty.New()
	resp, err := client.R().
		SetResult(&[]apiFruit{}).
		Get("https://www.fruityvice.com/api/fruit/all")

	if err != nil {
		log.Println("Erro ao consumir API externa:", err)
		return
	}

	fruits := *resp.Result().(*[]apiFruit)

	for _, apiF := range fruits {
		fruit := model.Fruit{
			ID:    apiF.ID,
			Name:  apiF.Name,
			Sugar: int(apiF.Nutritions.Sugar), // Armazena como inteiro
		}

		// Salva no banco (atualiza se já existir, senão cria)
		database.DB.Save(&fruit)
	}

	log.Printf("Crawler finalizado: %d frutas processadas.\n", len(fruits))
}
