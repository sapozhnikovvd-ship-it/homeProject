package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func mySplit(data []byte, atEOF bool) (advance int, token []byte, err error) {
	sep := ","
	index := strings.Index(string(data), sep)

	if index >= 0 {
		return index + len(sep), data[:index], nil
	} else if len(data) > 0 {
		return len(data), data, nil
	}

	return 0, nil, nil
}

func inputNumberSeries() ([]int, error) {
	var result []int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(mySplit)

	for {
		if scanner.Scan() {
			num, err := strconv.ParseInt(strings.TrimSpace(scanner.Text()), 10, strconv.IntSize)

			if err == nil {
				result = append(result, int(num))
			} else {
				return make([]int, 0), err
			}
		} else if err := scanner.Err(); err != nil {
			return make([]int, 0), err
		} else {
			break
		}
	}

	return result, nil
}

func performingOperation(numberSeries []int, operation string) (int, error) {
	if !(len(numberSeries) > 0) {
		return -1, fmt.Errorf("Empty number series")
	}

	switch operation {
	case "AVG":
		sum := 0
		for _, value := range numberSeries {
			sum += value
		}
		return int(sum / len(numberSeries)), nil
	case "SUM":
		sum := 0
		for _, value := range numberSeries {
			sum += value
		}
		return sum, nil
	case "MED":
		copyNumberSeries := slices.Clone(numberSeries)
		slices.Sort(copyNumberSeries)
		n := len(copyNumberSeries)
		if n%2 == 1 {
			return copyNumberSeries[n/2], nil
		}
		return (copyNumberSeries[n/2-1] + copyNumberSeries[n/2]) / 2, nil
	default:
		return -1, fmt.Errorf("Unknow command")
	}
}

func main() {
	var command string
	fmt.Print("Input command (AVG, SUM, MED): ")
	_, err := fmt.Scanf("%s", &command)
	if err != nil {
		fmt.Printf("Invalid input: %v\n", err)
		os.Exit(1)
	}

	fmt.Print("Input number series: ")
	numberSeries, err := inputNumberSeries()
	if err != nil {
		fmt.Printf("Invalid input: %v\n", err)
		os.Exit(1)
	}

	result, err := performingOperation(numberSeries, command)
	if err != nil {
		fmt.Printf("Invalid input: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Result for \"%s\": %d\n", command, result)
}
