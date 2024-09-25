package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

// types
type Story map[string]Chapter

type Handler struct {
	story Story
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == "/" {
		path = "/intro"
	}
	chapter, ok := h.story[path[1:]]
	if !ok {
		http.NotFound(w, r)
	} else {
		tml := template.Must(template.ParseFiles("static/html/template.html"))
		err := tml.Execute(w, chapter)
		check(err)
	}
}

type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func loadStories(filename string) Story {
	var Stories Story
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(file, &Stories)
	if err != nil {
		log.Fatal(err)
	}
	return Stories
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	stories := loadStories("gopher.json")
	http.ListenAndServe(":8080", Handler{stories})
	fmt.Println("Listening on :8080")

}
