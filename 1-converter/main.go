package main

import "fmt"

// Форматированный ввод пользователем информации для конверсии валют
func inputConvertInfo() (amount float64, sourceCurrency string, finalCurrency string) {
	fmt.Scanf("%f %s %s", &amount, &sourceCurrency, &finalCurrency)
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
	}

	return amount
}

func main() {
	fmt.Printf("Result: %f\n", convertCurrency(inputConvertInfo()))
}
