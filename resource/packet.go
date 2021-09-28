package resource

import (
	"github.com/sakuraapp/shared/resource/opcode"
	"gopkg.in/guregu/null.v4"
)

type Packet struct {
	Opcode opcode.Opcode `json:"op"`
	Data map[string]interface{} `json:"d"`
	Time null.Int `json:"t"`
}