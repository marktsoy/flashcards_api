package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/marktsoy/flashcards_api/internal/app/apiserver"
)

var (
	configPath string
)

// init function - called once
// More info: https://tutorialedge.net/golang/the-go-init-function/#:~:text=The%20init%20Function,will%20only%20be%20called%20once.
func init() {
	fmt.Println("Init function. I am called only once. I am called first")

	// Flag (cmd args) definitions
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "Path to config file")
}

func main() {

	// After all flags are defined call the Parse function to bind args to vars
	flag.Parse()
	var config *apiserver.Config = apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)

	log.Println(config.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
