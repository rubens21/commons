package GameState

import "github.com/maketplay/commons/BasicTypes"

type State BasicTypes.State

const (
	WAITINGTEAMS State = "waiting"
	GETREADY     State = "get-ready"
	LISTENING    State = "listening"
	PAUSE        State = "pause"
	PLAYING      State = "playing"
	RESULTS      State = "results"
	OVER         State = "game-over"
)
