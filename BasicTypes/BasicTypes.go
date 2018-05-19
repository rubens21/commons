package BasicTypes

import (
	"github.com/makeitplay/commons/Units"
	"github.com/makeitplay/commons/Physics"
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
	InitialCoords Physics.Point
}
