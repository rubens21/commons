package GameState

import "github.com/makeitplay/commons/BasicTypes"

type State BasicTypes.State

const (
	WAITINGTEAMS State = "waiting"
	GETREADY     State = "get-ready"
	LISTENING    State = "listening"
	PAUSE        State = "pause" // debbuging mode, waiting for "nextStep" signal
	PLAYING      State = "playing"
	RESULTS      State = "results"
	OVER         State = "game-over"
)
