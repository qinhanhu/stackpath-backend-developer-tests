# Input Processing
## Test Result
```
❯ go test -v
=== RUN   TestErrorDetect-developer-tests/input-processing │ input_processing !3 ?2  go test -v              ✔ │ 6s │ base  │ 22:13:15 
=== RUN   TestErrorDetect/Detected_error_in_multilines
=== RUN   TestErrorDetect/No_error_detected
=== RUN   TestErrorDetect/Detected_Random_Large_Data
=== RUN   TestErrorDetect/Detected_Random_Large_Data/Detected_Random_Large_Data
=== RUN   TestErrorDetect/Detected_Large_Data_with_error
=== RUN   TestErrorDetect/Detected_Large_Data_with_error/Detected_Large_Data_with_error
--- PASS: TestErrorDetect (5.64s)
    --- PASS: TestErrorDetect/Detected_error_in_multilines (0.00s)
    --- PASS: TestErrorDetect/No_error_detected (0.00s)
    --- PASS: TestErrorDetect/Detected_Random_Large_Data (0.51s)
        --- PASS: TestErrorDetect/Detected_Random_Large_Data/Detected_Random_Large_Data (0.22s)
    --- PASS: TestErrorDetect/Detected_Large_Data_with_error (5.13s)
        --- PASS: TestErrorDetect/Detected_Large_Data_with_error/Detected_Large_Data_with_error (2.21s)
PASS
ok      github.com/stackpath/backend-developer-tests/input-processing   5.892s

```
## Some questions you might have
1. How to process large size input line by line? 
   - https://devmarkpro.com/working-big-files-golang
2. How to mock `Stdin` in unit test?  
   - https://stackoverflow.com/questions/46365221/fill-os-stdin-for-function-that-reads-from-it
3. How to mock `Stdout` in unit test?
   - https://stackoverflow.com/questions/26937770/golang-test-stdout
4. How to generate largesize byte array with random content? 
   - https://stackoverflow.com/questions/35781197/generating-a-random-fixed-length-byte-array-in-go
