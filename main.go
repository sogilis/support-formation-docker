package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/sogilis/support-formation-docker/handlers"
)

func main() {
	cfg, err := GetConfigFromCLI()
	if err != nil {
		log.Fatal("fail to fetch info from CLI", err)
	}

	log.Info("Storage dir location : ", cfg.StoragePath)
	if err = os.MkdirAll(cfg.StoragePath, 0700); err != nil {
		log.Fatal("Unable to create dir with error ", err)
	}

	log.Info("Creating routes")
	if err = handlers.Launch(handlers.Create_router(cfg.StoragePath), cfg.Port); err != nil {
		log.Fatal("Crash with error: ", err)
	}
}
