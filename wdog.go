package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"

	"github.com/songjiayang/wdog/config"
	"github.com/songjiayang/wdog/process"
)

var (
	configFile string
)

func init() {
	flag.StringVar(&configFile, "c", "./config.json", "config file")
}

func main() {
	flag.Parse()
	processes, err := config.Load(configFile)
	if err != nil {
		panic(err)
	}

	for _, proc := range processes {
		go process.NewProcess(proc).Run()
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	select {
	case <-sigChan:
		fmt.Printf("stopping...\n")
		break
	}
}
