package Physics

import "fmt"

type Velocity struct {
	Direction *Vector  `json:"direction"`
	Speed     float64 `json:"speed"`
}

func NewZeroedVelocity() Velocity {
	s := Velocity{}
	s.Direction = new(Vector)
	s.Speed = 0
	return s
}

func (v *Velocity) Copy() Velocity {
	copyS := NewZeroedVelocity()
	copyS.Speed = v.Speed
	copyS.Direction = v.Direction.Copy()
	return copyS
}

func (v *Velocity) Target(from Point) Point {
	if v.Speed == 0 {
		return from
	} else {
		v.Direction.SetLength(v.Speed)
		return v.Direction.TargetFrom(from)
	}
}

func (v *Velocity) Add(velocity Velocity) {
	v.Speed += velocity.Speed
	v.Direction.Add(velocity.Direction)
}

func (v *Velocity) String() string {
	return fmt.Sprintf("[%.2fx,%.2fy => %.2fs]", v.Direction.GetX(), v.Direction.GetY(), v.Speed)
}
