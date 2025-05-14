package blogrenderer_test

import (
	"bytes"
	"github.com/jspaaks/learn-go-with-tests/templating/pkg/blogposts"
	"github.com/jspaaks/learn-go-with-tests/templating/pkg/blogrenderer"
	"github.com/approvals/go-approval-tests"
	"testing"
)

func assertNoErrors(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("Didn't expect an error but got one: %s", err)
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
		assertNoErrors(t, err)
		approvals.VerifyString(t, buf.String())
	})
}
