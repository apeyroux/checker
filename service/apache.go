package service

import (
	"fmt"
	"log"
	"net"
)

type ApacheService struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func (as *ApacheService) State() state {
	uri := fmt.Sprintf("%s:%d", as.Host, as.Port)
	cnx, err := net.Dial("tcp", uri)
	defer cnx.Close()
	if err != nil {
		log.Printf("apache err:%s", err)
		return DOWN
	}
	return ALIVE
}

func (as *ApacheService) Name() string {
	return "apache"
}

func (as *ApacheService) String() string {
	str := fmt.Sprintf("%s %s:%d", as.Name(), as.Host, as.Port)
	return str
}
