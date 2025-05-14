package blogrenderer

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

func Render(writer io.Writer, post blogposts.Post) error {
	parsed, err := template.ParseFS(templates, "template.*.gohtml")
	if err != nil {
		return err
	}
	if err := parsed.ExecuteTemplate(writer, "template.index.gohtml", post); err != nil {
		return err
	}
	return nil
}
