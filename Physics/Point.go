package Physics

import (
	"math"
)

type Point struct {
	PosX int `json:"x"`
	PosY int `json:"y"`
}

func (p *Point) DistanceTo(coordTo Point) (distance float64) {
	catA := float64(coordTo.PosX - p.PosX)
	catO := float64(coordTo.PosY - p.PosY)
	return math.Hypot(catA, catO)
}

