package infrastructure

import (
	"encoding/csv"
	"os"
	"strconv"
)

// SaveToFile сохраняет срез float64 в csv-файл (одна колонка)
func SaveToFile(data []float64, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, v := range data {
		if err := writer.Write([]string{strconv.FormatFloat(v, 'f', 6, 64)}); err != nil {
			return err
		}
	}
	return nil
}

// LoadFromFile читает csv-файл, содержащий по одному числу в строке
func LoadFromFile(filename string) ([]float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	data := make([]float64, len(lines))
	for i, line := range lines {
		v, err := strconv.ParseFloat(line[0], 64)
		if err != nil {
			return nil, err
		}
		data[i] = v
	}
	return data, nil
}
