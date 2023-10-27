package stack

// Add necessary fields and synchronization primitives.
import (
	"sync"
)

// SafeStack holds the top element of the stack and its size.
type SafeStack struct {
	mutex sync.Mutex
	size  int
	top   *Element
}

// Size returns the size of the stack.
func (ss *SafeStack) Size() int {
	defer ss.mutex.Unlock()
	ss.mutex.Lock()
	return ss.size
}

// Push pushes value onto the stack.
func (ss *SafeStack) Push(value interface{}) {
	defer ss.mutex.Unlock()
	ss.mutex.Lock()
	ss.top = &Element{value, ss.top}
	ss.size++
}

// Pop pops the value at the top of the stack and returns it.
func (ss *SafeStack) Pop() (value interface{}) {
	defer ss.mutex.Unlock()
	ss.mutex.Lock()
	if ss.size > 0 {
		value, ss.top = ss.top.value, ss.top.next
		ss.size--
		return
	}
	return nil
}
