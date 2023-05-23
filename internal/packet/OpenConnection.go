package packet

import (
	"encoding/binary"
	"errors"
)

// ... (предыдущие структуры и функции)

type OpenConnectionRequest1 struct {
	Magic       [16]byte
	Protocol    byte
	NullPayload []byte
	ClientGUID  int64
	MTU         int16
}

type OpenConnectionReply1 struct {
	Magic      [16]byte
	ServerGUID int64
	Security   bool
	MTU        int16
}

// ReadOpenConnectionRequest1 - функция для чтения пакета OpenConnectionRequest1
func ReadOpenConnectionRequest1(data []byte) (*OpenConnectionRequest1, error) {
	if len(data) < 22 {
		return nil, errors.New("неверный размер пакета OpenConnectionRequest1")
	}

	var req1 OpenConnectionRequest1
	copy(req1.Magic[:], data[1:17])
	req1.Protocol = data[17]
	nullPayloadEnd := 18
	for ; data[nullPayloadEnd] == 0x00; nullPayloadEnd++ {
	}
	req1.NullPayload = data[18:nullPayloadEnd]
	req1.ClientGUID = int64(binary.BigEndian.Uint64(data[nullPayloadEnd : nullPayloadEnd+8]))
	req1.MTU = int16(binary.BigEndian.Uint16(data[nullPayloadEnd+8 : nullPayloadEnd+10]))

	return &req1, nil
}

// WriteOpenConnectionReply1 - функция для записи пакета OpenConnectionReply1
func WriteOpenConnectionReply1(reply1 *OpenConnectionReply1) []byte {
	data := make([]byte, 28)

	data[0] = 0x06
	copy(data[1:17], reply1.Magic[:])
	binary.BigEndian.PutUint64(data[17:25], uint64(reply1.ServerGUID))
	data[25] = 0x00 // Security flag (0 = no security, 1 = security enabled)
	binary.BigEndian.PutUint16(data[26:28], uint16(reply1.MTU))

	return data
}
