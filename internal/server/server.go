package server

import (
	"Adamant/internal/packet"
	"log"
	"net"
)

type Server struct {
	conn *net.UDPConn
}

func New(port string) (*Server, error) {
	serverAddr, err := net.ResolveUDPAddr("udp", ":"+port)
	if err != nil {
		return nil, err
	}

	conn, err := net.ListenUDP("udp", serverAddr)
	if err != nil {
		return nil, err
	}

	return &Server{conn: conn}, nil
}

func (s *Server) Run() {
	defer s.conn.Close()

	buf := make([]byte, 2048)

	for {
		n, addr, err := s.conn.ReadFromUDP(buf)
		if err != nil {
			log.Println("Ошибка при чтении данных:", err)
			continue
		}

		go packet.Process(s.conn, addr, buf[:n])
	}
}
