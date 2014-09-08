package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/communaute-cimi/checker/service"
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

var (
	cfg *string = flag.String("config", "checker.json", "config file")
)

func dispacher(aservice service.Service) service.Serve {
	switch aservice.Name {
	case "apache":
		cfg := new(service.ApacheService)
		cfgjson, _ := json.Marshal(aservice.Config)
		json.Unmarshal(cfgjson, &cfg)
		aservice.Config = cfg
		return cfg
	case "memcache":
		cfg := new(service.MemcacheService)
		cfgjson, _ := json.Marshal(aservice.Config)
		json.Unmarshal(cfgjson, &cfg)
		aservice.Config = cfg
		return cfg
	case "postgres":
		cfg := new(service.PostgresService)
		cfgjson, _ := json.Marshal(aservice.Config)
		json.Unmarshal(cfgjson, &cfg)
		aservice.Config = cfg
		return cfg
	default:
		return nil
	}
}

func main() {

	flag.Parse()

	cfgbuf, err := ioutil.ReadFile(*cfg)
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
			if serve != nil {
				log.Printf("%s %s", serve.String(), service.StateToString(serve.State()))
			} else {
				log.Printf("%s is not define. Check your config.", servi.Name)
			}
		}

	}
}
