package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Post struct {
	Title      string
	Url        string
	Id         string
	FaviconUrl string
	Points 	   int
}

type APIResponse struct {
	Hits []Post `json:"hits"`
}

var allPosts []Post = make([]Post, 0)

func NewPost(title, url string, points int) Post {
	// auto generate uuid id
	return Post{
		Title: title, 
		Url: url, 
		Id: uuid.New().String(), 
		FaviconUrl: "https://www.google.com/s2/favicons?domain=" + url,
		Points: points,
	}
}

func main() {

	// create test posts
	allPosts = append(allPosts, NewPost("Post 1", "https://www.google.com", 0))
	allPosts = append(allPosts, NewPost("Post 2", "https://www.facebook.com", 0))

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

	topPosts := getTopPosts()
	fmt.Println(topPosts)

	allPosts = topPosts

	posts := map[string][]Post{
		"Posts": allPosts,
	}

	tmpl.Execute(c.Writer, posts)
}

func submitPostHandler(c *gin.Context) {
	title := c.PostForm("title")
	url := c.PostForm("url")

	post := NewPost(title, url, 0)
	allPosts = append(allPosts, post)

	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.ExecuteTemplate(c.Writer, "post-element", post)
}

func getTopPosts() []Post {
	url := "http://hn.algolia.com/api/v1/search?tags=front_page"
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	var apiResponse APIResponse

	// Decode the JSON response
	err = json.NewDecoder(response.Body).Decode(&apiResponse)
	if err != nil {
		panic(err)
	}

	foundPosts := make([]Post, 0)
	for i := range apiResponse.Hits {	
		newPost := NewPost(apiResponse.Hits[i].Title, apiResponse.Hits[i].Url, apiResponse.Hits[i].Points)
		foundPosts = append(foundPosts, newPost)
	}

	return foundPosts
}
