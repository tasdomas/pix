package main

import (
	"flag"
	"log"

	"github.com/tasdomas/pixserver/config"
)

var cfgFile = flag.String("cfg", "", "configuration file")

func main() {
	flag.Parse()

	cfg, err := config.Load(*cfgFile)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
}
