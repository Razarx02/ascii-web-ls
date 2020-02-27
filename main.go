package main

import (
		"fmt"
		"net/http"
		"log"
		"html/template"

)

var templates *template.Template
type Page struct {
	Output  string
}

func main() {
	templates  = template.Must(template.ParseGlob("*.html"))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	http.HandleFunc("/", homepage)
	http.HandleFunc("/toart/", savestr)
	
	fmt.Println("localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}



func savestr(w http.ResponseWriter, r *http.Request ) {
	if r.URL.Path != "/toart/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	tmpl, err := template.ParseFiles("index.html")
	text := r.FormValue("body")
	shablon := r.FormValue("fs")
	fmt.Println("You used: " + shablon)
	Outputs(text, shablon + ".txt")
	toWrite := Writebanner()
	instring := Bytetostr(toWrite)
	lol := &Page{Output: instring}
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if err := tmpl.Execute(w, lol); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

}

func homepage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
 	templates.ExecuteTemplate(w, "index.html",nil)
}


func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	if status == http.StatusNotFound {
		http.ServeFile(w, r, "404.html")
	} 
}