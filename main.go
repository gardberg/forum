package main

import (
	"html/template"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Post struct {
	Title string
	Url   string
	Id    string
}

var allPosts []Post = make([]Post, 0)

func NewPost(title, url string) Post {
	// auto generate uuid id
	return Post{Title: title, Url: url, Id: uuid.New().String()}
}

func main() {

	// create test posts
	allPosts = append(allPosts, NewPost("Post 1", "https://www.google.com"))
	allPosts = append(allPosts, NewPost("Post 2", "https://www.facebook.com"))

	// start server
	r := gin.Default()

	r.GET("/", defaultHandler)
	r.POST("/submit-post", submitPostHandler)

	if err := r.Run("127.0.0.1:8000"); err != nil {
		panic(err)
	}
}

func defaultHandler(c *gin.Context) {
	tmpl := template.Must(template.ParseFiles("index.html"))

	posts := map[string][]Post{
		"Posts": allPosts,
	}

	tmpl.Execute(c.Writer, posts)
}


func submitPostHandler(c *gin.Context) {
	title := c.PostForm("title")
	url := c.PostForm("url")

	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.ExecuteTemplate(c.Writer, "post-element", NewPost(title, url))
}
