package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"sync"
	"time"

	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/stat"
)

func loadCSV2(filename string) ([][]float64, error) {
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

func fitLinearModel2(x, y []float64) (alpha, beta, mse float64) {
	alpha, beta = stat.LinearRegression(x, y, nil, false)
	residuals := make([]float64, len(y))
	for i := range residuals {
		residuals[i] = math.Pow(y[i]-(alpha+beta*x[i]), 2)
	}
	mse = floats.Sum(residuals) / float64(len(y))
	return alpha, beta, mse
}

func fitLinearModelConcurrent(wg *sync.WaitGroup, data [][]float64, y []float64, i int, results chan<- string) {
	defer wg.Done()

	x := make([]float64, len(data))
	for j, row := range data {
		x[j] = row[i]
	}

	alpha, beta, mse := fitLinearModel2(x, y)
	result := fmt.Sprintf("Feature %d: Alpha: %f, Beta: %f, MSE: %f", i, alpha, beta, mse)
	results <- result
}

func main() {
	data, err := loadCSV2("C:/Users/haoyx/Desktop/MSDS431/boston.csv")
	if err != nil {
		log.Fatalf("error loading CSV: %v", err)
	}

	y := make([]float64, len(data))
	for i, row := range data {
		y[i] = row[len(row)-1]
	}

	var totalTime time.Duration

	for n := 0; n < 100; n++ {
		start := time.Now()

		results := make(chan string)
		var wg sync.WaitGroup

		for i := 0; i < len(data[0])-1; i++ {
			wg.Add(1)
			go fitLinearModelConcurrent(&wg, data, y, i, results)
		}

		go func() {
			wg.Wait()
			close(results)
		}()

		// Collect results
		for result := range results {
			fmt.Println(result)
		}
		totalTime += time.Since(start)
	}
	avgTime := totalTime / 100
	fmt.Printf("Total execution time over 100 runs: %v\n", totalTime)
	fmt.Printf("Average execution time per run: %v\n", avgTime)
}
