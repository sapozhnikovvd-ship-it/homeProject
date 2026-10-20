package bins

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Bin struct {
	Id         string    `json:"Id"`
	Name       string    `json:"Name"`
	Private    bool      `json:"Private"`
	CreatedAt  time.Time `json:"CreatedAt"`
	ModifiedAt time.Time `json:"ModifiedAt"`
}

func NewBin(id string, private bool, name string) *Bin {
	return &Bin{Id: id, Name: name, Private: private, CreatedAt: time.Now(), ModifiedAt: time.Now()}
}

func NewBinConsole(scanner *bufio.Scanner) (*Bin, error) {
	fmt.Print("Please! Input your id: ")
	id, err := scannerHendler(scanner)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %w", err)
	}

	fmt.Print("Please! Input your name: ")
	name, err := scannerHendler(scanner)
	if err != nil {
		return nil, fmt.Errorf("invalid input name: %w", err)
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

	return &Bin{Id: id, Name: name, Private: private, CreatedAt: time.Now(), ModifiedAt: time.Now()}, nil
}

func NewBinFromJson(jsonName string) (*Bin, error) {
	newBinFromJson := &Bin{}

	bytesFromJson, err := os.ReadFile(jsonName)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytesFromJson, newBinFromJson)
	if err != nil {
		return nil, err
	}

	return newBinFromJson, nil
}

type BinList struct {
	List       []Bin     `json:"List"`
	CreatedAt  time.Time `json:"CreatedAt"`
	ModifiedAt time.Time `json:"ModifiedAt"`
}

func NewBinList(bins ...Bin) *BinList {
	newBinList := BinList{List: bins, CreatedAt: time.Now(), ModifiedAt: time.Now()}
	return &newBinList
}

func NewBinListConsole(scanner *bufio.Scanner) (*BinList, error) {

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
				return &BinList{List: newBinSlice, CreatedAt: time.Now(), ModifiedAt: time.Now()}, err
			}
			return nil, err
		}
	}

	return &BinList{List: newBinSlice, CreatedAt: time.Now(), ModifiedAt: time.Now()}, nil
}

func NewBinListFromJson(jsonName string) (*BinList, error) {
	binListFromJson := &BinList{}

	bytesFromJson, err := os.ReadFile(jsonName)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytesFromJson, binListFromJson)
	if err != nil {
		return nil, err
	}

	return binListFromJson, err
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
