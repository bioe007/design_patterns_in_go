package main

type Burger struct {
	bun, cheese                             string
	pickles, onions, ketchup, mustard, mayo bool
}

func BurgerBuilder() *Burger {
	b := &Burger{}
	return b
}

func (b *Burger) WithBun(s string) *Burger {
	b.bun = s
	return b
}

func (b *Burger) WithPickles(v bool) *Burger {
	b.pickles = v
	return b
}

func (b *Burger) WithOnions(v bool) *Burger {
	b.onions = v
	return b
}

func (b *Burger) WithKetchup(v bool) *Burger {
	b.ketchup = v
	return b
}

func (b *Burger) WithMustard(v bool) *Burger {
	b.mustard = v
	return b
}

func (b *Burger) WithMayo(v bool) *Burger {
	b.mayo = v
	return b
}

func (b *Burger) WithCheese(s string) *Burger {
	b.cheese = s
	return b
}
