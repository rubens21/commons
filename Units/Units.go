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

// BallSize size of the element ball
const BallSize = 2 * BaseUnit

// BallDeceleration is the deceleration rate of the ball speed.
const BallDeceleration = 0.4 * BaseUnit

// BallMaxSpeed is the max speed of the ball
const BallMaxSpeed = 8.8 * BaseUnit

// BallMinSpeed is the minimal speed of the ball. When the ball was at this speed or slower, it will be considered stopped.
const BallMinSpeed = 0.05 * BaseUnit

// BallTimeInGoalZone is the max number of turns that the ball may be in a goal zone. After that, the ball will be auto kicked
// towards the center of the field.
const BallTimeInGoalZone = 15

// GoalWidth is the goal width
const GoalWidth = 30 * BaseUnit

// GoalMinY is the coordinate Y of the lower pole of the goals
const GoalMinY = (CourtHeight - GoalWidth) / 2

// GoalMaxY is the coordinate Y of the upper pole of the goals
const GoalMaxY = GoalMinY + GoalWidth

// GoalZoneRange is the minimal distance that a player can stay from the opponent goal
const GoalZoneRange = 14 * BaseUnit

// GoalKeeperJumpLength is the distance that the goalkeeper jumps when he executes the order CATCH towards the ball.
const GoalKeeperJumpLength = 4 * BaseUnit
