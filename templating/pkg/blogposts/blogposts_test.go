package blogposts_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	"github.com/jspaaks/learn-go-with-tests/reading-files/pkg/blogposts"
)

type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, i always fail")
}

func TestNewPostsFromFilesystem(t *testing.T) {
	t.Run("failure", func(t *testing.T) {
		assertHasError := func(t *testing.T, err error) {
			t.Helper()
			if err == nil {
				t.Fatal("Unexpectedly didn't get an error while trying to read from mocked filesystem.")
			}
		}
		assertPostsNil := func(t *testing.T, posts []blogposts.Post) {
			t.Helper()
			if posts != nil {
				t.Fatal("Expected posts to be nil")
			}
		}
		filesystem := StubFailingFS{}
		posts, err := blogposts.NewPostsFromFilesystem(filesystem)
		assertHasError(t, err)
		assertPostsNil(t, posts)
	})
	t.Run("success", func(t *testing.T) {
		assertNoErrors := func(t *testing.T, err error) {
			t.Helper()
			if err != nil {
				t.Fatal("Unexpectedly got an error while trying to read from mocked filesystem.")
			}
		}
		assertCorrectNumberOfPosts := func(t *testing.T, posts []blogposts.Post, fs fstest.MapFS) {
			t.Helper()
			actual := len(posts)
			expected := len(fs)
			if actual != expected {
				t.Fatalf("got %d posts instead of the expected %d", actual, expected)
			}
		}
		assertPostHasCorrectContent := func(t *testing.T, actual blogposts.Post, expected blogposts.Post) {
			t.Helper()
			if !reflect.DeepEqual(actual, expected) {
				t.Errorf("got %+v, want %+v", actual, expected)
			}
		}
		const first = "Title: Post 1\nDescription: Description 1\nTags: tdd, go\n---\nHello\nWorld"
		const second = "Title: Post 2\nDescription: Description 2\nTags: rust, borrow-checker\n---\nB\nL\nM"
		fs := fstest.MapFS{
			"post1.md": {
				Data: []byte(first),
			},
			"post2.md": {
				Data: []byte(second),
			},
		}
		posts, err := blogposts.NewPostsFromFilesystem(fs)
		expected := []blogposts.Post{
			{
				Title:       "Post 1",
				Description: "Description 1",
				Tags:        []string{"tdd", "go"},
				Body:        "Hello\nWorld",
			},
			{
				Title:       "Post 2",
				Description: "Description 2",
				Tags:        []string{"rust", "borrow-checker"},
				Body:        "B\nL\nM",
			},
		}
		assertNoErrors(t, err)
		assertCorrectNumberOfPosts(t, posts, fs)
		assertPostHasCorrectContent(t, posts[0], expected[0])
	})
}
