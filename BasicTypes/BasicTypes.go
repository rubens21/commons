package BasicTypes

import (
	"github.com/maketplay/commons/Units"
	"github.com/maketplay/commons/Physics"
)

type MsgType string

type OrderType string

const (
	// msg type
	ORDER        MsgType = "ORDER"        // just started
	ANNOUNCEMENT MsgType = "ANNOUNCEMENT" // let you know
	SCORE        MsgType = "SCORE"        // score has changed
	RIP          MsgType = "RIP"          // main process has died
	WELCOME      MsgType = "WELCOME"          // main process has died
)

type State string


type PlayerSpecifications struct {
	Number Units.PlayerNumber
	InitialCoors Physics.Point
}
