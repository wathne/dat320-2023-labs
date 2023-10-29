package water

import (
	"sync"
)

// Water structure holds the synchronization primitives and
// data required to solve the water molecule problem.
// moleculeCount holds the number of molecules formed so far.
// result string contains the sequence of "H" and "O".
// wg WaitGroup is used to wait for goroutine completion.
type Water struct {
	wg            sync.WaitGroup
	moleculeCount int
	result        string
	// Add missing fields, if necessary.
	done     chan bool
	hydrogen chan string
	oxygen   chan string
}

// New initializes the water structure.
func New() *Water {
	water := &Water{}
	// Initialize the Water struct.
	water.done = make(chan bool)
	water.hydrogen = make(chan string)
	water.oxygen = make(chan string)
	return water
}

// releaseOxygen produces one oxygen atom if no oxygen atom is already present.
// If an oxygen atom is already present, it will block until enough hydrogen
// atoms have been produced to consume the atoms necessary to produce water.
//
// The w.wg.Done() must be called to indicate the completion of the goroutine.
func (w *Water) releaseOxygen() {
	defer w.wg.Done()
	// Implement the releaseOxygen routine.
	w.oxygen <- "O"
}

// releaseHydrogen produces one hydrogen atom unless two hydrogen atoms are already present.
// If two hydrogen atoms are already present, it will block until another oxygen
// atom has been produced to consume the atoms necessary to produce water.
//
// The w.wg.Done() must be called to indicate the completion of the goroutine.
func (w *Water) releaseHydrogen() {
	defer w.wg.Done()
	// Implement the releaseHydrogen routine.
	w.hydrogen <- "H"
}

// produceMolecule forms the water molecules.
func (w *Water) produceMolecule(done chan bool) {
	// Implement the produceMolecule routine.
	var result *string = &w.result
	var atom string
production:
	for {
		select {
		case atom = <-w.hydrogen:
			*result += atom
			select {
			case atom = <-w.hydrogen:
				*result += atom
				*result += <-w.oxygen
			case atom = <-w.oxygen:
				*result += atom
				*result += <-w.hydrogen
			}
		case atom = <-w.oxygen:
			*result += atom
			*result += <-w.hydrogen
			*result += <-w.hydrogen
		case <-w.done:
			break production
		}
		w.moleculeCount++
	}
	done <- true
}

func (w *Water) finish() {
	// Implement the finish routine to complete the water molecule formation.
	w.done <- true
}

// Molecules returns the number of water molecules that has been created.
func (w *Water) Molecules() int {
	// Add any missing code.
	return w.moleculeCount
}

// Make returns a sequence of water molecules derived from the input of hydrogen and oxygen atoms.
// DO NOT edit the Make method or modify the signatures of the other methods used by Make.
func (w *Water) Make(input string) string {
	done := make(chan bool)
	go w.produceMolecule(done)
	for _, ch := range input {
		w.wg.Add(1)
		switch ch {
		case 'O':
			go w.releaseOxygen()
		case 'H':
			go w.releaseHydrogen()
		}
	}
	w.wg.Wait()
	w.finish()
	<-done
	return w.result
}
