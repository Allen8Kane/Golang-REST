package main

import (
	"flag"
	"log"

	"github.com/Allen8Kane/Golang-REST/internal/app/apiserver"
	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init()  {
	flag.StringVar(&configPath, "config-path","configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}