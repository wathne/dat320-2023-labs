package stack

type stackOperation int

const (
	length stackOperation = iota
	push
	pop
)

type stackCommand struct {
	// Add necessary fields.
	op     stackOperation
	sender chan<- interface{}
	value  interface{}
}

// CspStack is a struct with methods needed to implement the Stack interface.
type CspStack struct {
	// Add necessary fields.
	receiver <-chan stackCommand
	sender   chan<- stackCommand
	size     int
	top      *Element
}

// NewCspStack returns an empty CspStack.
func NewCspStack() *CspStack {
	// Implement constructor and start handling commands.
	var channel chan stackCommand = make(chan stackCommand)
	var receiver <-chan stackCommand = channel
	var sender chan<- stackCommand = channel
	var cs *CspStack = &CspStack{
		receiver: receiver,
		sender:   sender,
		size:     0,
		top:      nil,
	}
	go cs.run()
	return cs
}

// Size returns the size of the stack.
func (cs *CspStack) Size() int {
	// Implement size.
	var channel chan interface{} = make(chan interface{})
	var receiver <-chan interface{} = channel
	var sender chan<- interface{} = channel
	cs.sender <- stackCommand{
		op:     length,
		sender: sender,
		value:  nil,
	}
	return (<-receiver).(int)
}

// Push pushes value onto the stack.
func (cs *CspStack) Push(value interface{}) {
	// Implement push.
	var channel chan interface{} = make(chan interface{})
	var receiver <-chan interface{} = channel
	var sender chan<- interface{} = channel
	cs.sender <- stackCommand{
		op:     push,
		sender: sender,
		value:  value,
	}
	<-receiver
}

// Pop pops the value at the top of the stack and returns it.
func (cs *CspStack) Pop() (value interface{}) {
	// Implement pop.
	var channel chan interface{} = make(chan interface{})
	var receiver <-chan interface{} = channel
	var sender chan<- interface{} = channel
	cs.sender <- stackCommand{
		op:     pop,
		sender: sender,
		value:  nil,
	}
	return <-receiver
}

func (cs *CspStack) run() {
	// Implement handlers for each stack command.
	var cmd stackCommand
	for cmd = range cs.receiver {
		switch cmd.op {
		case length:
			cmd.sender <- cs.size
		case push:
			cs.top = &Element{cmd.value, cs.top}
			cs.size++
			cmd.sender <- nil
		case pop:
			var value interface{}
			if cs.size > 0 {
				value, cs.top = cs.top.value, cs.top.next
				cs.size--
				cmd.sender <- value
			} else {
				cmd.sender <- nil
			}
		}
		close(cmd.sender)
	}
}
