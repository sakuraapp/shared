package resource

import (
	"github.com/sakuraapp/shared/resource/opcode"
	"gopkg.in/guregu/null.v4"
)

type Packet struct {
	Opcode opcode.Opcode `json:"op"`
	Data interface{} `json:"d"`
	Time null.Int `json:"t"`
}

func (p *Packet) DataMap() map[string]interface{} {
	return p.Data.(map[string]interface{})
}