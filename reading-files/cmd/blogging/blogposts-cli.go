package main

import (
	"github.com/jspaaks/learn-go-with-tests/reading-files/pkg/blogposts"
	"log"
	"os"
)

func main() {
	posts, err := blogposts.NewPostsFromFilesystem(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}
	for _, post := range posts {
		log.Printf("%+v", post)
	}
}
