package main

import (
	"net/http"
	"html/template"
)

var hogeTmpl = template.Must(template.New("hoge").ParseFiles("./templates/base.html", "./templates/hoge.html"))

func hogeHandler(w http.ResponseWriter, r *http.Request) {
	hogeTmpl.ExecuteTemplate(w, "base", "Hoge")
}

var piyoTmpl = template.Must(template.New("piyo").ParseFiles("./templates/base.html", "./templates/piyo.html"))

func piyoHandler(w http.ResponseWriter, r *http.Request) {
	piyoTmpl.ExecuteTemplate(w, "base", "Piyo")
}

func main() {
	// hoge
	http.HandleFunc("/", hogeHandler)
	http.HandleFunc("/hoge", hogeHandler)

	// piyo
	http.HandleFunc("/piyo", piyoHandler)

	// css
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))

	http.ListenAndServe(":8000", nil)
}