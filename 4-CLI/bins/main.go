package bins

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

type BinList = []Bin

func NewBin() Bin {
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

	return Bin{Id: id, Private: private, CreatedAt: time.Now(), Name: name}
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

func NewBinList(lenOfBinList, capOfBinList int) (BinList, error) {
	if capOfBinList < 0 {
		return make([]Bin, 0), errors.New("negative capacity")
	} else if lenOfBinList < 0 {
		return make([]Bin, 0), errors.New("negative lenth")
	} else if lenOfBinList < capOfBinList {
		return make([]Bin, 0), errors.New("length is less than capacity")
	}

	newBinList := make([]Bin, lenOfBinList, capOfBinList)

	for i := 0; i < lenOfBinList; i++ {
		newBinList[i] = NewBin()
	}

	return BinList(newBinList), nil
}
