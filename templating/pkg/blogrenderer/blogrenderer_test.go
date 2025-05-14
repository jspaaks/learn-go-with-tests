package blogrenderer_test

import (
	"bytes"
	"github.com/jspaaks/learn-go-with-tests/templating/pkg/blogposts"
	"github.com/jspaaks/learn-go-with-tests/templating/pkg/blogrenderer"
	"testing"
)

func assertNoErrors(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("Didn't expect an error but got one: %s", err)
	}
}
func assertCorrectHtml(t *testing.T, actual string, expected string) {
	t.Helper()
	if actual != expected {
		t.Fatalf("actual HTML string %q didn't match expected HTML string %q", actual, expected)
	}
}

func TestBlogRenderer(t *testing.T) {
	t.Run("rendering a Post struct", func(t *testing.T) {
		post := blogposts.Post{
			Title:       "This is the post title",
			Description: "This is the post description",
			Tags:        []string{"go", "tdd"},
			Body:        "This is the post body",
		}
		buf := bytes.Buffer{}
		err := blogrenderer.Render(&buf, post)
		expected := 
`<h1>This is the post title</h1>
<p>This is the post description</p>
Tags:
<ul>
    <li>go</li>
    <li>tdd</li>
</ul>
`
		assertNoErrors(t, err)
		assertCorrectHtml(t, buf.String(), expected)
	})
}
