package Units

import "github.com/maketplay/commons/Physics"

type TeamPlace string

const HomeTeam TeamPlace = "home"
const AwayTeam TeamPlace = "away"


const BaseUnit = 100
const PlayerSize = 4 * BaseUnit
const BallSize = 2 * BaseUnit
const BallMinVelocity = 0.25 // minimal power to make the ball move

const CourtWidth = 200 * BaseUnit // unitss
const CourtHeight = 100 * BaseUnit // units


const BallSlowerRatio = 0.70                                 // ratio value for slowing the ball
const PlayerMaxSpeed = 5.0 * BaseUnit                        // unit/lance
const BallMaxSpeed = 20.0 * BaseUnit                         // units/lance
const DistanceCatchBall = float64(PlayerSize+BallSize) * 0.6 // units float

var HomeTeamGoalCenter = Physics.Point{0, CourtHeight / 2}
var AwayTeamGoalCenter = Physics.Point{CourtWidth, CourtHeight / 2}
var CourtCenter = Physics.Point{CourtWidth / 2, CourtHeight / 2}
const GoalWidth = PlayerMaxSpeed * 4

type PlayerNumber string

const (
	PositionA PlayerNumber = "1"
	PositionB PlayerNumber = "2"
	PositionC PlayerNumber = "3"
	PositionD PlayerNumber = "4"
	PositionE PlayerNumber = "5"
)

var InitialPostionHomeTeam = map[PlayerNumber]Physics.Point{
	PositionA: {45 * BaseUnit,75 * BaseUnit},
	PositionB: {45 * BaseUnit,50 * BaseUnit},
	PositionC: {45 * BaseUnit,25 * BaseUnit},
	PositionD: {75 * BaseUnit,60 * BaseUnit},
	PositionE: {75 * BaseUnit,45 * BaseUnit},
}

var InitialPostionAwayTeam = map[PlayerNumber]Physics.Point{
	PositionA: {155 * BaseUnit, 25 * BaseUnit},
	PositionB: {155 * BaseUnit, 50 * BaseUnit},
	PositionC: {155 * BaseUnit, 75 * BaseUnit},
	PositionD: {125 * BaseUnit, 45 * BaseUnit},
	PositionE: {125 * BaseUnit, 60 * BaseUnit},
}

// Invert the coords X and Y as in a mirror to found out the same position seen from the away team field
// Keep in mind that all coords in the field is based on the bottom left corner!
func MirrorCoordToAway(coords Physics.Point) Physics.Point {
	return Physics.Point{
		PosX: CourtWidth - coords.PosX,
		PosY: CourtHeight - coords.PosY,
	}
}