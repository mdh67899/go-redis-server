package server

import (
	"log"
	"net"
	"runtime"
	"strings"
)

type TCPHandler interface {
	Handle(net.Conn)
}

func TCPServer(listener net.Listener, handler TCPHandler) {
	log.Println("TCP: listening on", listener.Addr())

	for {
		clientConn, err := listener.Accept()
		if err != nil {
			if nerr, ok := err.(net.Error); ok && nerr.Temporary() {
				log.Println("temporary Accept() failure -", err)
				runtime.Gosched()
				continue
			}
			// theres no direct way to detect this error because it is not exposed
			if !strings.Contains(err.Error(), "use of closed network connection") {
				log.Println("listener.Accept() -", err)
			}
			break
		}
		go handler.Handle(clientConn)
	}

	log.Println("TCP: closing", listener.Addr())
}
