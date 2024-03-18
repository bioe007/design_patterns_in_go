package main

import "testing"

func TestBuildCheese(t *testing.T) {
	b := BurgerBuilder().WithBun("seed").
		WithCheese("american").
		WithMayo(true).
		WithOnions(true).
		WithKetchup(true)

	if b.cheese != "american" {
		t.Error("We WANT AMERICaN CHEEZE")
	}
}
