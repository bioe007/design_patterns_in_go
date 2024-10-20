package main

import (
	"bytes"
	"log"
	"testing"

	decoraterobots "github.com/bioe007/design_patterns_in_go/decorateRobots"
)

func TestUnDecoratedMove(t *testing.T) {
	r := decoraterobots.Robot{X: 0, Y: 0, Name: "suzy"}
	r.Move(5, 5)

	if r.X != 5 || r.Y != 5 {
		t.Errorf("expected (5,5) got (%d,%d)", r.X, r.Y)
	}
}

// Fast moving robots will move faster..
func TestFastMove(t *testing.T) {
	r_vanilla := decoraterobots.Robot{X: 0, Y: 0, Name: "eva"}
	r := FastRobotMover{robot: r_vanilla}
	r.Move(5, 5)

	if r.robot.X != 5+10 || r.robot.Y != 5+10 {
		t.Errorf("expected (5,5) got (%d,%d)", r.robot.X, r.robot.Y)
	}
}

// ...and slow ones slower.
func TestSlowMove(t *testing.T) {
	r_vanilla := decoraterobots.Robot{X: 0, Y: 0, Name: "eva"}
	r := SlowRobotMover{robot: r_vanilla}
	// moving now will be slower
	r.Move(5, 5)

	if r.robot.X != 5/2 || r.robot.Y != 5/2 {
		t.Errorf("expected (5,5) got (%d,%d)", r.robot.X, r.robot.Y)
	}
}

// But that depends on the 3rd party library, how do we know the decorator
// itself is working correctly?

// Create a mock interface
type robotMoveMock interface {
	Move(x, y int)
}

type robotMock struct {
	X, Y int
	Name string
}

// i can't remember where i was going with this
// func TestSlowMoveIso(t *testing.T) {
// 	r := robotMock{0, 0, "Mock"}
// 	r.Move(42, 42)
// }

func TestWrapLog(t *testing.T) {
	var buf bytes.Buffer
	r := decoraterobots.Robot{X: 42, Y: 42, Name: "marvin"}
	logger := log.New(&buf, "wrap: ", log.LstdFlags)
	logged_move := wrap_move_time(r.Move, logger)
	logged_move(42, 42)
	t.Log(buf.String())
}
