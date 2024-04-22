package main

import (
	"testing"

	"github.com/montanaflynn/stats"
	"github.com/stretchr/testify/assert"
)

func BenchmarkLinearRegression1(b *testing.B) {
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

	for i := 0; i < b.N; i++ {
		_, _ = stats.LinearRegression(dataSet1)
	}
}

func BenchmarkLinearRegression2(b *testing.B) {
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

	for i := 0; i < b.N; i++ {
		_, _ = stats.LinearRegression(dataSet2)
	}
}

func BenchmarkLinearRegression3(b *testing.B) {
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

	for i := 0; i < b.N; i++ {
		_, _ = stats.LinearRegression(dataSet3)
	}
}

func BenchmarkLinearRegression4(b *testing.B) {
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

	for i := 0; i < b.N; i++ {
		_, _ = stats.LinearRegression(dataSet4)
	}
}

func TestLinearRegression1(t *testing.T) {
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

	slope1, intercept1, _ := calculateRegressionCoefficients(dataSet1)
	assert.InDelta(t, 0.5001, slope1, 0.0001)
	assert.InDelta(t, 3.0001, intercept1, 0.0001)
}

func TestLinearRegression2(t *testing.T) {
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

	slope2, intercept2, _ := calculateRegressionCoefficients(dataSet2)
	assert.InDelta(t, 0.5000, slope2, 0.0001)
	assert.InDelta(t, 3.0009, intercept2, 0.0001)
}

func TestLinearRegression3(t *testing.T) {
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

	slope3, intercept3, _ := calculateRegressionCoefficients(dataSet3)
	assert.InDelta(t, 0.4997, slope3, 0.0001)
	assert.InDelta(t, 3.0025, intercept3, 0.0001)
}

func TestLinearRegression4(t *testing.T) {
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

	slope4, intercept4, _ := calculateRegressionCoefficients(dataSet4)
	assert.InDelta(t, 0.4999, slope4, 0.0001)
	assert.InDelta(t, 3.0017, intercept4, 0.0001)
}
