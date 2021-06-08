package main

import (
	"net/http"
	"strings"
	"text/template"
	"unicode"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*gohtml"))
}

type Braille struct {
	Blist []string
}

func stringToBraille(s string) *Braille {

	brailleMap := map[string]string{
		"a": "a.png", "b": "b.png", "c": "c.png",
		"d": "d.png", "e": "e.png", "f": "f.png",
		"g": "g.png", "h": "h.png", "i": "i.png",
		"j": "j.png", "k": "k.png", "l": "l.png",
		"m": "m.png", "n": "n.png", "o": "o.png",
		"p": "p.png", "q": "q.png", "r": "r.png",
		"s": "s.png", "t": "t.png", "u": "u.png",
		"v": "v.png", "w": "w.png", "x": "x.png",
		"y": "y.png", "z": "z.png",
		" ": "space.png", "cap": "cap.png",
	}

	brailleString := Braille{[]string{}}

	for _, letter := range s {
		// add capitalized letter symbol
		if strings.ToUpper(string(letter)) == string(letter) && unicode.IsLetter(letter) {
			brailleString.Blist = append(brailleString.Blist, brailleMap["cap"])
		}
		// making letter lower
		lower_letter := strings.ToLower(string(letter))
		brailleString.Blist = append(brailleString.Blist, brailleMap[lower_letter])
	}
	return &brailleString
}

func TextValidator(str string) bool {
	for _, charVariable := range str {
		if (charVariable < 'a' || charVariable > 'z') &&
			(charVariable < 'A' || charVariable > 'Z') &&
			(charVariable != ' ') {
			return false
		}
	}
	return true
}

func index(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("text")
	braille_output := stringToBraille(text)
	tpl.ExecuteTemplate(w, "index.gohtml", braille_output)
}

func getText(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	inputText := r.FormValue("input_s")

	if TextValidator(inputText) {
		http.Redirect(w, r, "/?text="+inputText, http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}

func handleRequest() {
	http.HandleFunc("/", index)
	http.HandleFunc("/getText", getText)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()
}
