package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/ludaNOFX/probability/internal"
	computersimulation "github.com/ludaNOFX/probability/internal/computer_simulation"
	"github.com/ludaNOFX/probability/internal/computer_simulation/app"
	"github.com/ludaNOFX/probability/internal/computer_simulation/infrastructure"
	"github.com/ludaNOFX/probability/internal/computer_simulation/infrastructure/random"
)

func main() {
	csv := "csv"
	plot := "plot"
	prj := internal.NewPrjMap(
		"computer_simulation",
		[]string{
			csv,
			plot,
		},
	)
	pathMap, err := computersimulation.SetupPaths(prj)
	if err != nil {
		log.Fatal(err)
	}
	uniformL1 := filepath.Join(pathMap[csv], "uniform_l1.csv")
	uniformL2 := filepath.Join(pathMap[csv], "uniform_l2.csv")
	normalL1 := filepath.Join(pathMap[csv], "normal_l1.csv")
	normalL2 := filepath.Join(pathMap[csv], "normal_l2.csv")

	uniformL1Q7 := filepath.Join(pathMap[plot], "uniform_l1_q7.png")
	uniformL1Q5 := filepath.Join(pathMap[plot], "uniform_l1_q5.png")
	uniformL2Q7 := filepath.Join(pathMap[plot], "uniform_l2_q7.png")
	uniformL2Q5 := filepath.Join(pathMap[plot], "uniform_l2_q5.png")

	normalL1Q7 := filepath.Join(pathMap[plot], "normal_l1_q7.png")
	normalL1Q5 := filepath.Join(pathMap[plot], "normal_l1_q5.png")
	normalL2Q7 := filepath.Join(pathMap[plot], "normal_l2_q7.png")
	normalL2Q5 := filepath.Join(pathMap[plot], "normal_l2_q5.png")

	N := 14
	l1, l2 := 100, 1000
	a := -float64(N) / 10.0
	b := float64(N) / 2.0

	ug := random.NewUniform(0)
	ng := random.NewBoxMuller(0)
	sim := app.NewNormalGenerator(ng)

	uniformSampleL1 := make([]float64, l1)
	uniformSampleL2 := make([]float64, l2)

	for i := 0; i < l1; i++ {
		u := ug.Float64()
		uniformSampleL1[i] = a + (b-a)*u
	}
	for i := 0; i < l2; i++ {
		u := ug.Float64()
		uniformSampleL2[i] = a + (b-a)*u
	}

	if err := infrastructure.SaveToFile(uniformSampleL1, uniformL1); err != nil {
		log.Fatal(err)
	}
	if err := infrastructure.SaveToFile(uniformSampleL2, uniformL2); err != nil {
		log.Fatal(err)
	}

	dataUniformL1, err := infrastructure.LoadFromFile(uniformL1)
	if err != nil {
		log.Fatal(err)
	}
	dataUniformL2, err := infrastructure.LoadFromFile(uniformL2)
	if err != nil {
		log.Fatal(err)
	}

	statsUniformL1 := app.ComputeStats(dataUniformL1)
	statsUniformL2 := app.ComputeStats(dataUniformL2)

	fmt.Println("Равномерное распределение")
	fmt.Printf("Параметры a=%.1f, b=%.0f\n\n", a, b)
	fmt.Printf("Статистика l1: %v\n", statsUniformL1)
	fmt.Printf("Статистика l2: %v\n\n", statsUniformL2)

	fmt.Println("Гистограмма q=7 l1:")
	for idx, hist := range app.Histogram(dataUniformL1, 7) {
		fmt.Printf("%d) %v\n", idx+1, hist)
	}
	fmt.Println()
	fmt.Println("Гистограмма q=5 l1:")
	for idx, hist := range app.Histogram(dataUniformL1, 5) {
		fmt.Printf("%d) %v\n", idx+1, hist)
	}
	fmt.Println()
	fmt.Println("Гистограмма q=7 l2:")
	for idx, hist := range app.Histogram(dataUniformL2, 7) {
		fmt.Printf("%d) %v\n", idx+1, hist)
	}
	fmt.Println()
	fmt.Println("Гистограмма q=5 l2:")
	for idx, hist := range app.Histogram(dataUniformL2, 5) {
		fmt.Printf("%d) %v\n", idx+1, hist)
	}
	fmt.Println()

	Mn := float64(N)
	sigma := float64(N) / 3.0

	normalStandardL1 := sim.Sample(l1)
	normalStandardL2 := sim.Sample(l2)

	normalSampleL1 := make([]float64, l1)
	normalSampleL2 := make([]float64, l2)

	for i := 0; i < l1; i++ {
		normalSampleL1[i] = Mn + sigma*normalStandardL1[i]
	}
	for i := 0; i < l2; i++ {
		normalSampleL2[i] = Mn + sigma*normalStandardL2[i]
	}

	if err := infrastructure.SaveToFile(normalSampleL1, normalL1); err != nil {
		log.Fatal(err)
	}
	if err := infrastructure.SaveToFile(normalSampleL2, normalL2); err != nil {
		log.Fatal(err)
	}

	dataNormalL1, err := infrastructure.LoadFromFile(normalL1)
	if err != nil {
		log.Fatal(err)
	}
	dataNormalL2, err := infrastructure.LoadFromFile(normalL2)
	if err != nil {
		log.Fatal(err)
	}

	statsNormalL1 := app.ComputeStats(dataNormalL1)
	statsNormalL2 := app.ComputeStats(dataNormalL2)

	err = infrastructure.PlotHistogram(dataUniformL1, uniformL1Q7, 7)
	if err != nil {
		log.Fatal(err)
	}
	err = infrastructure.PlotHistogram(dataUniformL1, uniformL1Q5, 5)
	if err != nil {
		log.Fatal(err)
	}
	err = infrastructure.PlotHistogram(dataUniformL2, uniformL2Q7, 7)
	if err != nil {
		log.Fatal(err)
	}
	err = infrastructure.PlotHistogram(dataUniformL2, uniformL2Q5, 5)
	if err != nil {
		log.Fatal(err)
	}

	err = infrastructure.PlotHistogram(dataNormalL1, normalL1Q7, 7)
	if err != nil {
		log.Fatal(err)
	}
	err = infrastructure.PlotHistogram(dataNormalL1, normalL1Q5, 5)
	if err != nil {
		log.Fatal(err)
	}
	err = infrastructure.PlotHistogram(dataNormalL2, normalL2Q7, 7)
	if err != nil {
		log.Fatal(err)
	}
	err = infrastructure.PlotHistogram(dataNormalL2, normalL2Q5, 5)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nНормальное распределение")
	fmt.Printf("Параметры M=%.0f, sigma=%.5f\n\n", Mn, sigma)
	fmt.Println("Статистика l1:", statsNormalL1)
	fmt.Println("Статистика l2:", statsNormalL2)

	fmt.Println("Гистограмма q=7 l1:")
	for idx, hist := range app.Histogram(dataNormalL1, 7) {
		fmt.Printf("%d) %v\n", idx+1, hist)
	}
	fmt.Println()
	fmt.Println("Гистограмма q=5 l1:")
	for idx, hist := range app.Histogram(dataNormalL1, 5) {
		fmt.Printf("%d) %v\n", idx+1, hist)
	}
	fmt.Println()
	fmt.Println("Гистограмма q=7 l2:")
	for idx, hist := range app.Histogram(dataNormalL2, 7) {
		fmt.Printf("%d) %v\n", idx+1, hist)
	}
	fmt.Println()
	fmt.Println("Гистограмма q=5 l2:")
	for idx, hist := range app.Histogram(dataNormalL2, 5) {
		fmt.Printf("%d) %v\n", idx+1, hist)
	}
}
