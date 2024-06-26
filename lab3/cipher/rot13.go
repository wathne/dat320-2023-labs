package cipher

import (
	"io"
)

/*
Task 3: Rot 13

This task is taken from http://tour.golang.org.

A common pattern is an io.Reader that wraps another io.Reader, modifying the
stream in some way.

For example, the gzip.NewReader function takes an io.Reader (a stream of
compressed data) and returns a *gzip.Reader that also implements io.Reader (a
stream of the decompressed data).

Implement a rot13Reader that implements io.Reader and reads from an io.Reader,
modifying the stream by applying the rot13 substitution cipher to all
alphabetical characters.

The rot13Reader type is provided for you. Make it an io.Reader by implementing
its Read method.
*/

type rot13Reader struct {
	r io.Reader
}

var rot13Map map[byte]byte = make(map[byte]byte, 52)
var rot13MapReady bool = false

func rot13MapInit() {
	if rot13MapReady {
		return
	}
	rot13MapInput := []byte(
		"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	rot13MapOutput := []byte(
		"NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm")
	for i, v := range rot13MapInput {
		rot13Map[v] = rot13MapOutput[i]
	}
	rot13MapReady = true
}

func (r rot13Reader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	rot13MapInit()
	for i := 0; i < n; i++ {
		v, ok := rot13Map[p[i]]
		if ok {
			p[i] = v
		}
	}
	return
}
