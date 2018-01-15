package server

import (
	//"io"
	"log"
	"net"
)

func (p *RedisServer) Handle(clientConn net.Conn) {
	log.Println("TCP: new client", clientConn.RemoteAddr())
	/*
		// The client should initialize itself by sending a 4 byte sequence indicating
		// the version of the protocol that it intends to communicate, this will allow us
		// to gracefully upgrade the protocol away from text/line oriented to whatever...
		buf := make([]byte, 4)
		_, err := io.ReadFull(clientConn, buf)
		if err != nil {
			log.Println("failed to read protocol version - %s", err)
			return
		}
		protocolMagic := string(buf)

		log.Println("CLIENT(%s): desired protocol magic '%s'", clientConn.RemoteAddr(), protocolMagic)

		var prot protocol
		switch protocolMagic {
		case "  V2":
			prot = protocol{ctx: &context{ctx: p}}
		default:
			prot.SendFramedResponse(clientConn, []byte("E_BAD_PROTOCOL"))
			clientConn.Close()
			log.Println("client(%s) bad protocol magic '%s'", clientConn.RemoteAddr(), protocolMagic)
			return
		}
	*/
	prot := protocol{ctx: &context{ctx: p}}
	err := prot.IOLoop(clientConn)
	if err != nil {
		log.Println("client(%s) - %s", clientConn.RemoteAddr(), err)
		return
	}
}
