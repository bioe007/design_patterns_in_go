package main

import (
	"testing"

	"github.com/bioe007/design_patterns_in_go/singleton"
)

func TestMainSingletonAdd(t *testing.T) {
	u := singleton.GetSingleton()
	u.Add(42)
	if u.Get() != 42 {
		t.Error("not the answser")
	}
}

func TestMainSingletonMultiAdd(t *testing.T) {
	sng0 := singleton.GetSingleton()
	sng1 := singleton.GetSingleton()

	sng1.Add(1)
	sng1.Add(1)

	if sng0.Get() != sng1.Get() {
		t.Errorf("Expected 2 == 2, got %d != %d", sng1.Get(), sng0.Get())
	}
}
