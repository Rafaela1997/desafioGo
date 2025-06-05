package schedule

import (
	"desafioGo/service"
	"github.com/robfig/cron/v3"
	"log"
)

func Start() {
	c := cron.New()

	// Roda todo dia à meia-noite
	_, err := c.AddFunc("0 0 * * *", func() {
		log.Println("Executando crawler programado...")
		service.CrawlFruits()
	})

	if err != nil {
		log.Println("Erro ao agendar crawler:", err)
	}

	// Executa o crawler uma vez na inicialização também (opcional, mas útil)
	go service.CrawlFruits()

	c.Start()
}
