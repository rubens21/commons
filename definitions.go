package commons

import (
	"github.com/makeitplay/commons/BasicTypes"
	"github.com/makeitplay/commons/Units"
	"github.com/makeitplay/commons/Physics"
)

const (
	GoalkeeperNumber BasicTypes.PlayerNumber = "1"
)

var HomeTeamGoal = BasicTypes.Goal{
	Place: Units.HomeTeam,
	Center: Physics.Point{0, Units.CourtHeight / 2},
	TopPole: Physics.Point{0, Units.GoalMaxY},
	BottomPole: Physics.Point{0, Units.GoalMinY},
}

var AwayTeamGoal = BasicTypes.Goal{
	Place: Units.HomeTeam,
	Center: Physics.Point{Units.CourtWidth, Units.CourtHeight / 2},
	TopPole: Physics.Point{Units.CourtWidth, Units.GoalMaxY},
	BottomPole: Physics.Point{Units.CourtWidth, Units.GoalMinY},
}