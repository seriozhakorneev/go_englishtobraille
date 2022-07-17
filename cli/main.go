package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type braillePrint struct {
	top, middle, bottom string
}

func (b braillePrint) String() string {
	//len of row with 10 Braille symbols(130)
	lengthOfRow := 130
	firstIndex := 0
	for i := 1; i <= len(b.top)/lengthOfRow; i++ {
		fmt.Println(strings.Trim(b.top[firstIndex:lengthOfRow], "[]"))
		fmt.Println(strings.Trim(b.middle[firstIndex:lengthOfRow], "[]"))
		fmt.Println(strings.Trim(b.bottom[firstIndex:lengthOfRow], "[]"))
		fmt.Printf("\n")
		firstIndex += 130
		lengthOfRow += 130
	}

	if len(b.top)%130 != 0 {
		fmt.Println(strings.Trim(b.top[firstIndex:], "[]"))
		fmt.Println(strings.Trim(b.middle[firstIndex:], "[]"))
		fmt.Println(strings.Trim(b.bottom[firstIndex:], "[]"))
		fmt.Printf("\n")
	}
	return ""
}

func stringToBraillePrint(s string) (top string, middle string, bottom string) {
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

func textValidator(str string) bool {
	for _, charVariable := range str {
		if (charVariable < 'a' || charVariable > 'z') &&
			(charVariable < 'A' || charVariable > 'Z') &&
			(charVariable != ' ') {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("---------English to Braille---------")
	for {
		fmt.Print("Enter text (A-Z a-z . , ex1t): ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		if scanner.Text() == "ex1t" {
			fmt.Println("Exiting...")
			break
		}
		fmt.Print("\n")
		b := braillePrint{}
		s := scanner.Text()
		if textValidator(s) {
			// TODO:
			// эти символы принтуются в консоль как хотят особенно если строка больше 42
			// где-то разъехалось, надо чинить
			b.top, b.middle, b.bottom = stringToBraillePrint(s)
			fmt.Println(b)
		}
	}
}
