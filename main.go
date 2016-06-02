package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/tasdomas/pix/config"
	"github.com/tasdomas/pix/ui"
)

var cfgFile = flag.String("cfg", "", "configuration file")

func main() {
	flag.Parse()

	cfg, err := config.Load(*cfgFile)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	uiServer, err := ui.NewUIServer()
	if err != nil {
		log.Fatalf("failed to load UI server: %v", err)
	}

	log.Fatal(http.ListenAndServe(cfg.Port, uiServer))
}
