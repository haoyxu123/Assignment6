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
		if i == 0 {
			continue
		}
		var floatRecord []float64
		for j, value := range record {
			if j == 0 {
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
	startTime := time.Now()
	for i := 0; i < 100; i++ {
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
		for j := 0; j < 1000; j++ {
			_ = j * j
		}
	}
	elapsedTime := time.Since(startTime)

	fmt.Printf("Total execution time for 100 runs: %s\n", elapsedTime)
	fmt.Printf("Average execution time per run: %s\n", elapsedTime/time.Duration(100))
}

