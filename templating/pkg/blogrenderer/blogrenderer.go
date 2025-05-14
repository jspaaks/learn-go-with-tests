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
	templ, err := template.ParseFS(templates, "template.*.gohtml")
	if err != nil {
		return err
	}
	if err := templ.Execute(writer, post); err != nil {
		return err
	}
	return nil
}
