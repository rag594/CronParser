## CronParser

### Package structure

`models/` package contains the basic cron modelling \
`parser/` package contains the different parsers for different parts(minutes, hours etc) \
`cronExpressionParser.go`  file has the cron expression parser structure \
`main.go` is the driver file

Each parts(hours, minutes etc) has different parsers. Month parser looks like below:

```go
type ICronMonthParser interface {
    WildCardParser() []int // parses the wildcard(* * * * *)
    NumberParser(values []int) []int // parses just the number(* * 5 * *)
    StepParser(values []int) []int // parses the steps(* * */5 * *)
    RangeParser(values []int) []int // parses the range(* * 4-8 * *)
    Parse() []int // parse constructs the pipeline using the above functions and additional business logic if any.
}
```


### Build
``
go build -o cronParser
``

### Run

``
./cronParser <cron_expression> <command>
``

```
foo$: ./cronParser "*/15 0 1-15 * 2-4" /usr/bin/find
minute         [0 15 30 45]
hour           [0]
day of month   [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15]
month          [1 2 3 4 5 6 7 8 9 10 11 12]
day of week    [2 3 4]
command        /usr/bin/find

```

### Tests

#### Run suite
``
go test -v ./...
``

#### Minute Parser tests
``
go test -v ./parser -run TestParseMinutes
``

#### Hour Parser tests
``
go test -v ./parser -run TestParseHour
``

#### Day of Month Parser tests
``
go test -v ./parser -run TestParseDayOfMonth
``

#### Month Parser tests
``
go test -v ./parser -run TestParseMonth
``

#### Day of Week Parser tests
``
go test -v ./parser -run TestParseDayOfWeek
``