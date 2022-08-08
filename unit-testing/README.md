# Test Result
```
❯ go test -v -bench=.
=== RUN   TestFizzBuzz
=== RUN   TestFizzBuzz/test_empty
=== RUN   TestFizzBuzz/test_err:_divided_by_zero,_fizzAt_=_0
=== RUN   TestFizzBuzz/test_err:_divided_by_zero,_buzzAt_=_0
=== RUN   TestFizzBuzz/test_result
--- PASS: TestFizzBuzz (0.00s)
    --- PASS: TestFizzBuzz/test_empty (0.00s)
    --- PASS: TestFizzBuzz/test_err:_divided_by_zero,_fizzAt_=_0 (0.00s)
    --- PASS: TestFizzBuzz/test_err:_divided_by_zero,_buzzAt_=_0 (0.00s)
    --- PASS: TestFizzBuzz/test_result (0.00s)
=== RUN   ExampleFizzBuzz
--- PASS: ExampleFizzBuzz (0.00s)
goos: darwin
goarch: amd64
pkg: github.com/stackpath/backend-developer-tests/unit-testing/pkg/fizzbuzz
cpu: VirtualApple @ 2.50GHz
BenchmarkFizzBuzz
BenchmarkFizzBuzz-10    	 7807476	       145.7 ns/op
PASS
ok  	github.com/stackpath/backend-developer-tests/unit-testing/pkg/fizzbuzz	1.513s
```
