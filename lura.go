package main

import (
	"flag"
	"log"
	"os"

	"github.com/luraproject/lura/config"
	"github.com/luraproject/lura/logging"
	"github.com/luraproject/lura/proxy"
	"github.com/luraproject/lura/router/gin"
)

func serviceLura() {
	port := flag.Int("p", 0, "Port of the service")
	logLevel := flag.String("l", "ERROR", "Logging Level")

	configFile := flag.String("c", "/etc/lura/configuration.json", "Path to config filename")

	parser := config.NewParser()
	serviceConfig, err := parser.Parse(*configFile)

	if err != nil {
		log.Fatal("Error:", err.Error())
	}

	serviceConfig.Port = *port
	serviceConfig.Endpoints = make([]*config.EndpointConfig, 0)

	logger, _ := logging.NewLogger(*logLevel, os.Stdout, "[LURA]")
	routerFactory := gin.DefaultFactory(proxy.DefaultFactory(logger), logger)
	routerFactory.New().Run(serviceConfig)

}
