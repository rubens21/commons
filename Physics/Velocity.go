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
	copied := velocity.Copy()

	copied.Direction.SetLength(copied.Speed)
	v.Direction.SetLength(v.Speed)
	//if the vector is the inverse of the actual, we cannot sum them because they would null each other
	if copied.Copy().Direction.Invert().IsEqualTo(v.Direction) {
		v.Direction.Invert()
		v.Speed = 0
	} else {
		v.Direction.Add(copied.Direction)
		v.Speed = v.Direction.Length()
	}

	v.Direction.Normalize()
}

func (v *Velocity) String() string {
	return fmt.Sprintf("[%.2fx,%.2fy => %.2fs]", v.Direction.GetX(), v.Direction.GetY(), v.Speed)
}
