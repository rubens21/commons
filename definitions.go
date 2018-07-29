package commons

import (
	"github.com/makeitplay/commons/BasicTypes"
	"github.com/makeitplay/commons/Physics"
	"github.com/makeitplay/commons/Units"
)

const (
	// GoalkeeperNumber defines the goalkeeper number
	GoalkeeperNumber BasicTypes.PlayerNumber = "1"
)

// HomeTeamGoal works as a constant value to help to retrieve a Goal struct with the values of the Home team goal
var HomeTeamGoal = BasicTypes.Goal{
	Place:      Units.HomeTeam,
	Center:     Physics.Point{0, Units.CourtHeight / 2},
	TopPole:    Physics.Point{0, Units.GoalMaxY},
	BottomPole: Physics.Point{0, Units.GoalMinY},
}

// AwayTeamGoal works as a constant value to help to retrieve a Goal struct with the values of the Away team goal
var AwayTeamGoal = BasicTypes.Goal{
	Place:      Units.HomeTeam,
	Center:     Physics.Point{Units.CourtWidth, Units.CourtHeight / 2},
	TopPole:    Physics.Point{Units.CourtWidth, Units.GoalMaxY},
	BottomPole: Physics.Point{Units.CourtWidth, Units.GoalMinY},
}


// CourtCenter works as a constant value to help to retrieve a Point struct with the values of the center of the court
var CourtCenter = Physics.Point{Units.CourtWidth / 2, Units.CourtHeight / 2}
