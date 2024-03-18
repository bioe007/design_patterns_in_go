package singleton

import (
	"sync"
)

// This is only truly safe outside of it's package.
type theSingletonThing struct {
	m              sync.Mutex
	myPrivateIdaho bool
	value          int
}

type PublicThing struct {
	tst *theSingletonThing
}

var aSingleThing theSingletonThing

func init() {
	aSingleThing.myPrivateIdaho = false
}

func (a *theSingletonThing) Add(n int) {
	a.m.Lock()
	defer a.m.Unlock()
	a.value += n
}

func (a *theSingletonThing) Sub(n int) {
	a.m.Lock()
	defer a.m.Unlock()
	a.value -= n
}

func (a *theSingletonThing) Get() int {
	a.m.Lock()
	defer a.m.Unlock()
	return a.value
}

func GetSingleton() *theSingletonThing {
	aSingleThing.m.Lock()
	defer aSingleThing.m.Unlock()
	if aSingleThing.myPrivateIdaho == false {
		aSingleThing.myPrivateIdaho = true
	}
	return &aSingleThing
}
