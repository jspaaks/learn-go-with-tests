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

// Don't instantiate with Renderer{}, use NewRenderer() instead.
type Renderer struct {
	templ *template.Template
}

// Renders the post as HTML using the embedded templates.
func (r *Renderer) RenderPost(writer io.Writer, post blogposts.Post) error {
	if err := r.templ.ExecuteTemplate(writer, "template.index.gohtml", post); err != nil {
		return err
	}
	return nil
}

//Parses the embedded HTML templates and returns a renderer.
func NewRenderer() (*Renderer, error) {
	parsed, err := template.ParseFS(templates, "template.*.gohtml")
	if err != nil {
		return nil, err
	}
	return &Renderer{templ: parsed}, nil
}
