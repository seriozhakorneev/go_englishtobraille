package main

import (
	"log"
	"net/http"
	"strings"
	"text/template"
	"unicode"
)

var tpl *template.Template

func init() { tpl = template.Must(template.ParseGlob("templates/*gohtml")) }

var brailleMap = map[string]string{
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

func stringToBraille(s string) (brailleString []string) {
	for _, letter := range s {
		// add capitalized letter symbol
		if strings.ToUpper(string(letter)) == string(letter) && unicode.IsLetter(letter) {
			brailleString = append(brailleString, brailleMap["cap"])
		}
		// making letter lower
		lowerLetter := strings.ToLower(string(letter))
		brailleString = append(brailleString, brailleMap[lowerLetter])
	}
	return
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

func index(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("text")
	brailleOutput := stringToBraille(text)

	err := tpl.ExecuteTemplate(
		w,
		"index.gohtml",
		map[string][]string{
			"signs": brailleOutput,
		},
	)
	if err != nil {
		log.Fatalln("Failed to execute template: ", err)
	}
}

func getText(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	inputText := r.FormValue("input_s")

	if textValidator(inputText) {
		http.Redirect(w, r, "/?text="+inputText, http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
	return
}

func main() {
	port := ":8080"
	http.HandleFunc("/", index)
	http.HandleFunc("/getText", getText)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Server starts on", "http://localhost"+port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("failed to ListenAndServe: ", err)
	}
}
