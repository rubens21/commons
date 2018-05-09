package tests

import (
	"testing"
	"github.com/maketplay/commons/Physics"
	"math"
)

func TestColisionDectection(t *testing.T) {
	elementA := Physics.Element{
		Size:     9,
		Coords:   Physics.Point{},
		Velocity: Physics.NewZeroedVelocity(),
	}
	elementB := Physics.Element{
		Size:     12,
		Coords:   Physics.Point{PosX: 50},
		Velocity: Physics.NewZeroedVelocity()}

	elementsBodySpace := float64(elementB.Size+elementA.Size) / 2

	input := []struct {
		Name   string
		A      Physics.Point
		B      Physics.Point
		Collid bool
		Dist   float64
	}{
		// X
		{"X1", Physics.Point{}, Physics.Point{50, 0}, false, float64(50) - elementsBodySpace},
		{"X2", Physics.Point{}, Physics.Point{25, 0}, false, float64(25) - elementsBodySpace},
		{"X3", Physics.Point{}, Physics.Point{11, 0}, false, float64(11) - elementsBodySpace},
		{"X4", Physics.Point{}, Physics.Point{10, 0}, true, float64(10) - elementsBodySpace},
		{"X5", Physics.Point{}, Physics.Point{0, 0}, true, float64(0) - elementsBodySpace},
		{"X6", Physics.Point{}, Physics.Point{-10, 0}, true, float64(10) - elementsBodySpace},
		{"X7", Physics.Point{}, Physics.Point{-11, 0}, false, float64(11) - elementsBodySpace},
		{"X8", Physics.Point{}, Physics.Point{-25, 0}, false, float64(25) - elementsBodySpace},

		// Y
		{"Y1", Physics.Point{}, Physics.Point{0, 50}, false, float64(50) - elementsBodySpace},
		{"Y2", Physics.Point{}, Physics.Point{0, 25}, false, float64(25) - elementsBodySpace},
		{"Y3", Physics.Point{}, Physics.Point{0, 11}, false, float64(11) - elementsBodySpace},
		{"Y4", Physics.Point{}, Physics.Point{0, 10}, true, float64(10) - elementsBodySpace},
		{"Y5", Physics.Point{}, Physics.Point{0, 0}, true, float64(0) - elementsBodySpace},
		{"Y6", Physics.Point{}, Physics.Point{0, -10}, true, float64(10) - elementsBodySpace},
		{"Y7", Physics.Point{}, Physics.Point{0, -11}, false, float64(11) - elementsBodySpace},
		{"Y8", Physics.Point{}, Physics.Point{0, -25}, false, float64(25) - elementsBodySpace},

		// diagonal

		{"D1", Physics.Point{}, Physics.Point{2, 5}, true, float64(5.38516480713) - elementsBodySpace},
		{"D2", Physics.Point{}, Physics.Point{5, 9}, true, float64(10.295630141) - elementsBodySpace},
		{"D3", Physics.Point{}, Physics.Point{1, 1}, true, float64(1.41421356237) - elementsBodySpace},
		{"D4", Physics.Point{}, Physics.Point{0, 0}, true, float64(0) - elementsBodySpace},
		{"D5", Physics.Point{}, Physics.Point{5, 10}, false, float64(11.1803398875) - elementsBodySpace},
		{"D6", Physics.Point{}, Physics.Point{10, 5}, false, float64(11.1803398875) - elementsBodySpace},
		{"D7", Physics.Point{}, Physics.Point{20, 20}, false, float64(28.2842712475) - elementsBodySpace},

		//diagnoal negative
		{"D1", Physics.Point{}, Physics.Point{-2, 5}, true, float64(5.38516480713) - elementsBodySpace},
		{"D2", Physics.Point{}, Physics.Point{-5, 9}, true, float64(10.295630141) - elementsBodySpace},
		{"D3", Physics.Point{}, Physics.Point{1, -1}, true, float64(1.41421356237) - elementsBodySpace},
		{"D4", Physics.Point{}, Physics.Point{0, 0}, true, float64(0) - elementsBodySpace},
		{"D5", Physics.Point{}, Physics.Point{5, -10}, false, float64(11.1803398875) - elementsBodySpace},
		{"D6", Physics.Point{}, Physics.Point{-10, 5}, false, float64(11.1803398875) - elementsBodySpace},
		{"D7", Physics.Point{}, Physics.Point{20, -20}, false, float64(28.2842712475) - elementsBodySpace},
		{"D7", Physics.Point{}, Physics.Point{-20, -20}, false, float64(28.2842712475) - elementsBodySpace},
	}

	for _, testCase := range input {
		elementA.Coords = testCase.A
		elementB.Coords = testCase.B
		collide, dist := elementA.HasCollide(&elementB)
		if collide != testCase.Collid {
			if testCase.Collid {
				t.Errorf("%s: A and B should had collide. Error dist: %f", testCase.Name, dist)
			} else {
				t.Errorf("%s: A and B should not being colliding. Error dist: %f", testCase.Name, dist)
			}
		}

		if math.Abs(math.Abs(dist) - math.Abs(testCase.Dist)) > 0.00001 {
			t.Errorf("%s: Wrong distance dist: %f (expected %f)", testCase.Name, dist, testCase.Dist)
		}

	}

	// invert
	for _, testCase := range input {
		elementA.Coords = testCase.B
		elementB.Coords = testCase.A
		collide, dist := elementA.HasCollide(&elementB)
		if collide != testCase.Collid {
			if testCase.Collid {
				t.Errorf("%s: A and B should had collide. Error dist: %f", testCase.Name, dist)
			} else {
				t.Errorf("%s: A and B should not being colliding. Error dist: %f", testCase.Name, dist)
			}
		}

		if math.Abs(math.Abs(dist) - math.Abs(testCase.Dist)) > 0.00001 {
			t.Errorf("%s: Wrong distance dist: %f (expected %f)", testCase.Name, dist, testCase.Dist)
		}

	}
}
