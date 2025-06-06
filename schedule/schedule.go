package schedule

import (
	"desafioGo/service"
	"github.com/robfig/cron/v3"
	"log"
)

func Start() {
	c := cron.New()


	_, err := c.AddFunc("0 0 * * *", func() {
		log.Println("Executando crawler programado...")
		service.CrawlFruits()
	})

	if err != nil {
		log.Println("Erro ao agendar crawler:", err)
	}

	go service.CrawlFruits()

	c.Start()
}
