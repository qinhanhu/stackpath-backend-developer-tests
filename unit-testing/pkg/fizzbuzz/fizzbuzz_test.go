package fizzbuzz

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var fizzBuzzTests = []struct {
	name                  string
	total, fizzAt, buzzAt int64
	expected              []string
	wantErr               bool
	wantPanic             bool
}{
	// Write test cases here
	{"test empty", 0, 0, 0, []string{}, false, false},
	{"test err: divided by zero, fizzAt = 0", 1, 0, 1, []string{}, true, false},
	{"test err: divided by zero, buzzAt = 0", 1, 1, 0, []string{}, true, false},
	{"test result", 15, 3, 5, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}, false, false},
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

			result, err := FizzBuzz(test.total, test.fizzAt, test.buzzAt)
			if (err != nil) != test.wantErr {
				t.Errorf("unexpected err: err = %v, case = %v", err, test)
			} else if err != nil {
				errString := fmt.Sprintf("fizzAt or buzzAt equals zero, fizzAt = %v, buzzAt = %v", test.fizzAt, test.buzzAt)
				assert.EqualError(t, err, errString)
			}
			assert.Equal(t, result, test.expected)
		})
	}
}

func BenchmarkFizzBuzz(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = FizzBuzz(15, 3, 5)
	}
}

func ExampleFizzBuzz() {
	fmt.Println(FizzBuzz(15, 3, 5))
	fmt.Println(FizzBuzz(0, 3, 5))
	fmt.Println(FizzBuzz(1, 0, 1))
	fmt.Println(FizzBuzz(1, 1, 0))
	// Output:
	// [1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz] <nil>
	// [] <nil>
	// [] fizzAt or buzzAt equals zero, fizzAt = 0, buzzAt = 1
	// [] fizzAt or buzzAt equals zero, fizzAt = 1, buzzAt = 0
}
