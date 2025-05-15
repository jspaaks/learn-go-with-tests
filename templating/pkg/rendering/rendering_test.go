package rendering_test

import (
	"bytes"
	"github.com/approvals/go-approval-tests"
	"github.com/jspaaks/learn-go-with-tests/templating/pkg/blogposts"
	"github.com/jspaaks/learn-go-with-tests/templating/pkg/rendering"
	"io"
	"testing"
)

func assertNoErrors(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("Didn't expect an error but got one: %s", err)
	}
}

func BenchmarkRenderPost(b *testing.B) {
	post := blogposts.Post{
		Title:       "This is the post title",
		Description: "This is the post description",
		Tags:        []string{"go", "tdd"},
		Body:        "This is the post body",
	}
	renderer, _ := rendering.NewRenderer()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		renderer.RenderPost(io.Discard, post)
	}
}

func TestRendering(t *testing.T) {
	posts := []blogposts.Post{
		{
			Title:       "This is the first post title",
			Description: "This is the first post description",
			Tags:        []string{"go", "tdd"},
			Body:        "This is the first post body",
		},
		{
			Title:       "This is the second post title",
			Description: "This is the second post description",
			Tags:        []string{"rust", "borrow-checker"},
			Body:        "This is the second post body",
		},
	}
	renderer, err := rendering.NewRenderer()
	if err != nil {
		t.Fatal(err)
	}
	t.Run("rendering a single post page", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := renderer.RenderPost(&buf, posts[0])
		assertNoErrors(t, err)
		approvals.VerifyString(t, buf.String())
	})
	t.Run("rendering a list of posts page", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := renderer.RenderListOfPosts(&buf, posts)
		assertNoErrors(t, err)
		approvals.VerifyString(t, buf.String())
	})
}
