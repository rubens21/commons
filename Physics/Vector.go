package Physics

import (
	"math"
	"encoding/json"
)

type Vector struct {
	x float64
	y float64
}

func NewVector(from Point, to Point) *Vector {
	v := new(Vector)
	v.x = float64(to.PosX) - float64(from.PosX)
	v.y = float64(to.PosY) - float64(from.PosY)
	if v.x == 0 && v.y == 0 {
		panic("vector can not have zero length")
	}
	return v
}

func (v Vector) Copy() *Vector {
	nv := new(Vector)
	nv.x = v.x
	nv.y = v.y
	return nv
}

func (v *Vector) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"x":   v.x,
		"y":   v.y,
		"ang": v.AngleDegrees(),
	})
}

func (v *Vector) UnmarshalJSON(b []byte) error {
	var tmp struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}
	v.x = tmp.X
	v.y = tmp.Y
	return nil
}

func (v *Vector) Normalize() *Vector {
	length := v.Length()
	if length > 0 {
		v.Scale(100 / length)
	}
	return v
}

func (v *Vector) SetLength(length float64) *Vector {
	if length == 0 {
		panic("vector can not have zero length")
	}
	v.Scale(length / v.Length())
	return v
}

func (v *Vector) SetX(x float64) *Vector {
	v.x = x
	return v
}

func (v *Vector) SetY(y float64) *Vector {
	v.y = y
	return v
}

func (v *Vector) Invert() *Vector {
	v.x = -v.x
	v.y = -v.y
	return v
}

func (v *Vector) Scale(t float64) *Vector {
	v.x *= t
	v.y *= t
	return v
}

func (v *Vector) Sin() float64 {
	return v.y / v.Length()
}

func (v *Vector) Cos() float64 {
	return v.x / v.Length()
}

func (v *Vector) Angle() float64 {
	return math.Acos(v.Cos())
}

func (v *Vector) AngleDegrees() float64 {
	angX := math.Acos(v.Cos()) * 180 / math.Pi

	if v.y < 0 {
		angX *= -1
	}
	return angX

}

func (v *Vector) OppositeAngle() float64 {
	return math.Acos(v.Cos())
}

func (v *Vector) Length() float64 {
	return math.Hypot(v.x, v.y)
}

func (v *Vector) Add(vector *Vector) *Vector {
	v.x += vector.x
	v.y += vector.y
	return v
}

func (v *Vector) Sub(vector *Vector) *Vector {
	v.x -= vector.x
	v.y -= vector.y
	return v
}

func (v *Vector) TargetFrom(point Point) Point {
	return Point{
		point.PosX + int(math.Round(v.x)),
		point.PosY + int(math.Round(v.y)),
	}
}

func (v *Vector) GetX() float64 {
	return v.x
}
func (v *Vector) GetY() float64 {
	return v.y
}

func (v *Vector) IsEqualTo(b *Vector) bool {
	copyMe := v.Copy().Normalize()
	copyOther := b.Copy().Normalize()
	return copyMe.y == copyOther.y && copyMe.x == copyOther.y
}

func (v *Vector) AngleWith(b *Vector) float64 {
	//http://onlinemschool.com/math/assistance/vector/angl/
	copyMe := v.Copy().Normalize()
	copyOther := b.Copy().Normalize()

	dotProduct := (copyMe.x * copyOther.x) + (copyMe.y * copyOther.y)
	cos := dotProduct / (copyMe.Length() * copyOther.Length())
	ang := math.Round(math.Acos(cos)*(180/math.Pi)*100) / 100
	if copyMe.y > copyOther.y {
		ang *= -1
	}
	return ang
}

func (v *Vector) IsObstacle(from Point, obstacle Point) bool {
	to := v.TargetFrom(from)
	a := from.DistanceTo(obstacle)
	b := obstacle.DistanceTo(to)
	hypo := from.DistanceTo(to)
	return math.Round(a+b-hypo) < 0.1
}
