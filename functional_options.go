package main

type LotsOfOptions struct {
	id                   int
	url, userauth        string
	maxconn, maxage, buz int // note
}

// One _could_ do something like this - but then client has to worry about
// every default value, even if they're not changing them. Let's say every time
// though we do have to set the 'id' value.
func PainfulNew(id int, url, userauth string, maxconn, maxage, buz int) *LotsOfOptions {
	lo := &LotsOfOptions{
		id,
		"www.foo.com",
		"deadbeef",
		10,
		10,
		10,
	}
	return lo
}

// The Option type allows us to concisely declare the `WithXXX` methods.
type Option func(*LotsOfOptions)

// Each `With<Option Name>` method will set the value of a struct member and
// return a callable.
func WithUrl(s string) Option {
	return func(lo *LotsOfOptions) {
		lo.url = s
	}
}

func WithUserauth(s string) Option {
	return func(lo *LotsOfOptions) {
		lo.userauth = s
	}
}

func WithMaxconn(v int) Option {
	return func(lo *LotsOfOptions) {
		lo.maxconn = v
	}
}

func WithMaxage(v int) Option {
	return func(lo *LotsOfOptions) {
		lo.maxage = v
	}
}

func WithBuz(v int) Option {
	return func(lo *LotsOfOptions) {
		lo.buz = v
	}
}

func New(id int, opts ...Option) *LotsOfOptions {
	lo := &LotsOfOptions{
		id,
		"www.foo.com",
		"deadbeef",
		10,
		10,
		10,
	}
	// every `opt` will be func that accepts the `LotsOfOptions` struct
	for _, opt := range opts {
		opt(lo)
	}
	return lo
}
