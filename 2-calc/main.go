package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readLine(scanner *bufio.Scanner) (string, error) {
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return "", err
		}
		return "", io.EOF
	}

	return strings.TrimSpace(scanner.Text()), nil
}

func parseNumberSeries(line string) ([]int, error) {
	var result []int

	for _, part := range strings.Split(line, ",") {
		part = strings.TrimSpace(part)

		if part == "" {
			continue
		}

		number, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		result = append(result, number)
	}

	return result, nil
}

func performingOperation(numberSeries []int, operation string) (float64, error) {
	if !(len(numberSeries) > 0) {
		return -1, fmt.Errorf("empty number series")
	}

	switch operation {
	case "AVG":
		sum := 0
		for _, value := range numberSeries {
			sum += value
		}
		return float64(sum) / float64(len(numberSeries)), nil
	case "SUM":
		sum := 0
		for _, value := range numberSeries {
			sum += value
		}
		return float64(sum), nil
	case "MED":
		copyNumberSeries := slices.Clone(numberSeries)
		slices.Sort(copyNumberSeries)
		n := len(copyNumberSeries)
		if n%2 == 1 {
			return float64(copyNumberSeries[n/2]), nil
		}
		return float64(copyNumberSeries[n/2-1]+copyNumberSeries[n/2]) / 2.0, nil
	default:
		return -1, fmt.Errorf("unknow command")
	}
}

func main() {
	var command string
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < 101; i++ {
		if i == 100 {
			fmt.Printf("So mach try!\n")
			os.Exit(1)
		}

		fmt.Print("Input command (AVG, SUM, MED): ")
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Printf("Invalid input: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("End of input\n")
			os.Exit(1)
		}

		command = strings.TrimSpace(scanner.Text())
		if slices.Contains([]string{"AVG", "SUM", "MED"}, command) {
			break
		}
		fmt.Println("Unknown command! Try again!")
	}

	var numberSeries []int
	for i := 0; i < 101; i++ {
		if i == 100 {
			fmt.Printf("So mach try!\n")
			os.Exit(1)
		}

		fmt.Print("Input number series: ")
		lineOfNuberSeries, err := readLine(scanner)
		if err != nil {
			fmt.Printf("Invalid input: %v\n", err)
			os.Exit(1)
		}

		numberSeries, err = parseNumberSeries(lineOfNuberSeries)
		if err != nil {
			fmt.Printf("Invalid input: %v\n", err)
			continue
		}
		if numberSeries == nil {
			continue
		}

		break
	}

	result, err := performingOperation(numberSeries, command)
	if err != nil {
		fmt.Printf("Invalid input: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Result for \"%s\": %4.2f\n", command, result)
}
