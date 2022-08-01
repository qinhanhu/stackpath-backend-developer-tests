package fizzbuzz

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

//type fizzBuzzTest struct {
//	name                  string
//	total, fizzAt, buzzAt int64
//	expected              []string
//	wantPanic             bool
//}

var fizzBuzzTests = []struct {
	name                  string
	total, fizzAt, buzzAt int64
	expected              []string
	wantPanic             bool
}{
	// Write test cases here
	{"test empty", 0, 0, 0, []string{}, false},
	{"test panic: divided by zero, fizzAt = 0", 1, 0, 1, []string{}, true},
	{"test panic: divided by zero, buzzAt = 0", 1, 1, 0, []string{}, true},
	{"test result", 15, 3, 5, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}, false},
}

func TestFizzBuzz(t *testing.T) {
	for _, test := range fizzBuzzTests {
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != test.wantPanic {
					t.Errorf("unexpected panic: recover = %v, case = %v", r, test)
				}
			}()
			//if result := FizzBuzz(test.total, test.fizzAt, test.buzzAt); !stringSlicesEqual(result, test.expected){
			//	t.Errorf("Output %q not equal to expected %q\n", result, test.expected)
			//}
			result := FizzBuzz(test.total, test.fizzAt, test.buzzAt)
			assert.Equal(t, result, test.expected)
		})
	}
}

//func integerDivideByZeroPanic(t *testing.T) {
//	defer func() {
//		if r := recover(); r == nil {
//			t.Errorf("The func did not panic as expected")
//		}
//	}()
//	a := 5
//	b := 0
//	_ = a / b
//}

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
