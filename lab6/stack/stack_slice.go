package stack

// Add necessary fields and synchronization primitives.
import (
	"sync"
)

// DefaultCap is the default stack capacity.
const DefaultCap = 10

// SliceStack is a struct with methods needed to implement the Stack interface.
type SliceStack struct {
	sync.Mutex
	slice []interface{}
	top   int
}

// NewSliceStack returns an empty SliceStack.
func NewSliceStack() *SliceStack {
	return &SliceStack{
		slice: make([]interface{}, DefaultCap),
		top:   -1,
	}
}

// Size returns the size of the stack.
func (ss *SliceStack) Size() int {
	defer ss.Unlock()
	ss.Lock()
	return ss.top + 1
}

// Push pushes value onto the stack.
func (ss *SliceStack) Push(value interface{}) {
	defer ss.Unlock()
	ss.Lock()
	ss.top++
	if ss.top == len(ss.slice) {
		// Reallocate
		newSlice := make([]interface{}, len(ss.slice)*2)
		copy(newSlice, ss.slice)
		ss.slice = newSlice
	}
	ss.slice[ss.top] = value
}

// Pop pops the value at the top of the stack and returns it.
func (ss *SliceStack) Pop() (value interface{}) {
	defer ss.Unlock()
	ss.Lock()
	if ss.top > -1 {
		defer func() { ss.top-- }()
		return ss.slice[ss.top]
	}
	return nil
}
