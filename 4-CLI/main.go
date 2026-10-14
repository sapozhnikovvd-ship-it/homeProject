package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Bin struct {
	Id        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

func CreatBinAndFilling() *Bin {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Please! Input your id (number): ")
	id, err := scannerHendler(scanner)
	if err != nil {
		fmt.Printf("Stop programm: %v. Goodbye!\n", err)
		os.Exit(1)
	}

	fmt.Println("Please! Select private or public: ")
	choice, err := scannerHendler(scanner)
	if err != nil {
		fmt.Printf("Stop programm: %v. Goodbye!\n", err)
		os.Exit(1)
	}

	var private bool
	if choice == "private" {
		private = true
	} else if choice == "public" {
		private = false
	} else {
		counter := 0
		for ; counter < 100; counter++ {
			fmt.Printf("Invalid choice! Try again (private or public):")
			choice, err = scannerHendler(scanner)
			if err != nil {
				fmt.Printf("Stop programm: %v. Goodbye!\n", err)
				os.Exit(1)
			}

			if choice == "private" {
				private = true
				break
			} else if choice == "public" {
				private = false
				break
			}
		}

		if counter == 100 {
			fmt.Printf("Stop programm: so many try of input. Goodbye!\n", err)
			os.Exit(1)
		}
	}

	fmt.Println("Please! Input your name: ")
	name, err := scannerHendler(scanner)
	if err != nil {
		fmt.Printf("Stop programm: %v. Goodbye!\n", err)
		os.Exit(1)
	}

	return &Bin{Id: id, Private: private, CreatedAt: time.Now(), Name: name}
}

func scannerHendler(scanner *bufio.Scanner) (string, error) {
	counter := 0
	for ; counter < 100; counter++ {
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Println("Invalid input! Try again: ")
				continue
			}
			return "", errors.New("end of input")
		}
		break
	}
	if counter == 100 {
		return "", errors.New("so many try of input")
	}
	return strings.TrimSpace(scanner.Text()), nil
}

func NewBinList(lenOfBin, capOfBin int) *[]Bin {
	newBinList := make([]Bin, lenOfBin, capOfBin)
	return &newBinList
}

func NewBin() *Bin {
	return &Bin{}
}

func main() {
}
