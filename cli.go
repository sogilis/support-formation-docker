package main

import (
	"flag"
	"os"
	"path"

	"github.com/google/uuid"
)

func GetConfigFromCLI() (*Config, error) {
	config := Config{}

	dir := path.Join(os.TempDir(), uuid.New().String())

	flag.UintVar(&config.Port, "port", 7777, "HTTP server port used to listen")
	flag.StringVar(&config.StoragePath, "storage-path", dir, "Image storage folder")

	flag.Parse()

	return &config, nil
}
