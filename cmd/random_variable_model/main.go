package main

import (
	"fmt"
	"log"

	randomvariablemodel "github.com/ludaNOFX/probability/internal/random_variable_model"
)

func main() {
	fmt.Print("Введите реальное число из своего варианта: ")
	inputD, err := randomvariablemodel.ScanRealDigit(randomvariablemodel.GetStdin())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Введите минимальную границу (default -0.1): ")
	xmin, err := randomvariablemodel.ScanRealDigit(randomvariablemodel.GetStdin())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Введите максимальную границу (default 1.1): ")
	xmax, err := randomvariablemodel.ScanRealDigit(randomvariablemodel.GetStdin())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Введите шаг (default 0.01): ")
	step, err := randomvariablemodel.ScanRealDigit(randomvariablemodel.GetStdin())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Введите ширину графика (default 600.0): ")
	weight, err := randomvariablemodel.ScanRealDigit(randomvariablemodel.GetStdin())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Введите высоту графика (default 400.0): ")
	height, err := randomvariablemodel.ScanRealDigit(randomvariablemodel.GetStdin())
	if err != nil {
		log.Fatal(err)
	}

	params := randomvariablemodel.NewParams(inputD, randomvariablemodel.CalcC)
	fmt.Printf("Константа c = %.4f\n", params.C)

	density := randomvariablemodel.GeneratePoints(params, xmin, xmax, step, randomvariablemodel.Fx)
	cdf := randomvariablemodel.GeneratePoints(params, xmin, xmax, step, randomvariablemodel.F)

	metaDensity := randomvariablemodel.NewMetaInfo(
		"Density",
		"x",
		"f(x)",
		weight,
		height,
	)

	mean := randomvariablemodel.Mean(params)
	median := randomvariablemodel.Median(params)
	mode := randomvariablemodel.Mode(params)
	variance := randomvariablemodel.Variance(params)
	sigma := randomvariablemodel.Sigma(params)

	fmt.Println("M =", mean)
	fmt.Println("Me =", median)
	fmt.Println("Mo =", mode)
	fmt.Println("D =", variance)
	fmt.Println("σ =", sigma)

	err = randomvariablemodel.Plot(density, "density.png", metaDensity, mean, median, mode)
	if err != nil {
		log.Fatal(err)
	}

	metaCdf := randomvariablemodel.NewMetaInfo(
		"CDF",
		"x",
		"F(x)",
		weight,
		height,
	)

	err = randomvariablemodel.Plot(cdf, "cdf.png", metaCdf, mean, median, mode)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n=== ВЫВОДЫ ПО РАБОТЕ ===")
	fmt.Printf("1. Для a = %.1f нормировочная константа c = %.4f\n", inputD, params.C)
	fmt.Printf("2. Плотность распределения: f(x) = %.4f(x + %.1f) на [0, 1]\n", params.C, params.A)
	fmt.Printf("3. Функция распределения: F(x) = %.4f(0.5x² + %.1fx) на [0, 1]\n", params.C, params.A)
	fmt.Printf("4. Математическое ожидание M = %.4f\n", mean)
	fmt.Printf("5. Медиана Me = %.4f\n", median)
	fmt.Printf("6. Мода Mo = %.4f\n", mode)
	fmt.Printf("7. Дисперсия D = %.4f, СКО σ = %.4f\n", variance, sigma)

	if mean > median {
		fmt.Printf("8. M > Me (распределение правостороннее)\n")
	} else if mean < median {
		fmt.Printf("8. M < Me (распределение левостороннее)\n")
	} else {
		fmt.Printf("8. M = Me (симметричное распределение)\n")
	}

	fmt.Printf("9. Все характеристики отображены на графиках\n")
}
