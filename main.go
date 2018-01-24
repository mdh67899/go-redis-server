package main

import (
	"log"
	"net"
	"os"
	"runtime"

	"github.com/mdh67899/go-redis-server/server"
)

var Redis *server.RedisServer = &server.RedisServer{}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	tcpListener, err := net.Listen("tcp", ":8003")
	if err != nil {
		log.Println("listen (:8003) failed - %s", err)
		os.Exit(1)
	}
	Redis.Lock()
	Redis.TCPListener = tcpListener
	Redis.Unlock()

	Redis.Add(1)
	go func(redis *server.RedisServer) {
		defer redis.Done()
		//tcpServer, 从tcpListener收到连接, 交给tcpServer的handler方法处理
		server.TCPServer(redis.TCPListener, redis)
	}(Redis)

	Redis.Wait()
}
