package generator_test

import (
	"testing"

	"github.com/bioe007/design_patterns_in_go/generator"
)

func TestGenerate(t *testing.T) {
	x := []int{1, 2, 3, 5, 6, 7, 8, 9}

	// g := generator.Generate(x)

	count := 0
	for v := range generator.Generate(x) {
		if v != x[count] {
			t.Fatalf("Received %d, expected %d", v, x[count])
		}
		count++
		// fmt.Println("generating", v)
	}
}
