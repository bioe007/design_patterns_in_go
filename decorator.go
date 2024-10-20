package main

import (
	"log"
	"time"

	decoraterobots "github.com/bioe007/design_patterns_in_go/decorateRobots"
)

// Decorators allow wrapping another interface in some new behavior

// Pretend we have a robot that moves about which is made by another company.
// We have to go through their API to move the robot.
// But sometimes we want to log how long it takes between stops. We can decorate
// it's move interface to collect these times

// Implement the third party interface
type FastRobotMover struct {
	robot decoraterobots.Robot
}

func (r *FastRobotMover) Move(x, y int) {
	// fmt.Println("decorated move add +1 to args")
	r.robot.Move(x+10, y+10)
}

type SlowRobotMover struct {
	robot decoraterobots.Robot
}

func (r *SlowRobotMover) Move(x, y int) {
	r.robot.Move(x/2, y/2)
}

// An alternative is to decorate by wrapping the function itself.
// TODO - don't really test this..
type robotMoveFunc func(x, y int)

func wrap_move_time(fun robotMoveFunc, logger *log.Logger) robotMoveFunc {
	return func(x, y int) {
		defer func(t time.Time) {
			logger.Printf("move(%v,%v) took %v", x, y, time.Since(t))
		}(time.Now())
		fun(x, y)
	}
}
