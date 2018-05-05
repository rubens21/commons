package Physics

type Velocity struct {
	Direction Vector  `json:"direction"`
	Speed     float64 `json:"speed"`
}

func NewZeroedVelocity() Velocity {
	s := Velocity{}
	s.Direction = Vector{0,0}
	s.Speed = 0
	return s
}

func (v *Velocity) Copy() Velocity {
	copyS := NewZeroedVelocity()
	copyS.Speed = v.Speed
	copyS.Direction = *v.Direction.Copy()
	return copyS
}
func (v *Velocity) Add(velocity Velocity) {
	v.Speed += velocity.Speed
	v.Direction.Add(&velocity.Direction)
}
