package main

import "testing"

func TestNewWithDefault(t *testing.T) {
	lo := New(42)
	if lo.id != 42 || lo.url != "www.foo.com" {
		t.Errorf("we're not in kansas anymore: %#v", lo)
	}
}

// Build the struct in a single call to New
func TestNewWithReddit(t *testing.T) {
	lo := New(42,
		WithUrl("reddit.com"),
		WithMaxage(22),
		WithBuz(96),
	)

	if lo.id != 42 || lo.url != "reddit.com" {
		t.Error("wrong")
	}
}

// or pass in my custom settings to build up Options
func TestNewWithIndividuals(t *testing.T) {
	u := WithUrl("individual.co")
	lo := New(42, u)

	if lo.url != "individual.co" || lo.id != 42 {
		t.Errorf("damnit: %#v", lo)
	}
}
