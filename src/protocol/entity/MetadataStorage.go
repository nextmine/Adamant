package entity

/*

  ___      _                             _
 / _ \    | |                           | |
/ /_\ \ __| | __ _ _ __ ___   __ _ _ __ | |_
|  _  |/ _` |/ _` | '_ ` _ \ / _` | '_ \| __|
| | | | (_| | (_| | | | | | | (_| | | | | |_
\_| |_/\__,_|\__,_|_| |_| |_|\__,_|_| |_|\__|

Adamant - высокопроизводительное и надежное ядро сервера Minecraft Bedrock Edition написанное на Golang
Автор: dm-vev(vk.com/wideputhin; t.me/dmvev)
25.07.23

*/

import "fmt"

const (
	TypeByte = iota
	TypeShort
	TypeInt
	TypeFloat
	TypeString
	TypeCompound
	TypeVector3Int
	TypeLong
	TypeVector3Float

	KeyFlags         = 0
	KeyFlagsExtended = 1
)

type Vector3Int struct {
	X int
	Y int
	Z int
}

type Vector3Float struct {
	X float32
	Y float32
	Z float32
}

type Metadata struct {
	metadata map[int]map[string]interface{}
}

func NewMetadata() *Metadata {
	return &Metadata{
		metadata: make(map[int]map[string]interface{}),
	}
}

func (m *Metadata) GetEntry(key int) map[string]interface{} {
	return m.metadata[key]
}

func (m *Metadata) SetEntry(key int, value interface{}, entryType int) {
	m.metadata[key] = map[string]interface{}{
		"value": value,
		"type":  entryType,
	}
}

func (m *Metadata) GetByte(key int) int {
	entry := m.GetEntry(key)
	if entry != nil && entry["type"] == TypeByte {
		return entry["value"].(int)
	}
	return 0
}

func (m *Metadata) SetByte(key int, value int) {
	m.SetEntry(key, value, TypeByte)
}

func (m *Metadata) GetShort(key int) int {
	entry := m.GetEntry(key)
	if entry != nil && entry["type"] == TypeShort {
		return entry["value"].(int)
	}
	return 0
}

func (m *Metadata) SetShort(key int, value int) {
	m.SetEntry(key, value, TypeShort)
}

func (m *Metadata) GetInt(key int) int {
	entry := m.GetEntry(key)
	if entry != nil && entry["type"] == TypeInt {
		return entry["value"].(int)
	}
	return 0
}

func (m *Metadata) SetInt(key int, value int) {
	m.SetEntry(key, value, TypeInt)
}

func (m *Metadata) GetFloat(key int) float32 {
	entry := m.GetEntry(key)
	if entry != nil && entry["type"] == TypeFloat {
		return entry["value"].(float32)
	}
	return 0
}

func (m *Metadata) SetFloat(key int, value float32) {
	m.SetEntry(key, value, TypeFloat)
}

func (m *Metadata) GetString(key int) string {
	entry := m.GetEntry(key)
	if entry != nil && entry["type"] == TypeString {
		return entry["value"].(string)
	}
	return ""
}

func (m *Metadata) SetString(key int, value string) {
	m.SetEntry(key, value, TypeString)
}

func (m *Metadata) GetCompound(key int) interface{} {
	entry := m.GetEntry(key)
	if entry != nil && entry["type"] == TypeCompound {
		return entry["value"]
	}
	return nil
}

func (m *Metadata) SetCompound(key int, value interface{}) {
	m.SetEntry(key, value, TypeCompound)
}

func (m *Metadata) GetVector3Int(key int) Vector3Int {
	entry := m.GetEntry(key)
	if entry != nil && entry["type"] == TypeVector3Int {
		return entry["value"].(Vector3Int)
	}
	return Vector3Int{}
}

func (m *Metadata) SetVector3Int(key int, value Vector3Int) {
	m.SetEntry(key, value, TypeVector3Int)
}

func (m *Metadata) GetLong(key int) int64 {
	entry := m.GetEntry(key)
	if entry != nil && entry["type"] == TypeLong {
		return entry["value"].(int64)
	}
	return 0
}

func (m *Metadata) SetLong(key int, value int64) {
	m.SetEntry(key, value, TypeLong)
}

func (m *Metadata) GetVector3Float(key int) Vector3Float {
	entry := m.GetEntry(key)
	if entry != nil && entry["type"] == TypeVector3Float {
		return entry["value"].(Vector3Float)
	}
	return Vector3Float{}
}

func (m *Metadata) SetVector3Float(key int, value Vector3Float) {
	m.SetEntry(key, value, TypeVector3Float)
}

func (m *Metadata) GetFlag(flag int, extended bool) bool {
	var flags int64
	if !extended {
		flags = m.GetLong(KeyFlags)
	} else {
		flags = m.GetLong(KeyFlagsExtended)
	}
	if flags != 0 {
		return flags&(1<<uint(flag)) > 0
	}
	return false
}

func (m *Metadata) SetFlag(flag int, value bool, extended bool) {
	currentValue := m.GetFlag(flag, extended)
	if currentValue == value {
		return
	}

	var flags int64
	if !extended {
		flags = m.GetLong(KeyFlags)
	} else {
		flags = m.GetLong(KeyFlagsExtended)
	}

	if value {
		flags |= 1 << uint(flag)
	} else {
		flags &^= 1 << uint(flag)
	}

	if !extended {
		m.SetLong(KeyFlags, flags)
	} else {
		m.SetLong(KeyFlagsExtended, flags)
	}
}

func main() {
	m := NewMetadata()

	m.SetByte(10, 255)
	fmt.Println(m.GetByte(10))

	m.SetFlag(5, true, false)
	fmt.Println(m.GetFlag(5, false))
}
