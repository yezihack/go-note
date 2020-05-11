package main

import "github.com/yezihack/studyGo/org/service/services"

func main() {
	var engine services.AppService

	engineFlag := services.EngineRedis
	engine = services.NewService(engineFlag)
	engine.Find()

	engine = services.NewService(engineFlag)
	engine.Find()

	engine = services.NewService(engineFlag)
	engine.Find()
}
