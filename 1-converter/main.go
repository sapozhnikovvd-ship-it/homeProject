package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var exchangeRate = map[string]float64{"USD": 1.0, "EUR": 0.932, "RUB": 70.323}

// Форматированный ввод пользователем информации для конверсии валют
func inputConvertInfo() (amount float64, sourceCurrency string, finalCurrency string) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Input amount: ")

		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Printf("Invalid input amount: %v\nTry again\n", err)
			}
			fmt.Println("Goodbye!")
			os.Exit(1)
		}

		line := strings.TrimSpace(scanner.Text())
		amountLocal, err := strconv.ParseFloat(line, 64)

		if err != nil {
			fmt.Printf("Invalid input amount: %v\nTry again\n", err)
			continue
		}

		if amountLocal < 0 {
			fmt.Printf("Invalid input: the amount must be equal to zero or be positive\nTry again\n")
			continue
		}
		amount = amountLocal
		break
	}

	for {
		fmt.Print("Input source currency (RUB, EUR, USD): ")
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Printf("Invalid input source currency: %v\nTry again\n", err)
				continue
			}
			fmt.Println("Goodbye!")
			os.Exit(1)
		}

		sourceCurrency = strings.TrimSpace(scanner.Text())

		if _, ok := exchangeRate[sourceCurrency]; !ok {
			fmt.Printf("Invalid input: the source currency must be RUB, EUR or USD\nTry again\n")
			continue
		}
		break
	}

	for {
		fmt.Print("Input final currency (RUB, EUR, USD): ")
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Printf("Invalid input final currency: %v\nTry again\n", err)
				continue
			}
			fmt.Println("Goodbye!")
			os.Exit(1)
		}

		finalCurrency = strings.TrimSpace(scanner.Text())

		if _, ok := exchangeRate[finalCurrency]; !ok {
			fmt.Printf("Invalid input: the final currency must be RUB, EUR or USD\nTry again\n")
			continue
		}
		break
	}
	return
}

// Расчет эквивалента в другой валюте
func convertCurrency(amount float64, sourceCurrency string, finalCurrency string) float64 {
	return amount / exchangeRate[sourceCurrency] * exchangeRate[finalCurrency]
}

func main() {
	fmt.Printf("Result: %f\n", convertCurrency(inputConvertInfo()))
}
