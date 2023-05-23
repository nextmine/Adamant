package packet

import (
	"Adamant/internal"
	"fmt"
	"log"
	"net"
)

func Process(conn *net.UDPConn, addr *net.UDPAddr, data []byte) {
	if len(data) == 0 {
		return
	}

	packetID := data[0]
	switch packetID {
	case 0x01: // UnconnectedPing
		ping, err := ReadUnconnectedPing(data)
		if err != nil {
			log.Println("Ошибка при чтении UnconnectedPing:", err)
			return
		}

		pong := &UnconnectedPong{
			PingTime:   ping.Time,
			ServerID:   ping.ClientID + 12345, // Взять реальный ServerID из конфигурации сервера
			Magic:      ping.Magic,
			ServerName: fmt.Sprintf("MCPE;%s;%d;%s;0;20;0;Bedrock level;Creative;", internal.ServerName, internal.ProtocolVersion, internal.GameVersion),
		}

		response := WriteUnconnectedPong(pong)
		_, err = conn.WriteToUDP(response, addr)
		if err != nil {
			log.Println("Ошибка при отправке UnconnectedPong:", err)
		}
	case 0x05: // OpenConnectionRequest1
		req1, err := ReadOpenConnectionRequest1(data)
		if err != nil {
			log.Println("Ошибка при чтении OpenConnectionRequest1:", err)
			return
		}
		log.Println("Подключается новый клиент")
		log.Printf("Протокол клиента: %d\n", req1.Protocol)

		reply1 := &OpenConnectionReply1{
			Magic:      req1.Magic,
			ServerGUID: req1.ClientGUID + 12345, // Взять реальный ServerGUID из конфигурации сервера
			Security:   false,
			MTU:        req1.MTU,
		}

		response := WriteOpenConnectionReply1(reply1)
		_, err = conn.WriteToUDP(response, addr)
		if err != nil {
			log.Println("Ошибка при отправке OpenConnectionReply1:", err)
		}
	case 0x07: // OpenConnectionRequest2
		req2, err := ReadOpenConnectionRequest2(data)
		if err != nil {
			log.Println("Ошибка при чтении OpenConnectionRequest2:", err)
			return
		}

		reply2 := &OpenConnectionReply2{
			Magic:      req2.Magic,
			ServerGUID: req2.ClientGUID + 12345, // Взять реальный ServerGUID из конфигурации сервера
			ClientAddr: *addr,
			MTU:        1500,
			Security:   false,
		}

		clientIP := addr.IP
		clientPort := addr.Port
		log.Printf("IP клиента: %s, порт клиента: %d", clientIP.String(), clientPort)

		response := WriteOpenConnectionReply2(reply2)
		_, err = conn.WriteToUDP(response, addr)
		if err != nil {
			log.Println("Ошибка при отправке OpenConnectionReply2:", err)
		}
	default:
		log.Printf("Неизвестный пакет с ID 0x%02x\n", packetID)
	}
}
