package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"time"

	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/stat"
)

func loadCSV(filename string) ([][]float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rawCSVData, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var data [][]float64
	for i, record := range rawCSVData {
		if i == 0 { // Skip header row
			continue
		}
		var floatRecord []float64
		for j, value := range record {
			if j == 0 { // Assuming 'Nahant' is in the first column, skip it
				continue
			}
			f, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return nil, fmt.Errorf("parsing error on row %d, column %d: %v", i+1, j+1, err)
			}
			floatRecord = append(floatRecord, f)
		}
		data = append(data, floatRecord)
	}

	return data, nil
}

// fitLinearModel fits a linear regression model and calculates MSE for the provided subset of data.
func fitLinearModel(x, y []float64) (alpha, beta, mse float64) {
	alpha, beta = stat.LinearRegression(x, y, nil, false)
	residuals := make([]float64, len(y))
	for i := range residuals {
		residuals[i] = math.Pow(y[i]-(alpha+beta*x[i]), 2)
	}
	mse = floats.Sum(residuals) / float64(len(y))
	return alpha, beta, mse
}

func main() {
	// Load the dataset from CSV
	data, err := loadCSV("C:/Users/haoyx/Desktop/MSDS431/boston.csv")
	if err != nil {
		log.Fatalf("error loading CSV: %v", err)
	}

	y := make([]float64, len(data))
	for i, row := range data {
		y[i] = row[len(row)-1]
	}

	// Fit a model for each feature individually
	for i := 0; i < len(data[0])-1; i++ {
		x := make([]float64, len(data))
		for j, row := range data {
			x[j] = row[i]
		}

		alpha, beta, mse := fitLinearModel(x, y)
		fmt.Printf("Feature %d: Beta: %f, MSE: %f\n", i, alpha, beta, mse)
	}
	var totalTime time.Duration
	for i := 0; i < 100; i++ {
		start := time.Now()
		// ... (code to fit models without concurrency)
		totalTime += time.Since(start)
	}
	avgTime := totalTime / 100
	fmt.Printf("Average execution time (non-concurrent): %v\n", avgTime)
}
