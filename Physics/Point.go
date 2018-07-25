package Physics

import (
	"math"
	"fmt"
)

type Point struct {
	PosX int `json:"x"`
	PosY int `json:"y"`
}

func (p *Point) DistanceTo(coordTo Point) (distance float64) {
	catA := float64(coordTo.PosX) - float64(p.PosX)
	catO := float64(coordTo.PosY) - float64(p.PosY)
	return math.Hypot(catA, catO)
}

func (p *Point) MiddlePointTo(target Point) Point  {
	x := math.Abs(float64(p.PosX - target.PosX))
	y := math.Abs(float64(p.PosY - target.PosY))

	return Point{
		PosX: int(math.Round(math.Min(float64(p.PosX), float64(target.PosX)) + x)),
		PosY: int(math.Round(math.Min(float64(p.PosY), float64(target.PosY)) + y)),
	}
}

func (p *Point) String() string {
	return fmt.Sprintf("{%d, %d}", p.PosX, p.PosY)
}

