package main

import (
	"fmt"
	"time"

	"github.com/montanaflynn/stats"
)

func main() {
	// Define your data sets
	dataSet1 := []stats.Coordinate{
		{X: 10, Y: 8.04},
		{X: 8, Y: 6.95},
		{X: 13, Y: 7.58},
		{X: 9, Y: 8.81},
		{X: 11, Y: 8.33},
		{X: 14, Y: 9.96},
		{X: 6, Y: 7.24},
		{X: 4, Y: 4.26},
		{X: 12, Y: 10.84},
		{X: 7, Y: 4.82},
		{X: 5, Y: 5.68},
	}

	dataSet2 := []stats.Coordinate{
		{X: 10, Y: 9.14},
		{X: 8, Y: 8.14},
		{X: 13, Y: 8.74},
		{X: 9, Y: 8.77},
		{X: 11, Y: 9.26},
		{X: 14, Y: 8.1},
		{X: 6, Y: 6.13},
		{X: 4, Y: 3.1},
		{X: 12, Y: 9.13},
		{X: 7, Y: 7.26},
		{X: 5, Y: 4.74},
	}

	dataSet3 := []stats.Coordinate{
		{X: 10, Y: 7.46},
		{X: 8, Y: 6.77},
		{X: 13, Y: 12.74},
		{X: 9, Y: 7.11},
		{X: 11, Y: 7.81},
		{X: 14, Y: 8.84},
		{X: 6, Y: 6.08},
		{X: 4, Y: 5.39},
		{X: 12, Y: 8.15},
		{X: 7, Y: 6.42},
		{X: 5, Y: 5.73},
	}

	dataSet4 := []stats.Coordinate{
		{X: 8, Y: 6.58},
		{X: 8, Y: 5.76},
		{X: 8, Y: 7.71},
		{X: 8, Y: 8.84},
		{X: 8, Y: 8.47},
		{X: 8, Y: 7.04},
		{X: 8, Y: 5.25},
		{X: 19, Y: 12.5},
		{X: 8, Y: 5.56},
		{X: 8, Y: 7.91},
		{X: 8, Y: 6.89},
	}

	measureExecutionTime(func() { calculateRegressionCoefficients(dataSet1) }, "LinearRegression for dataSet1")
	measureExecutionTime(func() { calculateRegressionCoefficients(dataSet2) }, "LinearRegression for dataSet2")
	measureExecutionTime(func() { calculateRegressionCoefficients(dataSet3) }, "LinearRegression for dataSet3")
	measureExecutionTime(func() { calculateRegressionCoefficients(dataSet4) }, "LinearRegression for dataSet4")

}

func calculateRegressionCoefficients(dataSet []stats.Coordinate) (float64, float64, error) {
	_, err := stats.LinearRegression(dataSet)
	if err != nil {
		return 0, 0, err
	}

	sumX, sumY, sumXY, sumX2 := 0.0, 0.0, 0.0, 0.0
	for _, point := range dataSet {
		sumX += point.X
		sumY += point.Y
		sumXY += point.X * point.Y
		sumX2 += point.X * point.X
	}

	n := float64(len(dataSet))
	slope := (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)
	intercept := (sumY - slope*sumX) / n

	fmt.Printf("Slope: %f, Intercept: %f\n", slope, intercept)
	return slope, intercept, nil
}

func measureExecutionTime(f func(), label string) {
	start := time.Now()
	f()
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", label, elapsed)
}
