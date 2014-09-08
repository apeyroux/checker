package service

import (
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"net"
)

type MemcacheService struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func (mc *MemcacheService) State() state {
	uri := fmt.Sprintf("%s:%d", mc.Host, mc.Port)
	cnx, err := net.Dial("tcp", uri)
	if err != nil {
		return DOWN
	}
	defer cnx.Close()
	mcc := memcache.New(uri)
	mcitem := new(memcache.Item)
	mcitem.Key = "tmc"
	mcitem.Value = []byte("test memcache")
	err = mcc.Set(mcitem)
	if err != nil {
		return DOWN
	}
	err = mcc.Delete("tmc")
	if err != nil {
		return DOWN
	}
	return ALIVE
}

func (mc *MemcacheService) Name() string {
	return "memcache"
}

func (mc *MemcacheService) String() string {
	str := fmt.Sprintf("%s %s:%d", mc.Name(), mc.Host, mc.Port)
	return str
}
