package main

import (
	"desafioGo/database"
	"desafioGo/route"
	"desafioGo/schedule"
)


func main() {
	database.Init()        // Inicializa o banco (abre conexão e migra)
	schedule.Start()      // Inicia o cron para rodar o crawler diariamente
	r := route.SetupRouter() // Cria o router (como configurar @RequestMapping)
	r.Run(":8080")         // Sobe a API no localhost:8080
}