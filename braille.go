package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func stringToBraille(s string) (string, string, string) {

	braille := map[string]string{
		"a": "⚫️ ⚪️ ⚪️ ⚪️ ⚪️ ⚪️", "b": "⚫️ ⚪️ ⚫️ ⚪️ ⚪️ ⚪️", "c": "⚫️ ⚫️ ⚪️ ⚪️ ⚪️ ⚪️",
		"d": "⚫️ ⚫️ ⚪️ ⚫️ ⚪️ ⚪️", "e": "⚫️ ⚪️ ⚪️ ⚫️ ⚪️ ⚪️", "f": "⚫️ ⚫️ ⚫️ ⚪️ ⚪️ ⚪️",
		"g": "⚫️ ⚫️ ⚫️ ⚫️ ⚪️ ⚪️", "h": "⚫️ ⚪️ ⚫️ ⚫️ ⚪️ ⚪️", "i": "⚪️ ⚫️ ⚫️ ⚪️ ⚪️ ⚪️",
		"j": "⚪️ ⚫️ ⚫️ ⚫️ ⚪️ ⚪️", "k": "⚫️ ⚪️ ⚪️ ⚪️ ⚫️ ⚪️", "l": "⚫️ ⚪️ ⚫️ ⚪️ ⚫️ ⚪️",
		"m": "⚫️ ⚫️ ⚪️ ⚪️ ⚫️ ⚪️", "n": "⚫️ ⚫️ ⚪️ ⚫️ ⚫️ ⚪️", "o": "⚫️ ⚪️ ⚪️ ⚫️ ⚫️ ⚪️",
		"p": "⚫️ ⚫️ ⚫️ ⚪️ ⚫️ ⚪️", "q": "⚫️ ⚫️ ⚫️ ⚫️ ⚫️ ⚪️", "r": "⚫️ ⚪️ ⚫️ ⚫️ ⚫️ ⚪️",
		"s": "⚪️ ⚫️ ⚫️ ⚪️ ⚫️ ⚪️", "t": "⚪️ ⚫️ ⚫️ ⚫️ ⚫️ ⚪️", "u": "⚫️ ⚪️ ⚪️ ⚪️ ⚫️ ⚫️",
		"v": "⚫️ ⚪️ ⚫️ ⚪️ ⚫️ ⚫️", "w": "⚪️ ⚫️ ⚫️ ⚫️ ⚪️ ⚫️", "x": "⚫️ ⚫️ ⚪️ ⚪️ ⚫️ ⚫️",
		"y": "⚫️ ⚫️ ⚪️ ⚫️ ⚫️ ⚫️", "z": "⚫️ ⚪️ ⚪️ ⚫️ ⚫️ ⚫️", " ": "⚪️ ⚪️ ⚪️ ⚪️ ⚪️ ⚪️",
		".": "⚪️ ⚪️ ⚫️ ⚫️ ⚪️ ⚫️", ",": "⚪️ ⚪️ ⚫️ ⚪️ ⚪️ ⚪️", "cap": "⚪️ ⚪️ ⚪️ ⚪️ ⚪️ ⚫️",
	}

	top := ""
	middle := ""
	bottom := ""
	for _, letter := range s {
		// add capitalized letter symbol
		if strings.ToUpper(string(letter)) == string(letter) && unicode.IsLetter(letter) {

			cap_slice := strings.SplitN(braille["cap"], " ", -1)
			top += cap_slice[0] + cap_slice[1] + " "
			middle += cap_slice[2] + cap_slice[3] + " "
			bottom += cap_slice[4] + cap_slice[5] + " "
		}

		// making letter lower
		lower_letter := strings.ToLower(string(letter))
		slice := strings.SplitN(braille[lower_letter], " ", -1)

		top += slice[0] + slice[1] + " "
		middle += slice[2] + slice[3] + " "
		bottom += slice[4] + slice[5] + " "

	}
	return top, middle, bottom
}

func printBraille(top, middle, bottom string) {
	// len of row with 10 Braille symbols(130)
	length_of_row := 130
	loop_count := len(top) / length_of_row
	first_index := 0

	for i := 1; i <= loop_count; i++ {

		fmt.Println(strings.Trim(top[first_index:length_of_row], "[]"))
		fmt.Println(strings.Trim(middle[first_index:length_of_row], "[]"))
		fmt.Println(strings.Trim(bottom[first_index:length_of_row], "[]"))
		fmt.Printf("\n")

		first_index += 130
		length_of_row += 130
	}

	if len(top)%130 != 0 {

		fmt.Println(strings.Trim(top[first_index:], "[]"))
		fmt.Println(strings.Trim(middle[first_index:], "[]"))
		fmt.Println(strings.Trim(bottom[first_index:], "[]"))
		fmt.Printf("\n")
	}
}

func main() {
	fmt.Println("English to Braille.")

	for {
		fmt.Print("Enter text (A-Z a-z . , ex1t): ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		if scanner.Text() == "ex1t" {
			break
		}
		t, m, b := stringToBraille(scanner.Text())

		fmt.Print("\n")
		printBraille(t, m, b)
	}
}
