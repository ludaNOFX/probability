package randomvariablemodel

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
)

func ScanRealDigit(r io.Reader) (float64, error) {
	reader := bufio.NewScanner(r)
	var d float64
	count := 0
	for reader.Scan() {
		t := reader.Text()
		var err error
		d, err = strconv.ParseFloat(t, 64)
		if err != nil {
			fmt.Println("Input not real digit")
			if count > 10 {
				return 0, errors.New("Max retries over")
			}
			continue
		}
		break
	}
	return d, nil
}
