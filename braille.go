package main

import (
	"fmt"
	"strings"
	"unicode"
)

// braille var should be here
//var x string = `this is

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
	length_of_row := 130 // len of row with 10 Braille symbols(130)
	loop_count := len(top) / length_of_row
	first_index := 0

	for i := 1; i <= loop_count; i++ {
		//println(i)
		fmt.Println(strings.Trim(top[first_index:length_of_row], "[]"))
		fmt.Println(strings.Trim(middle[first_index:length_of_row], "[]"))
		fmt.Println(strings.Trim(bottom[first_index:length_of_row], "[]"))
		fmt.Printf("\n")

		first_index += 130
		length_of_row += 130

	}
	// запустить если не четное
	fmt.Println(strings.Trim(top[first_index:], "[]"))
	fmt.Println(strings.Trim(middle[first_index:], "[]"))
	fmt.Println(strings.Trim(bottom[first_index:], "[]"))
	fmt.Printf("\n")

}

func main() {

	var s_input string

	fmt.Print("Enter your name: ")
	fmt.Scanf("%s", &s_input)
	fmt.Println(s_input)

	//s := "The Quick Brown Fox, Jumps Over A Ly Dog."

	t, m, b := stringToBraille(s_input)
	printBraille(t, m, b)

}
