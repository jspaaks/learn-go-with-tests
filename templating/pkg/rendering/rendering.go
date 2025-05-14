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

type renderer struct {
	templ *template.Template
}

func (r *renderer) RenderPost(writer io.Writer, post blogposts.Post) error {
	if err := r.templ.ExecuteTemplate(writer, "template.index.gohtml", post); err != nil {
		return err
	}
	return nil
}

func NewRenderer() (*renderer, error) {
	parsed, err := template.ParseFS(templates, "template.*.gohtml")
	if err != nil {
		return nil, err
	}
	return &renderer{templ: parsed}, nil
}
