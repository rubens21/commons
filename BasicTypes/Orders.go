package BasicTypes

import (
	"encoding/json"
	"github.com/maketplay/commons/Physics"
	"github.com/pkg/errors"
	"fmt"
)

type Order struct {
	Type  OrderType       `json:"order"`
	Data  interface{}     `json:"data"`
}

type MoveOrderData struct {
	Velocity Physics.Velocity `json:"velocity"`
}

type KickOrderData struct {
	Velocity Physics.Velocity `json:"velocity"`
}


const (
	// orders sent by the PLAYER
	MOVE  OrderType = "MOVE"  // Move me to that position
	KICK  OrderType = "KICK"  // Cick the ball to that position
	CATCH OrderType = "CATCH" // I'll try to catch the ball of the player
)

func (o *Order) GetMoveOrderData() MoveOrderData {
	return o.Data.(MoveOrderData)
}

func (o *Order) GetKickOrderData() KickOrderData {
	return o.Data.(KickOrderData)
}

func (o *Order) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Type  OrderType       `json:"order"`
		Data  json.RawMessage     `json:"data"`
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}
	o.Type = tmp.Type
	switch {
	case tmp.Type == MOVE:
		var mov MoveOrderData
		err = json.Unmarshal(tmp.Data, &mov)
		o.Data = mov
	case tmp.Type == KICK:
		var mov KickOrderData
		err = json.Unmarshal(tmp.Data, &mov)
		o.Data = mov
	case tmp.Type == CATCH:
		o.Data = nil
	default:
		err = errors.New(fmt.Sprintf("Unknow order type %s", tmp.Type))
	}
	return err
}
