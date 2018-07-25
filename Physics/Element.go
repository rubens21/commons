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

func (e *Element) HasCollided(obstacle *Element) (bool, float64)  {
	minDistance := float64(e.Size + obstacle.Size) / 2
	centerDistance := NewVector(e.Coords, obstacle.Coords).Length()
	realDistance := centerDistance - minDistance
	return realDistance < 0, realDistance
}

func (e *Element) VectorCollides(vector Vector, from Point, margin float64) *Point {
	if collide, point1, point2 := e.LineCollides(from, vector.TargetFrom(from), margin); collide {
		if point2 != nil {
			var nearestPoint *Point
			nearestPoint = nil

			vectorLength := vector.Length()
			distance := vectorLength + 1//just initializing
			if vector.IsObstacle(from, *point1) {
				distance = from.DistanceTo(*point1)
				nearestPoint = point1
			}
			if vector.IsObstacle(from, *point2) && from.DistanceTo(*point2) < distance {
				distance = from.DistanceTo(*point2)
				nearestPoint = point2
			}
			// when the distance is too small or even zero, it means the obstale the right behind the vector
			if distance < 0.01 || vectorLength - distance < 0.01 {
				return nil
			}
			return nearestPoint
		}
	}
	return nil
}

func (e *Element) LineCollides(a, b Point, margin float64) (bool, *Point, *Point)  {
	// Credits: https://stackoverflow.com/a/1088058/2047138
	c := e.Coords
	// compute the euclidean distance between A and B
	LAB := math.Sqrt( math.Pow(float64(b.PosX-a.PosX), 2) + math.Pow(float64(b.PosY-a.PosY), 2) )

	// compute the direction vector D from A to B
	Dx := float64(b.PosX-a.PosX)/LAB
	Dy := float64(b.PosY-a.PosY)/LAB

	// Now the line equation is x = Dx*t + Ax, y = Dy*t + Ay with 0 <= t <= 1.

	// compute the value t of the closest point to the circle center (Cx, Cy)
	t := Dx*float64(c.PosX-a.PosX) + Dy*float64(c.PosY-a.PosY)

	// This is the projection of C on the line from A to B.

	// compute the coordinates of the point E on line and closest to C
	Ex := t*Dx+float64(a.PosX)
	Ey := t*Dy+float64(a.PosY)

	// compute the euclidean distance from E to C
	LEC := math.Sqrt( math.Pow(Ex-float64(c.PosX), 2)+ math.Pow(Ey-float64(c.PosY), 2) )

	R := (float64(e.Size) / 2) + margin
	// test if the line intersects the circle
	if LEC < R	{
		// compute distance from t to circle intersection point
		dt := math.Sqrt( math.Pow(R, 2) - math.Pow(LEC, 2))

		// compute first intersection point
		Fx := (t-dt)*Dx + float64(a.PosX)
		Fy := (t-dt)*Dy + float64(a.PosY)

		// compute second intersection point
		Gx := (t+dt)*Dx + float64(a.PosX)
		Gy := (t+dt)*Dy + float64(a.PosY)

		return true, &Point{
			int(math.Round(Fx)), int(math.Round(Fy)),
		}, &Point{
			int(math.Round(Gx)), int(math.Round(Gy)),
		}
	} else if LEC == R { // else test if the line is tangent to circle
		// tangent point to circle is E
		return true, &Point{
			int(math.Round(Ex)), int(math.Round(Ey)),
		}, nil

	} else {// line doesn't touch circle
		return false, nil, nil
	}
}