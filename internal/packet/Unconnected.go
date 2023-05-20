package packet

import (
	"encoding/binary"
	"errors"
)

// UnconnectedPing - структура для пакета UnconnectedPing
type UnconnectedPing struct {
	Time     int64
	Magic    [16]byte
	ClientID int64
}

// UnconnectedPong - структура для пакета UnconnectedPong
type UnconnectedPong struct {
	PingTime   int64
	ServerID   int64
	Magic      [16]byte
	ServerName string
}

// ReadUnconnectedPing - функция для чтения UnconnectedPing пакета
func ReadUnconnectedPing(data []byte) (*UnconnectedPing, error) {
	if len(data) < 25 {
		return nil, errors.New("неверный размер пакета UnconnectedPing")
	}

	var ping UnconnectedPing
	ping.Time = int64(binary.BigEndian.Uint64(data[1:9]))
	copy(ping.Magic[:], data[9:25])
	ping.ClientID = int64(binary.BigEndian.Uint64(data[25:]))

	return &ping, nil
}

// WriteUnconnectedPong - функция для записи UnconnectedPong пакета
func WriteUnconnectedPong(pong *UnconnectedPong) []byte {
	data := make([]byte, 35+len(pong.ServerName))

	data[0] = 0x1c
	binary.BigEndian.PutUint64(data[1:9], uint64(pong.PingTime))
	binary.BigEndian.PutUint64(data[9:17], uint64(pong.ServerID))
	copy(data[17:33], pong.Magic[:])
	data[33] = 0x00
	data[34] = byte(len(pong.ServerName))
	copy(data[35:], []byte(pong.ServerName))

	return data
}
