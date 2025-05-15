package rendering

import (
	"embed"
	"github.com/jspaaks/learn-go-with-tests/templating/pkg/blogposts"
	"html/template"
	"io"
)

var (
	//go:embed "template.*.gohtml"
	templates embed.FS
)

// Abstract data type.
type renderer struct {
	templ *template.Template
}

// Interface for renderers.
type Renderer interface {
	RenderListOfPosts(writer io.Writer, posts []blogposts.Post) error
	RenderPost(writer io.Writer, post blogposts.Post) error
}

// Renders the page as HTML using the embedded templates and the posts data.
func (r *renderer) RenderListOfPosts(writer io.Writer, posts []blogposts.Post) error {
	if err := r.templ.ExecuteTemplate(writer, "template.index.gohtml", posts); err != nil {
		return err
	}
	return nil
}

// Renders the page as HTML using the embedded templates and the data for a single data.
func (r *renderer) RenderPost(writer io.Writer, post blogposts.Post) error {
	if err := r.templ.ExecuteTemplate(writer, "template.post.gohtml", post); err != nil {
		return err
	}
	return nil
}

// Parses the embedded HTML templates and returns a renderer.
func NewRenderer() (Renderer, error) {
	parsed, err := template.ParseFS(templates, "template.*.gohtml")
	if err != nil {
		return nil, err
	}
	return &renderer{templ: parsed}, nil
}
