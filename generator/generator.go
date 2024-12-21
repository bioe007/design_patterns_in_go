package generator

func Generate[T any](arr []T) chan T {
	gen := make(chan T)

	go func() {
		defer close(gen)
		for _, value := range arr {
			gen <- value
		}
	}()

	return gen
}
