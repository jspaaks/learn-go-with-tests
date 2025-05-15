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

func BenchmarkRenderPage(b *testing.B) {
	post := blogposts.Post{
		Title:       "This is the post title",
		Description: "This is the post description",
		Tags:        []string{"go", "tdd"},
		Body:        "This is the post body",
	}
	renderer, _ := rendering.NewRenderer()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		renderer.RenderPage(io.Discard, post)
	}
}

func TestRenderPage(t *testing.T) {
	t.Run("rendering a page", func(t *testing.T) {
		post := blogposts.Post{
			Title:       "This is the post title",
			Description: "This is the post description",
			Tags:        []string{"go", "tdd"},
			Body:        "This is the post body",
		}
		renderer, err1 := rendering.NewRenderer()
		if err1 != nil {
			t.Fatal(err1)
		}
		buf := bytes.Buffer{}
		err2 := renderer.RenderPage(&buf, post)
		assertNoErrors(t, err2)
		approvals.VerifyString(t, buf.String())
	})
}
