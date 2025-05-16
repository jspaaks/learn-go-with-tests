package stack_test

import (
	"github.com/jspaaks/learn-go-with-tests/generics/pkg/assertions"
	"github.com/jspaaks/learn-go-with-tests/generics/pkg/stack"
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("stack", func(t *testing.T) {
		stack := stack.NewStack[int]()

		// check stack is empty
		assertions.AssertTrue(t, stack.IsEmpty())

		// add a thing, then check it's not empty
		stack.Push(123)
		assertions.AssertFalse(t, stack.IsEmpty())

		// add another thing, pop it back again
		stack.Push(456)
		value, _ := stack.Pop()
		assertions.AssertEqual(t, value, 456)
		value, _ = stack.Pop()
		assertions.AssertEqual(t, value, 123)
		assertions.AssertTrue(t, stack.IsEmpty())

		// can get the numbers we put in as numbers, not untyped interface{}
		stack.Push(1)
		stack.Push(2)
		firstNum, _ := stack.Pop()
		secondNum, _ := stack.Pop()
		assertions.AssertEqual(t, firstNum+secondNum, 3)
	})
}
