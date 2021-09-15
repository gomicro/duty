package main

import (
	"net/http"
	"os"

	"github.com/gomicro/duty/config"
	log "github.com/gomicro/ledger"
)

var (
	conf *config.File
)

func configure() {
	c, err := config.ParseFromFile()
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err.Error())
		os.Exit(1)
	}

	conf = c
	log.Debug("Config file parsed")

	log.Debug("Configuration complete")
}

func main() {
	configure()

	log.Infof("Listening on %v:%v", "0.0.0.0", "4567")
	err := http.ListenAndServe("0.0.0.0:4567", conf)
	if err != nil {
		log.Errorf("server error: %v", err.Error())
	}
}
