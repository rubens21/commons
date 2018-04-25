package BasicTypes

type Order struct {
	Type OrderType              `json:"order"`
	Data map[string]interface{} `json:"data"`
}

type MsgType string
type OrderType string

const (
	// msg type
	ORDER        MsgType = "ORDER"        // just started
	ANNOUNCEMENT MsgType = "ANNOUNCEMENT" // let you know
	SCORE        MsgType = "SCORE"        // score has changed
	RIP          MsgType = "RIP"          // main process has died

	// orders sent by the PLAYER
	ENTER OrderType = "ENTER" // I want to enjoy the team
	MOVE  OrderType = "MOVE"  // Move me to that position
	KICK  OrderType = "KICK"  // Cick the ball to that position
	CATCH OrderType = "CATCH" // I'll try to catch the ball of the player
)

type State string
