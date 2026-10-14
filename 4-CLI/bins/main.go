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

func NewBin(id string, private bool, createdAt time.Time, name string) *Bin {
	return &Bin{Id: id, Private: private, CreatedAt: createdAt, Name: name}
}

func NewBinConsole(scanner *bufio.Scanner) (*Bin, error) {
	fmt.Print("Please! Input your id: ")
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

type BinList []Bin

func NewBinList(bins ...Bin) *BinList {
	newBinList := BinList(bins)
	return &newBinList
}

func NewBinListConsole(scanner *bufio.Scanner) (BinList, error) {

	var newBinSlice []Bin
	errEOF := errors.New("end of input")

	for {
		newBin, err := NewBinConsole(scanner)
		if err != nil {
			return nil, err
		}

		newBinSlice = append(newBinSlice, (*newBin))
		fmt.Printf("Next? Y or anything else: ")

		if value, err := scannerHendler(scanner); value != "Y" && err == nil {
			break
		} else if err != nil {
			if errors.Is(err, errEOF) {
				return BinList(newBinSlice), err
			}
			return nil, err
		}
	}

	return BinList(newBinSlice), nil
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
