package main

import (
	"flag"
	"log"
	"os"

	"github.com/luraproject/lura/config"
	"github.com/luraproject/lura/logging"
	"github.com/luraproject/lura/proxy"
	lura "github.com/luraproject/lura/router/gin"
)

func main() {
	port := flag.Int("p", 8080, "Port of the service")
	logLevel := flag.String("l", "ERROR", "Logging Level")

	configFile := flag.String("c", "lura/configs/config.json", "Path to config filename")

	parser := config.NewParser()
	serviceConfig, err := parser.Parse(*configFile)

	if err != nil {
		log.Fatal("Error:", err.Error())
	}

	serviceConfig.Port = *port
	// serviceConfig.Endpoints = make([]*config.EndpointConfig, 0)

	logger, _ := logging.NewLogger(*logLevel, os.Stdout, "[LURA]")
	routerFactory := lura.DefaultFactory(proxy.DefaultFactory(logger), logger)
	routerFactory.New().Run(serviceConfig)

}
