package main 

import (
	"github.com/YaguarEgor/gator_cli/internal/config"
	"log"
	"fmt"
)

const username = "Egor"

func main() {
	conf, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", conf)
	err = conf.SetUser(username)
	if err != nil {
		log.Fatalf("error setting username in config: %v", err)
	}
	conf, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config again: %v", err)
	}
	fmt.Printf("Read config again: %+v\n", conf)

}