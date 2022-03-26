# Test EdenFarm - Adan
## Test questions :
Create a function to get the largest value from an array of values and
reappears later in the reverse order.
Example :
- [1, 2, 3, 8, 9, 3, 2, 1] largest value 3 (dari deret [1, 2, 3])
- [1, 2, 1, 2, 2, 1] → 2
- [7, 1, 2, 9, 7, 2, 1] → 2
Create in Go language and unit tests with at least 90% coverage

## Cloning Program :
Clone this program
```
  git clone https://github.com/aaidant/edenfarm-be-test-adan.git
```
Must get package after cloning
```
  make get
```

## Run Program :
Run this program
```
  make run
```
Input array
```
  1,2,3,8,9,3,2,1
```
Output expected
```
  Largest value of an array of consecutive values :  3
```

## Test Coverage :
Test this program
```
  make test
```
Test Coverage Result
```
  go test -v -cover -covermode=atomic ./...
  ?       edenfarm.com/test-edennfarm-adan       [no test files]
  ?       edenfarm.com/test-edennfarm-adan/entity        [no test files]
  === RUN   TestService_FindLargestValue
  --- PASS: TestService_FindLargestValue (0.00s)
  === RUN   TestService_FindLargestValue_Error
  --- PASS: TestService_FindLargestValue_Error (0.00s)
  === RUN   TestService_ConvertStringToArray
  --- PASS: TestService_ConvertStringToArray (0.00s)
  === RUN   TestService_CountTotalNum
  --- PASS: TestService_CountTotalNum (0.00s)
  === RUN   TestService_CheckInArrayWithValue
  --- PASS: TestService_CheckInArrayWithValue (0.00s)
  === RUN   TestService_GetLargestNumAndTotal
  --- PASS: TestService_GetLargestNumAndTotal (0.00s)
  PASS
  coverage: 100.0% of statements
  ok      edenfarm.com/test-edennfarm-adan/service       (cached)        coverage: 100.0% of statements
```
