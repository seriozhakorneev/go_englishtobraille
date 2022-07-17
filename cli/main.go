package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func stringToBraille(s string) (top string, middle string, bottom string) {
	for _, letter := range s {
		// add capitalized letter symbol
		if strings.ToUpper(string(letter)) == string(letter) && unicode.IsLetter(letter) {
			capSlice := strings.SplitN(braille["cap"], " ", -1)
			top += capSlice[0] + capSlice[1] + " "
			middle += capSlice[2] + capSlice[3] + " "
			bottom += capSlice[4] + capSlice[5] + " "
		}

		// making letter lower
		lowerLetter := strings.ToLower(string(letter))
		slice := strings.SplitN(braille[lowerLetter], " ", -1)

		// make lower letter
		top += slice[0] + slice[1] + " "
		middle += slice[2] + slice[3] + " "
		bottom += slice[4] + slice[5] + " "
	}
	return top, middle, bottom
}

// TODO: Refactor this, it can be easier
func printBraille(s string) {
	top, middle, bottom := stringToBraille(s)

	//len of row with 10 Braille symbols(130)
	lengthOfRow := 130
	firstIndex := 0
	for i := 1; i <= len(top)/lengthOfRow; i++ {

		fmt.Println(strings.Trim(top[firstIndex:lengthOfRow], "[]"))
		fmt.Println(strings.Trim(middle[firstIndex:lengthOfRow], "[]"))
		fmt.Println(strings.Trim(bottom[firstIndex:lengthOfRow], "[]"))
		fmt.Printf("\n")

		firstIndex += 130
		lengthOfRow += 130
	}

	if len(top)%130 != 0 {
		fmt.Println(strings.Trim(top[firstIndex:], "[]"))
		fmt.Println(strings.Trim(middle[firstIndex:], "[]"))
		fmt.Println(strings.Trim(bottom[firstIndex:], "[]"))
		fmt.Printf("\n")
	}
}

func cliFlow() {
	for {
		fmt.Print("Enter text (A-Z a-z . , ex1t): ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		if scanner.Text() == "ex1t" {
			fmt.Println("Exiting...")
			break
		}
		fmt.Print("\n")
		printBraille(scanner.Text())
	}
}

func main() {
	fmt.Println("---------English to Braille---------")
	cliFlow()
}
