package assertions_test

import (
	"github.com/jspaaks/learn-go-with-tests/generics/pkg/assertions"
	"testing"
)

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {
		assertions.AssertEqual(t, 1, 1)
		assertions.AssertNotEqual(t, 1, 2)
	})
	t.Run("asserting on strings", func(t *testing.T) {
		assertions.AssertEqual(t, "abc", "abc")
		assertions.AssertNotEqual(t, "abc", "def")
	})
	//t.Run("compile error when comparing a strings to an int", func(t *testing.T) {
	//	assertions.AssertEqual(t, "1", 1)
	//})
}
