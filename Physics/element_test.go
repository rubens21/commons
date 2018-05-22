package Physics

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestElement_LineCollides(t *testing.T) {
	table := map[string]struct {
		a        Point
		b        Point
		elPos    Point
		elRadius int
		maring   float64

		expectedCollide bool
		expectedCol1    *Point
		expectedCol2    *Point
	}{
		"Element axys 0,0 lin in X ": {
			Point{-10, 0},
			Point{10, 0},
			Point{0, 0},
			4,
			0,
			true,
			&Point{-4, 0},
			&Point{4, 0},
		},
		"Element axys 10,0 lin in X ": {
			Point{-10, 0},
			Point{10, 0},
			Point{10, 0},
			4,
			0,
			true,
			&Point{6, 0},
			&Point{14, 0},
		},
		"Element axys -10,0 lin in X ": {
			Point{-10, 0},
			Point{10, 0},
			Point{-10, 0},
			4,
			0,
			true,
			&Point{-14, 0},
			&Point{-6, 0},
		},
		"Element axys 0,5 lin in X ": {
			Point{-10, 0},
			Point{10, 0},
			Point{0, 5},
			4,
			0,
			false,
			nil,
			nil,
		},
		"Element axys 0,5 lin in 0,5>10,5 ": {
			Point{0, 5},
			Point{10, 5},
			Point{0, 5},
			4,
			0,
			true,
			&Point{-4, 5},
			&Point{4, 5},
		},
		"Element axys 0,9 lin in 0,5>10,5 ": {
			Point{0, 5},
			Point{10, 5},
			Point{0, 9},
			4,
			0,
			true,
			&Point{0, 5},
			nil,
		},
		"Element axys 0,9 lin in 0,2>10,2 ": {
			Point{0, 2},
			Point{10, 2},
			Point{0, 9},
			4,
			0,
			false,
			nil,
			nil,
		},
		"Diagonal": {
			Point{-30, -40},
			Point{30, 40},
			Point{0, 0},
			2,
			6.0,
			true,
			&Point{-3, -4},
			&Point{3, 4},
		},
	}

	for title, set := range table {
		element := Element{}
		element.Size = set.elRadius * 2
		element.Coords = set.elPos

		collid, col1, col2 := element.LineCollides(set.a, set.b, set.maring)
		if set.expectedCollide {
			assert.True(t, collid, title)
		} else {
			assert.False(t, collid, title)
		}

		assert.Equal(t, set.expectedCol1, col1, title)
		assert.Equal(t, set.expectedCol2, col2, title)
	}
}

func TestElement_VectorCollides(t *testing.T) {
	element := Element{}
	element.Size = 10
	element.Coords = Point{0, 0}
	var collisionPoint *Point

	vecA := NewVector(Point{}, Point{12,0})

	assert.Nil(t, element.VectorCollides(*vecA, Point{-20, 0}, 0.0))
	assert.Nil(t, element.VectorCollides(*vecA, Point{13, 0}, 0.0))
	assert.Nil(t, element.VectorCollides(*vecA, Point{0, 11}, 0.0))
	assert.Nil(t, element.VectorCollides(*vecA, Point{0, -5}, 0.0))

	assert.Nil(t, element.VectorCollides(*vecA, Point{-17, 0}, 0.0))



	collisionPoint = element.VectorCollides(*vecA, Point{-11, 0}, 0.0)
	assert.Equal(t, &Point{-5, 0}, collisionPoint)

	collisionPoint = element.VectorCollides(*vecA, Point{5, 0}, 0.0)
	assert.Equal(t, &Point{5, 0}, collisionPoint)


	element.Coords = Point{20, 20}

	vecB := NewVector(Point{}, Point{10,10})


	vecB.SetLength(50)
	collisionPoint = element.VectorCollides(*vecB, Point{0, 0}, 0.0)
	assert.Equal(t, &Point{18, 18}, collisionPoint)
}
