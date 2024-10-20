package decoraterobots

type Robot struct {
	X, Y int
	Name string
}

// This is the interface that will get decorated
type RobotMover interface {
	Move(x, y int)
}

func (r *Robot) Move(x, y int) {
	// Pretend this step does some computation that takes time...
	// fmt.Println("regular move")
	r.X = x
	r.Y = y
}
