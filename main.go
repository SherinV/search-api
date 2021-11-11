package main

import (
	"flag"

	"github.com/SherinV/search-api/config"
	"github.com/SherinV/search-api/database"
	"github.com/SherinV/search-api/server"

	klog "k8s.io/klog/v2"
)

// const defaultPort = "8080"

func main() {
	// Initialize the logger.
	klog.InitFlags(nil)
	flag.Parse()
	defer klog.Flush()
	klog.Info("Starting search-api.")

	// Read the config from the environment.
	config := config.New()
	config.PrintConfig()
	database.GetConnection()
	server.StartAndListen()
}
