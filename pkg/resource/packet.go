package resource

import (
	"github.com/sakuraapp/shared/pkg/resource/opcode"
	"gopkg.in/guregu/null.v4"
	"time"
)

type Packet struct {
	Opcode opcode.Opcode `json:"op"`
	Data   interface{}   `json:"d"`
	Time null.Int `json:"t"`
}

func (p *Packet) DataMap() map[string]interface{} {
	return p.Data.(map[string]interface{})
}

func BuildPacket(op opcode.Opcode, data interface{}) *Packet {
	t := time.Now().UnixNano() / 1000000

	return &Packet{
		Opcode: op,
		Data: data,
		Time: null.IntFrom(t),
	}
}