package vector_test

import (
	"fmt"
	"testing"

	"github.com/kvartborg/vector"
)

type vec = vector.Vector

func TestMultiDimensionalVec(t *testing.T) {
	result := vec{1}.Add(vec{1, 2})

	if result[0] != 2 || result[1] != 2 {
		t.Error("dimension mishmatch")
	}
}

func Example() {
	// Create a new vector
	v1 := vec{4, 2}

	// Create a vector from a list of float64
	v2 := vec([]float64{1, 2, 4})

	fmt.Println(v1.Add(v2))
	// Output: [5 4 4]
}

func ExampleAdd() {
	result := vector.Add(vec{0, 2}, vec{1, 4})
	fmt.Println(result)
	// Output: [1 6]
}

func ExampleVector_Add() {
	fmt.Println(
		vec{0, 2}.Add(vec{1, 4}),
	)
	// Output: [1 6]
}

func ExampleSub() {
	result := vector.Sub(vec{1, 4}, vec{0, 2})
	fmt.Println(result)
	// Output: [1 2]
}

func ExampleVector_Sub() {
	fmt.Println(
		vec{1, 4}.Sub(vec{0, 2}),
	)
	// Output: [1 2]
}

func ExampleDot() {
	result := vector.Dot(vec{0, 2}, vec{2, 0})
	fmt.Println(result)
	// Output: 0
}

func ExampleVector_Dot() {
	v1, v2 := vec{0, 2}, vec{2, 0}
	fmt.Println(v1.Dot(v2))
	// Output: 0
}

func ExampleCross() {
	result := vector.Cross(vec{0, 1, 2}, vec{3, 2, 1})
	fmt.Println(result)
	// Output: [-3 6 -6]
}

func ExampleVector_Cross() {
	v1, v2 := vec{0, 1, 2}, vec{3, 2, 1}
	fmt.Println(v1.Cross(v2))
	// Output: [-3 6 -6]
}

func BenchmarkAdd(b *testing.B) {
	v1, v2 := vec{1, 2}, vec{2, 3}

	for i := 0; i < b.N; i++ {
		vector.Add(v1, v2)
	}
}