package Physics

import (
	"fmt"
	"math"
)

type Velocity struct {
	Direction *Vector  `json:"direction"`
	Speed     float64 `json:"speed"`
}

func NewZeroedVelocity(direction Vector) Velocity {
	s := Velocity{}
	s.Direction = &direction
	s.Speed = 0
	return s
}

func (v *Velocity) Copy() Velocity {
	copyS := NewZeroedVelocity(*v.Direction.Copy())
	copyS.Speed = v.Speed
	return copyS
}

func (v *Velocity) Target(from Point) Point {
	if v.Speed == 0 {
		return from
	} else {
		speedX := v.Speed * v.Direction.Cos()
		speedY := v.Speed * v.Direction.Sin()
		return Point{
			PosX: from.PosX + int(math.Round(speedX)),
			PosY: from.PosY + int(math.Round(speedY)),
		}
	}
}

func (v *Velocity) Add(velocity Velocity) {
	v.Speed += velocity.Speed
	v.Direction.Add(velocity.Direction)
}

func (v *Velocity) String() string {
	return fmt.Sprintf("[%.2fx,%.2fy => %.2fs]", v.Direction.GetX(), v.Direction.GetY(), v.Speed)
}
