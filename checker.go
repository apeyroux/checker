package main

import (
	"encoding/json"
	"fmt"
	"github.com/j4/checker/service"
	"io/ioutil"
	"log"
)

type Checker struct {
	Name     string            `json:"name"`
	Services []service.Service `json:"services"`
}

type Config struct {
	Checkers []Checker `json:"checkers"`
}

func dispacher(aservice service.Service) service.Serve {
	switch aservice.Name {
	case "apache":
		cfg := new(service.ApacheService)
		cfgjson, _ := json.Marshal(aservice.Config)
		json.Unmarshal(cfgjson, &cfg)
		aservice.Config = cfg
		return cfg
	case "memcache":
		cfg := new(service.MemCache)
		cfgjson, _ := json.Marshal(aservice.Config)
		json.Unmarshal(cfgjson, &cfg)
		aservice.Config = cfg
		return cfg
	}
	return nil
}

func main() {
	cfgbuf, err := ioutil.ReadFile("checker.json")
	if err != nil {
		log.Fatal(err)
	}
	cfg := new(Config)
	err = json.Unmarshal(cfgbuf, &cfg)
	if err != nil {
		log.Fatal(err)
	}
	for _, checker := range cfg.Checkers {
		fmt.Printf("--- %s ---\n", checker.Name)
		for _, servi := range checker.Services {
			serve := dispacher(servi)
			log.Printf("%s %s", serve.String(), service.StateToString(serve.State()))
		}

	}
}
