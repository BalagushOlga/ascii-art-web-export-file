package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", generateAscii)
	http.HandleFunc("/export/", exportAscii)

	fmt.Println("Server listening on port. Go to http://localhost:8080")
	Check(http.ListenAndServe(":8080", nil))
}

func generateAscii(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("page.html"))
	outputStr := "here will be ascii-art"
	if r.URL.Path == "/ascii-art" {
		inputString := r.FormValue("inputString")
		file := r.FormValue("font")
		outputStr = asciiArt(inputString, file, w)
	}
	if r.URL.Path != "/ascii-art" && r.URL.Path != "/" {
		outputStr = "Bad request"
		w.WriteHeader(http.StatusBadRequest)
	}
	tmpl.Execute(w, outputStr)
}

func exportAscii(w http.ResponseWriter, r *http.Request) {
	str := strings.Join(strings.Split(r.URL.Path[8:], "~")[1:], "\n")
	str = strings.ReplaceAll(str, "±", "\\")

	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.Header().Add("Content-Disposition", "attachment; filename=result.txt")
	w.Header().Add("Content-Length", strconv.Itoa(len(str)))
	fmt.Fprintf(w, str)
}

func Check(e error) {
	if e != nil {
		fmt.Println("\nERROR 500 - Internal server error")
	}
}
