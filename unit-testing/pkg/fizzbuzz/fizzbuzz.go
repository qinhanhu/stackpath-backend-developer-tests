package fizzbuzz

import (
	"fmt"
	"strconv"
)

// FizzBuzz performs a FizzBuzz operation over a range of integers
//
// Given a range of integers:
// - Return "Fizz" if the integer is divisible by the `fizzAt` value.
// - Return "Buzz" if the integer is divisible by the `buzzAt` value.
// - Return "FizzBuzz" if the integer is divisible by both the `fizzAt` and
//   `buzzAt` values.
// - Return the original number if it is not divisible by either the `fizzAt` or
//   the `buzzAt` values.
func FizzBuzz(total, fizzAt, buzzAt int64) ([]string, error) {
	result := make([]string, total)
	for i := int64(1); i <= total; i++ {
		if fizzAt == 0 || buzzAt == 0 {
			return []string{}, fmt.Errorf("fizzAt or buzzAt equals zero, fizzAt = %v, buzzAt = %v", fizzAt, buzzAt)
		}
		if !(i%fizzAt == 0) && !(i%buzzAt == 0) {
			result[i-1] = strconv.FormatInt(i, 10)
			continue
		}

		if i%fizzAt == 0 {
			result[i-1] = "Fizz"
		}

		if i%buzzAt == 0 {
			result[i-1] += "Buzz"
		}
	}

	return result, nil
}
