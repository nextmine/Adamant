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
	default:
		log.Printf("Неизвестный пакет с ID 0x%02x\n", packetID)
	}
}
