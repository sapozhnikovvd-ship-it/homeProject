package bins

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
	"time"
)

type Bin struct {
	Id        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

type BinList []Bin

func NewBin(scanner *bufio.Scanner) (*Bin, error) {
	fmt.Print("Please! Input your id (number): ")
	id, err := scannerHendler(scanner)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %w", err)
	}

	fmt.Print("Please! Select private or public: ")
	choice, err := scannerHendler(scanner)
	if err != nil {
		return nil, fmt.Errorf("unacceptable choice private or not: %w", err)
	}

	var private bool
	if choice == "private" {
		private = true
	} else if choice == "public" {
		private = false
	} else {
		counter := 0
		for ; counter < 100; counter++ {
			fmt.Print("Invalid choice! Try again (private or public): ")
			choice, err = scannerHendler(scanner)
			if err != nil {
				return nil, fmt.Errorf("unacceptable choice private or not: %w", err)
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
			return nil, fmt.Errorf("so many try of input")
		}
	}

	fmt.Print("Please! Input your name: ")
	name, err := scannerHendler(scanner)
	if err != nil {
		return nil, fmt.Errorf("invalid input name: %w", err)
	}

	return &Bin{Id: id, Private: private, CreatedAt: time.Now(), Name: name}, nil
}

func scannerHendler(scanner *bufio.Scanner) (string, error) {
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return "", err
		}
		return "", errors.New("end of input")
	}

	return strings.TrimSpace(scanner.Text()), nil
}

func NewBinList(scanner *bufio.Scanner) (BinList, error) {

	var newBinSlice []Bin

	for {
		newBin, err := NewBin(scanner)
		if err != nil {
			return nil, err
		}

		newBinSlice = append(newBinSlice, (*newBin))
		fmt.Printf("Next? Y or anything else: ")

		if !scanner.Scan() {
			if err = scanner.Err(); err != nil {
				return nil, err
			}
			return nil, errors.New("end of input")
		}

		if strings.TrimSpace(scanner.Text()) != "Y" {
			break
		}
	}

	return BinList(newBinSlice), nil
}
