package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ThiagoVicentini/go_jobs/config"
	"github.com/ThiagoVicentini/go_jobs/router"
)

var (
	logger *config.Logger
)

func main() {
	port := ":8080"
	logger = config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	flag.Parse()
	argv := flag.Args()
	argc := len(argv)
	if argc > 1 {
		fmt.Fprintf(os.Stderr, "usage: ./main <port>\n")
		flag.PrintDefaults()
		os.Exit(2)
	} else if argc == 1 {
		port = argv[0]
		port = ":" + port
	}

	router.Initialize(port)
}
