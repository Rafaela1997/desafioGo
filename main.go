package main

import (
	"desafioGo/database"
	"desafioGo/route"
	"desafioGo/schedule"
)


func main() {
	database.Init()        
	schedule.Start()     
	r := route.SetupRouter()
	r.Run(":8080")         
}