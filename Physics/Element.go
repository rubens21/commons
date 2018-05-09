package Physics

import (
	"math"
)

type Element struct {
	Size       int
	Coords     Point     `json:"position"`
	Velocity   Velocity `json:"velocity"`
}

func (e *Element) MoveTo(coords Point) {
	e.Coords = coords
}

func (e *Element) GetCoords() Point {
	return e.Coords
}

func (e *Element) IsObstacle(target Point, obstacle Point, errMarginDegree float64) (degree float64, inRange bool) {
	if e.Coords.DistanceTo(obstacle) > e.Coords.DistanceTo(target) {
		return 0, false
	} else {
		/*
		The angule is calculated based on the diff between the coseno angule of the target and the obstacle.
		However, this values will always be positive if both are in the same directions (less than 90 degrees
		of difference).
		Then, target figure out if the obstacle is on the left or on the right of the route, we have
		target check if the sin is positive or negative because the sin is base on the Y axis
		 */
		vectorTarget := NewVector(e.Coords, target)
		vectorObstacle := NewVector(e.Coords, obstacle)

		angTarg := math.Acos(vectorTarget.Cos()) * (180 / math.Pi)
		angObst := math.Acos(vectorObstacle.Cos()) * (180 / math.Pi)

		angDirctionObst := math.Asin(vectorObstacle.Sin()) * (180 / math.Pi)

		diff := angObst - angTarg
		if angDirctionObst < 0 {
			diff *= -1
		}
		return diff, math.Abs(diff) <= errMarginDegree
	}
}

func (e *Element) HasCollide(obstacle *Element) (bool, float64)  {
	minDistance := float64(e.Size + obstacle.Size) / 2
	centerDistance := NewVector(e.Coords, obstacle.Coords).Length()
	realDistance := centerDistance - minDistance
	return realDistance < 0, realDistance
}