package BasicTypes

import (
	"github.com/makeitplay/commons/Units"
	"github.com/makeitplay/commons/Physics"
)

type MsgType string

type OrderType string

type PlayerNumber string

const (
	// msg type
	ORDER        MsgType = "ORDER"        // just started
	ANNOUNCEMENT MsgType = "ANNOUNCEMENT" // let you know the state has changed
	DEBUG        MsgType = "DEBUG"        // a debug command has happened
	SCORE        MsgType = "SCORE"        // score has changed
	RIP          MsgType = "RIP"          // main process has died
	WELCOME      MsgType = "WELCOME"      // main process has died
)

type State string

type PlayerSpecifications struct {
	Number        PlayerNumber
	InitialCoords Physics.Point
}

type Goal struct {
	Center     Physics.Point
	Place      Units.TeamPlace
	TopPole    Physics.Point
	BottomPole Physics.Point
}
