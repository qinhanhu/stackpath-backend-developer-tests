# Project Structure
```
❯ tree -L 2
.
├── controllers         // gather handlers here
│   └── people.go
├── go.mod
├── go.sum
├── main.go
├── pkg
│   └── models          // gather models here
├── routers             // gather routers here
│   ├── main.go
│   └── people.go
└── tests               // gather tests here
    ├── main.go
    ├── people_test.go
    └── postman
```
# Tests
* Use Postman
* Use go test
```
❯ go test -v
=== RUN   TestPeople
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /people                   --> github.com/stackpath/backend-developer-tests/rest-service/controllers.GetPeople (3 handlers)
[GIN-debug] GET    /people/:id               --> github.com/stackpath/backend-developer-tests/rest-service/controllers.GetPeopleByID (3 handlers)
=== RUN   TestPeople/test_GetPeople_All
[GIN] 2022/08/06 - 23:46:25 | 200 |     405.167µs |       192.0.2.1 | GET      "/people"
=== RUN   TestPeople/test_GetPeople_ByName_DataExist
[GIN] 2022/08/06 - 23:46:25 | 200 |     109.416µs |       192.0.2.1 | GET      "/people?first_name=Brian&last_name=Smith"
=== RUN   TestPeople/test_GetPeople_ByName_DataEmpty
[GIN] 2022/08/06 - 23:46:25 | 200 |       2.125µs |       192.0.2.1 | GET      "/people?first_name=&last_name="
=== RUN   TestPeople/test_GetPeople_ByPhoneNumber_DataExist
[GIN] 2022/08/06 - 23:46:25 | 200 |       9.333µs |       192.0.2.1 | GET      "/people?phone_number=%2B1+%28800%29+555-1313"
=== RUN   TestPeople/test_GetPeople_ByPhoneNumber_DataEmpty
[GIN] 2022/08/06 - 23:46:25 | 200 |         1.5µs |       192.0.2.1 | GET      "/people?phone_number=112233"
=== RUN   TestPeople/test_GetPeopleByID_DataExist
[GIN] 2022/08/06 - 23:46:25 | 200 |       2.833µs |       192.0.2.1 | GET      "/people/5b81b629-9026-450d-8e46-da4f8c7bd513"
=== RUN   TestPeople/test_GetPeopleByID_404NotFound
[GIN] 2022/08/06 - 23:46:25 | 404 |       7.042µs |       192.0.2.1 | GET      "/people/112233"
--- PASS: TestPeople (0.00s)
    --- PASS: TestPeople/test_GetPeople_All (0.00s)
    --- PASS: TestPeople/test_GetPeople_ByName_DataExist (0.00s)
    --- PASS: TestPeople/test_GetPeople_ByName_DataEmpty (0.00s)
    --- PASS: TestPeople/test_GetPeople_ByPhoneNumber_DataExist (0.00s)
    --- PASS: TestPeople/test_GetPeople_ByPhoneNumber_DataEmpty (0.00s)
    --- PASS: TestPeople/test_GetPeopleByID_DataExist (0.00s)
    --- PASS: TestPeople/test_GetPeopleByID_404NotFound (0.00s)
PASS
ok  	github.com/stackpath/backend-developer-tests/rest-service/tests	0.302s
```
