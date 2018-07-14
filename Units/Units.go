package Units

import (
	"github.com/makeitplay/commons/Physics"
)

type TeamPlace string

const HomeTeam TeamPlace = "home"
const AwayTeam TeamPlace = "away"


const BaseUnit = 100

const PlayerSize = 4 * BaseUnit
const PlayerMaxSpeed = 8.0 * BaseUnit // unit/lance

const CourtWidth = 200 * BaseUnit // unitss
var CourtCenter = Physics.Point{CourtWidth / 2, CourtHeight / 2}
const CourtHeight = 100 * BaseUnit // units

const BallSize = 2 * BaseUnit
const BallDeceleration = 3 * BaseUnit // ratio value for slowing the ball
const BallMaxSpeed = 20.0 * BaseUnit  // units/lance
const BallMinSpeed = 0.05 * BaseUnit // minimal power to make the ball move
const BallTimeInGoalZone = 6 //turns

const GoalWidth = 30 * BaseUnit
const GoalMinY = (CourtHeight - GoalWidth) / 2
const GoalMaxY = GoalMinY + GoalWidth
const GoalZoneRange = 14 * BaseUnit

const GoalKeeperJumpLength = 2 * BaseUnit
