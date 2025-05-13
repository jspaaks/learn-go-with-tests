package blogposts_test

import (
	"errors"
	"io/fs"
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
		fs := fstest.MapFS{
			"hello_world.md": {
				Data: []byte("hi"),
			},
			"hello_world2.md": {
				Data: []byte("hola"),
			},
		}
		posts, err := blogposts.NewPostsFromFilesystem(fs)
		assertNoErrors(t, err)
		assertCorrectNumberOfPosts(t, posts, fs)
	})
}
