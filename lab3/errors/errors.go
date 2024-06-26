package errors

import (
	"fmt"
)

/*
Task 5: Errors needed for multiwriter

You may find this blog post useful:
http://blog.golang.org/error-handling-and-go

Similar to a the Stringer interface, the error interface also defines a
method that returns a string.

// Note(wathne): Added extra comments. The linter would misinterpret this part.
// type error interface {
//     Error() string
// }

Thus also the error type can describe itself as a string. The fmt package (and
many others) use this Error() method to print errors.

Implement the Error() method for the Errors type defined above.

The following conditions should be covered:

1. When there are no errors in the slice, it should return:

"(0 errors)"

2. When there is one error in the slice, it should return:

The error string return by the corresponding Error() method.

3. When there are two errors in the slice, it should return:

The first error + " (and 1 other error)"

4. When there are X>1 errors in the slice, it should return:

The first error + " (and X other errors)"
*/
func (m Errors) Error() string {
	if m == nil {
		return "(0 errors)"
	}
	validErrors := make([]error, 0)
	for _, v := range m {
		if v == nil {
			continue
		}
		validErrors = append(validErrors, v)
	}
	errorCount := len(validErrors)
	if errorCount == 1 {
		return validErrors[0].Error()
	}
	if errorCount == 2 {
		return fmt.Sprint(validErrors[0], " (and 1 other error)")
	}
	if errorCount > 2 {
		return fmt.Sprint(validErrors[0], " (and ", errorCount-1,
			" other errors)")
	}
	return "(0 errors)"
}
