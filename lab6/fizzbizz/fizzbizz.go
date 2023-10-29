package fizzbizz

import (
	"dat320/lab6/wathne" // Logging.
	"strconv"
	"sync"
)

// SyncBlock structure holds the synchronization constructs and
// data required to solve the fizzbizz problem.
// current contains the current value to be processed.
// max is the value upto which the goroutines should process.
// result holds the output of the assignment.
// wg is the WaitGroup used to wait for the completion of the goroutines.
type SyncBlock struct {
	sync.Mutex
	cond    *sync.Cond
	wg      sync.WaitGroup
	current int
	max     int
	result  string
	printer chan<- string // Logging.
}

// newSyncBlock initializes the SyncBlock.
func newSyncBlock(max int) *SyncBlock {
	block := &SyncBlock{}
	block.max = max
	block.current = 1
	block.cond = sync.NewCond(block)
	// Initialize a safe buffered printer.
	block.printer = wathne.NewPrinter(32) // Logging.
	return block
}

// appendToResult appends partialResult to the result
// generated by the goroutines and increments s.current.
// Must only be called when holding the lock on s.
func (s *SyncBlock) appendToResult(partialResult string) {
	s.result = s.result + partialResult
	s.current++
}

// fizz appends "Fizz" to the result if
// s.current is divisible by 3 and not divisible by 5,
// and increments the value of s.current.
//
// Otherwise, if the above condition is not satisfied,
// fizz waits for other goroutines to execute and
// increment the value of s.current.
func (s *SyncBlock) fizz() {
	defer s.wg.Done()
	// Implement the fizz routine.
	// Wait for available mutex. Lock mutex on resume.
	s.cond.L.Lock()
	for {
		s.printer <- wathne.LogFizzBizzRoutine(s.current) // Logging.
		if (s.current > s.max) {
			s.printer <- wathne.LogFizzBizzBreak() // Logging.
			s.cond.Broadcast()
			s.cond.L.Unlock()
			break
		}
		if (s.current % 3 == 0) && (s.current % 5 != 0) {
			s.appendToResult("Fizz")
			s.printer <- wathne.LogFizzBizzTrue(s.result) // Logging.
		} else {
			s.printer <- wathne.LogFizzBizzFalse(s.result) // Logging.
		}
		s.cond.Broadcast()
		// Unlock mutex. Wait for broadcast. Wait for available mutex.
		// Lock mutex on resume.
		s.cond.Wait()
	}
	s.printer <- wathne.LogFizzBizzDone() // Logging.
}

// bizz appends "Bizz" to the result if
// s.current is divisible by 5 and not divisible by 3,
// and increments the value of s.current.
//
// Otherwise, if the above condition is not satisfied,
// bizz waits for other goroutines to execute and
// increment the value of s.current.
func (s *SyncBlock) bizz() {
	defer s.wg.Done()
	// Implement the bizz routine.
	// Wait for available mutex. Lock mutex on resume.
	s.cond.L.Lock()
	for {
		s.printer <- wathne.LogFizzBizzRoutine(s.current) // Logging.
		if (s.current > s.max) {
			s.printer <- wathne.LogFizzBizzBreak() // Logging.
			s.cond.Broadcast()
			s.cond.L.Unlock()
			break
		}
		if (s.current % 5 == 0) && (s.current % 3 != 0) {
			s.appendToResult("Bizz")
			s.printer <- wathne.LogFizzBizzTrue(s.result) // Logging.
		} else {
			s.printer <- wathne.LogFizzBizzFalse(s.result) // Logging.
		}
		s.cond.Broadcast()
		// Unlock mutex. Wait for broadcast. Wait for available mutex.
		// Lock mutex on resume.
		s.cond.Wait()
	}
	s.printer <- wathne.LogFizzBizzDone() // Logging.
}

// number appends s.current (as a string) to the result if
// s.current is not divisible by 3 and 5,
// and increments the value of s.current.
//
// Otherwise, if the above condition is not satisfied,
// number waits for other goroutines to execute and
// increment the value of s.current.
func (s *SyncBlock) number() {
	defer s.wg.Done()
	// Implement the number routine.
	// Wait for available mutex. Lock mutex on resume.
	s.cond.L.Lock()
	for {
		s.printer <- wathne.LogFizzBizzRoutine(s.current) // Logging.
		if (s.current > s.max) {
			s.printer <- wathne.LogFizzBizzBreak() // Logging.
			s.cond.Broadcast()
			s.cond.L.Unlock()
			break
		}
		if (s.current % 5 != 0) && (s.current % 3 != 0) {
			s.appendToResult(strconv.Itoa(s.current))
			s.printer <- wathne.LogFizzBizzTrue(s.result) // Logging.
		} else {
			s.printer <- wathne.LogFizzBizzFalse(s.result) // Logging.
		}
		s.cond.Broadcast()
		// Unlock mutex. Wait for broadcast. Wait for available mutex.
		// Lock mutex on resume.
		s.cond.Wait()
	}
	s.printer <- wathne.LogFizzBizzDone() // Logging.
}

// fizzBizz appends "FizzBizz" to the result if
// s.current is divisible by 3 and 5,
// and increments the value of s.current.
//
// Otherwise, if the above condition is not satisfied,
// fizzBizz waits for other goroutines to execute and
// increment the value of s.current.
func (s *SyncBlock) fizzBizz() {
	defer s.wg.Done()
	// Implement the fizzbizz routine.
	// Wait for available mutex. Lock mutex on resume.
	s.cond.L.Lock()
	for {
		s.printer <- wathne.LogFizzBizzRoutine(s.current) // Logging.
		if (s.current > s.max) {
			s.printer <- wathne.LogFizzBizzBreak() // Logging.
			s.cond.Broadcast()
			s.cond.L.Unlock()
			break
		}
		if (s.current % 5 == 0) && (s.current % 3 == 0) {
			s.appendToResult("FizzBizz")
			s.printer <- wathne.LogFizzBizzTrue(s.result) // Logging.
		} else {
			s.printer <- wathne.LogFizzBizzFalse(s.result) // Logging.
		}
		s.cond.Broadcast()
		// Unlock mutex. Wait for broadcast. Wait for available mutex.
		// Lock mutex on resume.
		s.cond.Wait()
	}
	s.printer <- wathne.LogFizzBizzDone() // Logging.
}

// FizzBizz returns the result of the fizzbizz algorithm for the given max.
// The output is produced when all goroutines have completed.
// DO NOT edit the FizzBizz function or modify the signatures of the other methods used by FizzBizz.
func FizzBizz(max int) string {
	const numGoroutines = 4
	s := newSyncBlock(max)
	s.wg.Add(numGoroutines)
	go s.fizz()
	go s.bizz()
	go s.number()
	go s.fizzBizz()
	s.wg.Wait()
	return s.result
}
