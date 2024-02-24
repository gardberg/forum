package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Post struct {
	Title   string
	Content string
	Url     string
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))

	posts := map[string][]Post{
		"Posts": {
			{Title: "Post 1", Content: "Content 1", Url: "https://www.google.com"},
			{Title: "Post 2", Content: "Content 2", Url: "https://www.google.com"},
		},
	}

	tmpl.Execute(w, posts)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func submitPostHandler(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	content := r.FormValue("content")
	url := r.FormValue("url")

	// Print form values to the server logs
	fmt.Println("Received form values:")
	fmt.Println("Title:", title)
	fmt.Println("Content:", content)
	fmt.Println("URL:", url)

	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.ExecuteTemplate(w, "post-element", Post{Title: title, Content: content, Url: url})

}

func main() {
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/submit-post/", submitPostHandler)
	fmt.Println("Server is running on http://localhost:8000")
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", nil))
}
