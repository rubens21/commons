package Units

// TeamPlace defines a side of the team during the game (left for home team, and right for the away team)
type TeamPlace string

// HomeTeam identify the home team
const HomeTeam TeamPlace = "home"

// AwayTeam identify the home team
const AwayTeam TeamPlace = "away"

// BaseUnit is used to increase the integer units scale and improve the precision when the integer numbers
// come from float  calculations. Some units have to be integer to avoid infinite intervals (e.g. a point in the field, element sizes)
const BaseUnit = 100

// PlayerSize is the size of each player
const PlayerSize = 4 * BaseUnit

// PlayerMaxSpeed is the max speed that a play may move
const PlayerMaxSpeed = 2.5 * BaseUnit

// CourtWidth is the width of the court (horizontal view)
const CourtWidth = 200 * BaseUnit

// CourtHeight is the height of the court (horizontal view)
const CourtHeight = 100 * BaseUnit

const BallSize = 2 * BaseUnit
const BallDeceleration = 0.4 * BaseUnit // ratio value for slowing the ball
const BallMaxSpeed = 8.8 * BaseUnit     // units/lance
const BallMinSpeed = 0.05 * BaseUnit    // minimal power to make the ball move
const BallTimeInGoalZone = 3 * 10       //about 3 seconds

const GoalWidth = 30 * BaseUnit
const GoalMinY = (CourtHeight - GoalWidth) / 2
const GoalMaxY = GoalMinY + GoalWidth
const GoalZoneRange = 14 * BaseUnit

const GoalKeeperJumpLength = 4 * BaseUnit
