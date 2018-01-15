package server

import (
	"net"
	"sync"
)

type context struct {
	ctx *RedisServer
}

type RedisServer struct {
	TCPListener net.Listener
	sync.RWMutex
	sync.WaitGroup
}
