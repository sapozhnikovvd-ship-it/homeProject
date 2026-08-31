package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Форматированный ввод пользователем информации для конверсии валют
func inputConvertInfo() (amount float64, sourceCurrency string, finalCurrency string) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Input amount: ")

		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Printf("Invalid input amount: %v\nTry again\n", err)
				continue
			}
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
		}

		sourceCurrency = strings.TrimSpace(scanner.Text())

		if sourceCurrency != "RUB" && sourceCurrency != "EUR" && sourceCurrency != "USD" {
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
		}

		finalCurrency = strings.TrimSpace(scanner.Text())

		if finalCurrency != "RUB" && finalCurrency != "EUR" && finalCurrency != "USD" {
			fmt.Printf("Invalid input: the final currency must be RUB, EUR or USD\nTry again\n")
			continue
		}
		break
	}
	return
}

func convertCurrency(amount float64, sourceCurrency string, finalCurrency string) float64 {
	const USDToEUR float64 = 0.932
	const USDToRUB float64 = 70.323
	const EURToRUB float64 = USDToRUB / USDToEUR

	switch sourceCurrency {
	case "RUB":
		if finalCurrency == "USD" {
			return amount / USDToRUB
		} else if finalCurrency == "EUR" {
			return amount / EURToRUB
		}
	case "USD":
		if finalCurrency == "RUB" {
			return amount * USDToRUB
		} else if finalCurrency == "EUR" {
			return amount * USDToEUR
		}
	case "EUR":
		if finalCurrency == "USD" {
			return amount / USDToEUR
		} else if finalCurrency == "RUB" {
			return amount * EURToRUB
		}
	default:
		return amount
	}

	return amount
}

func main() {
	fmt.Printf("Result: %f\n", convertCurrency(inputConvertInfo()))
}
