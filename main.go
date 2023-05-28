package main

import (
	"battle_of_psychics/internal/api"
	"battle_of_psychics/internal/app"
	"battle_of_psychics/internal/def"
	"log"
)

func main() {
	cfg, err := def.NewConfig()
	if err != nil {
		log.Fatalln("new config: ", err)
	}

	l, err := def.NewLogger(cfg.LogLevel)
	if err != nil {
		log.Fatalln("new logger: ", err)
	}

	battleSvc := app.NewBattleServerService(l)

	server, err := api.NewServer(cfg, l, battleSvc)
	if err != nil {
		l.Fatalln("init rest server:", err)
	}

	err = server.Serve()
	if err != nil {
		l.Fatalln("rest api serve:", err)
	}

	l.Info("started server at: ", cfg.HTTPPort)
}
