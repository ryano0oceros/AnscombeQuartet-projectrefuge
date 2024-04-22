# AnscombeQuartet-projectrefuge
Project to ensure that the Go statistical package yields results comparable to those (estimated linear regression coefficients) obtained from Python and R

## Description
This project is a simple implementation of the Anscombe Quartet in Go. The Anscombe Quartet is a set of four datasets that have nearly identical statistical properties, but are visually distinct. The goal of this project is to ensure that the Go statistical package yields results comparable to those (estimated linear regression coefficients) obtained from Python and R.

## Recommendation to Management

Based on the results of this project, we recommend considering Go for data science applications. Our tests show that the Go statistical package yields results identical to those obtained from Python and R, which are widely used in data science. Furthermore, Go's performance is impressive; in my tests, Go's runtime was significantly faster than Python's (on the order of microseconds vs. seconds, although the python example had more functionality built in). In summary, I recommend using Go for data science applications where performance is critical and the necessary functionality is not too specialized. For more specialized data science tasks, Python or R might be more suitable.

### Caveats and Considerations
There are a few considerations to keep in mind. While Go has robust support for many general programming tasks, its data science ecosystem is not as mature as Python's or R's. Python and R have a vast array of libraries and tools specifically designed for data science, which Go currently lacks. This means that data scientists might have to implement some functionality from scratch, which could be time-consuming. Additionally, Go's syntax and paradigms are different from those of Python and R, which are more familiar to most data scientists. This could lead to a steeper learning curve for data scientists new to Go.


## Usage
To run the project, simply execute the following command:
```bash
go run main.go
```

To run the tests, execute the following command:
```bash
go test
```

## Outputs of `go run main.go` and `go test -bench .`
```bash
ryano0oceros@Ryans-MacBook-Pro-2 AnscombeQuartet-projectrefuge % go run main.go
Slope: 0.500091, Intercept: 3.000091
LinearRegression for dataSet1 took 273µs
Slope: 0.500000, Intercept: 3.000909
LinearRegression for dataSet2 took 5.5µs
Slope: 0.499727, Intercept: 3.002455
LinearRegression for dataSet3 took 3.792µs
Slope: 0.499909, Intercept: 3.001727
LinearRegression for dataSet4 took 3.625µs

ryano0oceros@Ryans-MacBook-Pro-2 AnscombeQuartet-projectrefuge % go test
Slope: 0.500091, Intercept: 3.000091
Slope: 0.500000, Intercept: 3.000909
Slope: 0.499727, Intercept: 3.002455
Slope: 0.499909, Intercept: 3.001727
PASS
ok      anscombequartet-projectrefug    0.338s
```