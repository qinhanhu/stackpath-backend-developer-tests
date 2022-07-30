package fizzbuzz

import (
	"fmt"
	"testing"
)

type fizzBuzzTest struct {
	total, fizzAt, buzzAt int64
	expected []string
}

var fizzBuzzTests = []fizzBuzzTest{
	fizzBuzzTest{0, 0, 0, []string{}},
	fizzBuzzTest{0, 0, 0, []string{}},
	fizzBuzzTest{0, 0, 0, []string{}},
	fizzBuzzTest{0, 0, 0, []string{}},
	fizzBuzzTest{0, 0, 0, []string{}},
}

func stringSlicesEqual(a, b []string) bool {
	if len(a) != len(b){
		return false
	}
	for i, v := range a{
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestFizzBuzz(t *testing.T) {
	//t.Skip("TODO: Add tests")
	for _, test := range fizzBuzzTests{
		if result := FizzBuzz(test.total, test.fizzAt, test.buzzAt); !stringSlicesEqual(result, test.expected){
			t.Errorf("Output %q not equal to expected %q\n", result, test.expected)
		}
	}
}

func BenchmarkFizzBuzz(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FizzBuzz(15, 3, 5)
	}
}

func ExampleFizzBuzz() {
	fmt.Println(FizzBuzz(15, 3, 5))
	fmt.Println(FizzBuzz(0, 3, 5))
	// Output:
	// [1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz]
	// []
}
