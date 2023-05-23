package packet

import (
	"encoding/binary"
	"errors"
	"net"
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

func ReadOpenConnectionRequest1(data []byte) (*OpenConnectionRequest1, error) {
	if len(data) < 22 {
		return nil, errors.New("неверный размер пакета OpenConnectionRequest1")
	}

	var req1 OpenConnectionRequest1
	copy(req1.Magic[:], data[1:17])
	req1.Protocol = data[17]
	nullPayloadEnd := 18
	for ; nullPayloadEnd < len(data) && data[nullPayloadEnd] == 0x00; nullPayloadEnd++ {
	}
	req1.NullPayload = data[18:nullPayloadEnd]
	/*
		if nullPayloadEnd+8 > len(data) {
			return nil, errors.New("неправильный формат пакета OpenConnectionRequest1")
		}

		Вообще, это должно было работать как проверка целостоности, однако оно почему то не работает. Оставлю это здесь, может понадобится потом.

	*/
	req1.ClientGUID = int64(binary.BigEndian.Uint64(data[nullPayloadEnd : nullPayloadEnd+8]))
	req1.MTU = int16(len(data) - (nullPayloadEnd + 8))

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

// OpenConnectionRequest2 структура пакета OpenConnectionRequest2.
type OpenConnectionRequest2 struct {
	Magic      [16]byte
	ServerAddr net.UDPAddr
	ClientMTU  int16
	ClientGUID int64
}

type OpenConnectionReply2 struct {
	Magic      [16]byte
	ServerGUID int64
	ClientAddr net.UDPAddr
	MTU        int16
	Security   bool
}

// ReadOpenConnectionRequest2 читает и анализирует данные пакета OpenConnectionRequest2.
func ReadOpenConnectionRequest2(data []byte) (*OpenConnectionRequest2, error) {
	if len(data) < 28 {
		return nil, errors.New("неверный размер пакета OpenConnectionRequest2")
	}

	var req2 OpenConnectionRequest2
	copy(req2.Magic[:], data[1:17])
	req2.ServerAddr.IP = net.IPv4(data[17], data[18], data[19], data[20])
	req2.ServerAddr.Port = int(binary.BigEndian.Uint16(data[21:23]))
	req2.ClientMTU = int16(binary.BigEndian.Uint16(data[23:25]))
	req2.ClientGUID = int64(binary.BigEndian.Uint64(data[25:33]))

	return &req2, nil
}

func WriteOpenConnectionReply2(reply2 *OpenConnectionReply2) []byte {
	data := make([]byte, 34)

	data[0] = 0x08
	copy(data[1:17], reply2.Magic[:])
	binary.BigEndian.PutUint64(data[17:25], uint64(reply2.ServerGUID))
	data[25] = byte(reply2.ClientAddr.IP[0])
	data[26] = byte(reply2.ClientAddr.IP[1])
	data[27] = byte(reply2.ClientAddr.IP[2])
	data[28] = byte(reply2.ClientAddr.IP[3])
	binary.BigEndian.PutUint16(data[29:31], uint16(reply2.ClientAddr.Port))
	binary.BigEndian.PutUint16(data[31:33], uint16(reply2.MTU))
	data[33] = 0x00 // Security flag (0 = no security, 1 = security enabled)
	if reply2.Security {
		data[33] = 0x01
	}

	return data
}
