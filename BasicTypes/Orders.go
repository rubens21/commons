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

type OrderMove struct {
	Velocity Physics.Velocity `json:"velocity"`
}

type OrderKick struct {
	Velocity Physics.Velocity `json:"velocity"`
}
type OrderCatch struct {
}


const (
	// orders sent by the PLAYER
	MOVE  OrderType = "MOVE"  // Move me to that position
	KICK  OrderType = "KICK"  // Cick the ball to that position
	CATCH OrderType = "CATCH" // I'll try to catch the ball of the player
)

func (o *Order) GetOrderMove() OrderMove {
	return o.Data.(OrderMove)
}

func (o *Order) GetOrderKick() OrderKick {
	return o.Data.(OrderKick)
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
		var mov OrderMove
		err = json.Unmarshal(tmp.Data, &mov)
		o.Data = mov
	case tmp.Type == KICK:
		var mov OrderKick
		err = json.Unmarshal(tmp.Data, &mov)
		o.Data = mov
	case tmp.Type == CATCH:
		o.Data = nil
	default:
		err = errors.New(fmt.Sprintf("Unknow order type %s", tmp.Type))
	}
	return err
}
