package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/luraproject/lura/config"
	"github.com/luraproject/lura/logging"
	"github.com/luraproject/lura/proxy"
	"github.com/luraproject/lura/router"
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

	rotersConfig := readConfig()
	for _, path := range rotersConfig {
		router, err := parser.Parse(path)
		if err != nil {
			log.Fatal("Error:", err.Error())
		}
		serviceConfig.Endpoints = append(serviceConfig.Endpoints, router.Endpoints...)
	}
	serviceConfig.Port = *port
	engine := gin.Default()

	logger, _ := logging.NewLogger(*logLevel, os.Stdout, "[LURA]")
	routerFactory := lura.NewFactory(lura.Config{
		Engine:         engine,
		Middlewares:    []gin.HandlerFunc{emptyMiddleware()}, // if engin.Use(Middleware) this is not work
		HandlerFactory: lura.EndpointHandler,
		ProxyFactory:   proxy.DefaultFactory(logger),
		Logger:         logger,
		RunServer:      router.RunServer,
	})
	// routerFactory := lura.DefaultFactory(proxy.DefaultFactory(logger), logger)
	routerFactory.New().Run(serviceConfig)
}

func emptyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("emptyMiddle work")
		c.Next()
	}
}
