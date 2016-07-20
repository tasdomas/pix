package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/tasdomas/pix/config"
	"github.com/tasdomas/pix/storage"
	"github.com/tasdomas/pix/ui"
)

var cfgFile = flag.String("cfg", "", "configuration file")

func main() {
	flag.Parse()

	cfg, err := config.Load(*cfgFile)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	st, err := storage.New(cfg.Storage)
	if err != nil {
		log.Fatalf("failed to create disk storage: %v", err)
	}
	uiServer, err := ui.NewServer(st, cfg.Name, cfg.Secret)
	if err != nil {
		log.Fatalf("failed to load UI server: %v", err)
	}

	log.Fatal(http.ListenAndServe(cfg.Port, uiServer))
}
