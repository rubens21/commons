package Units

import "github.com/maketplay/commons/Physics"

type TeamPlace string

const HomeTeam TeamPlace = "home"
const AwayTeam TeamPlace = "away"
const PlayerSize = 10
const BallSize = 5
const BallMinPower = 0.25 // minimal power to make the ball move

const CourtWidth = 200  // unitss
const CourtHeight = 100 // units


const BallSlowerRatio = 0.70                                 // ratio value for slowing the ball
const PlayerSpeed = 5.0                                      // unit/lance
const BallSpeed = 20.0                                       // units/lance
const DistanceCatchBall = float64(PlayerSize+BallSize) * 0.6 // units float

var HomeTeamGoalcenter = Physics.Point{0, CourtHeight / 2}
var AwayTeamGoalcenter = Physics.Point{CourtWidth, CourtHeight / 2}
const GoalWidth = PlayerSpeed * 4

type PlayerNumber string

const (
	PositionA PlayerNumber = "1"
	PositionB PlayerNumber = "2"
	PositionC PlayerNumber = "3"
	PositionD PlayerNumber = "4"
	PositionE PlayerNumber = "5"
)

var InitialPostionHomeTeam = map[PlayerNumber]Physics.Point{
	PositionA: {45,75},
	PositionB: {45,50},
	PositionC: {45,25},
	PositionD: {75,60},
	PositionE: {75,45},
}

var InitialPostionAwayTeam = map[PlayerNumber]Physics.Point{
	PositionA: {155, 25},
	PositionB: {155, 50},
	PositionC: {155, 75},
	PositionD: {125, 45},
	PositionE: {125, 60},
}

// Invert the coords X and Y as in a mirror to found out the same position seen from the away team field
// Keep in mind that all coords in the field is based on the bottom left corner!
func MirrorCoordToAway(coords Physics.Point) Physics.Point {
	return Physics.Point{
		PosX: CourtWidth - coords.PosX,
		PosY: CourtHeight - coords.PosY,
	}
}