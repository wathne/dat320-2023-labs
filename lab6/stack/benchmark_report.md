# Benchmark Report

## CPU and Memory Benchmarks for all Three Stacks

```console
$ go test -v -run none -bench Benchmark -memprofilerate=1 -benchmem
goos: darwin
goarch: arm64
pkg: dat320/lab6/stack
BenchmarkStacks
BenchmarkStacks/SafeStack
BenchmarkStacks/SafeStack-8                   76      15545400 ns/op      317956 B/op      19744 allocs/op
BenchmarkStacks/SliceStack
BenchmarkStacks/SliceStack-8                 223       5327767 ns/op       77952 B/op       9744 allocs/op
BenchmarkStacks/CspStack
BenchmarkStacks/CspStack-8                    21      57510647 ns/op     2238002 B/op      39744 allocs/op
PASS
ok      dat320/lab6/stack   5.193s
```

1. How much faster than the slowest is the fastest stack?
    - [ ] a) 2x-3x
    - [ ] b) 3x-4x
    - [x] c) 5x-20x
    - [ ] d) 20x-30x

2. Which stack requires the most allocated memory?
    - [x] a) CspStack
    - [ ] b) SliceStack
    - [ ] c) SafeStack
    - [ ] d) UnsafeStack

3. Which stack requires the least amount of allocated memory?
    - [ ] a) CspStack
    - [x] b) SliceStack
    - [ ] c) SafeStack
    - [ ] d) UnsafeStack

## Memory Profile of BenchmarkStacks/SafeStack

```console
$ go test -v -run none -bench BenchmarkStacks/SafeStack -memprofile=safe-stack.prof
goos: darwin
goarch: arm64
pkg: dat320/lab6/stack
BenchmarkStacks
BenchmarkStacks/SafeStack
BenchmarkStacks/SafeStack-8                 2854        417633 ns/op
PASS
ok      dat320/lab6/stack   2.251s
$ go tool pprof safe-stack.prof
Type: alloc_space
Time: Oct 27, 2023 at 8:03pm (CEST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 1569.53MB, 100% of 1569.53MB total
Dropped 1 node (cum <= 7.85MB)
      flat  flat%   sum%        cum   cum%
 1164.53MB 74.20% 74.20%  1164.53MB 74.20%  dat320/lab6/stack.(*SafeStack).Push
  405.01MB 25.80%   100%  1569.53MB   100%  dat320/lab6/stack.BenchmarkStacks.func1
         0     0%   100%  1568.53MB 99.94%  testing.(*B).launch
         0     0%   100%  1569.53MB   100%  testing.(*B).runN
(pprof) list Push
Total: 1.53GB
ROUTINE ======================== dat320/lab6/stack.(*SafeStack).Push in /Users/joar/Documents.nosync/uis/dat320-2023/workdir/lab6/stack/stack_sync.go
    1.14GB     1.14GB (flat, cum) 74.20% of Total
         .          .     23:func (ss *SafeStack) Push(value interface{}) {
         .          .     24:   defer ss.mutex.Unlock()
         .          .     25:   ss.mutex.Lock()
    1.14GB     1.14GB     26:   ss.top = &Element{value, ss.top}
         .          .     27:   ss.size++
         .          .     28:}
         .          .     29:
         .          .     30:// Pop pops the value at the top of the stack and returns it.
         .          .     31:func (ss *SafeStack) Pop() (value interface{}) {
```

4. Which function accounts for all memory allocations in the `SafeStack` implementation?
    - [ ] a) `Size`
    - [ ] b) `NewSafeStack`
    - [x] c) `Push`
    - [ ] d) `Pop`

5. Which line in `SafeStack` does the actual memory allocation?
    - [ ] a) `type SafeStack struct {`
    - [x] b) `ss.top = &Element{value, ss.top}`
    - [ ] c) `value, ss.top = ss.top.value, ss.top.next`
    - [ ] d) `top  *Element`

## CPU Profile of BenchmarkStacks/CspStack

```console
$ go test -v -run none -bench BenchmarkStacks/CspStack -cpuprofile=csp-stack.prof
goos: darwin
goarch: arm64
pkg: dat320/lab6/stack
BenchmarkStacks
BenchmarkStacks/CspStack
BenchmarkStacks/CspStack-8               156       7523796 ns/op
PASS
ok      dat320/lab6/stack   2.133s
$ go tool pprof csp-stack.prof
Type: cpu
Time: Oct 27, 2023 at 8:02pm (CEST)
Duration: 1.90s, Total samples = 2.13s (111.96%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top38
Showing nodes accounting for 2.10s, 98.59% of 2.13s total
Dropped 23 nodes (cum <= 0.01s)
      flat  flat%   sum%        cum   cum%
     0.95s 44.60% 44.60%      0.95s 44.60%  runtime.pthread_cond_signal
     0.66s 30.99% 75.59%      0.66s 30.99%  runtime.usleep
     0.43s 20.19% 95.77%      0.43s 20.19%  runtime.pthread_cond_wait
     0.03s  1.41% 97.18%      0.03s  1.41%  runtime.madvise
     0.02s  0.94% 98.12%      0.02s  0.94%  runtime.pthread_kill
     0.01s  0.47% 98.59%      0.02s  0.94%  runtime.gcDrain
         0     0% 98.59%      0.03s  1.41%  runtime.(*mheap).alloc.func1
         0     0% 98.59%      0.03s  1.41%  runtime.(*mheap).allocSpan
         0     0% 98.59%      0.91s 42.72%  runtime.findRunnable
         0     0% 98.59%      0.02s  0.94%  runtime.forEachP
         0     0% 98.59%      0.02s  0.94%  runtime.gcBgMarkWorker.func2
         0     0% 98.59%      0.02s  0.94%  runtime.gcMarkDone.func1
         0     0% 98.59%      0.31s 14.55%  runtime.lock (inline)
         0     0% 98.59%      0.31s 14.55%  runtime.lock2
         0     0% 98.59%      0.31s 14.55%  runtime.lockWithRank (inline)
         0     0% 98.59%      0.43s 20.19%  runtime.mPark (inline)
         0     0% 98.59%      0.92s 43.19%  runtime.mcall
         0     0% 98.59%      0.43s 20.19%  runtime.notesleep
         0     0% 98.59%      0.95s 44.60%  runtime.notewakeup
         0     0% 98.59%      0.31s 14.55%  runtime.osyield (inline)
         0     0% 98.59%      0.92s 43.19%  runtime.park_m
         0     0% 98.59%      0.02s  0.94%  runtime.preemptM
         0     0% 98.59%      1.12s 52.58%  runtime.ready
         0     0% 98.59%      0.24s 11.27%  runtime.recv.goready.func1
         0     0% 98.59%      0.35s 16.43%  runtime.runqgrab
         0     0% 98.59%      0.35s 16.43%  runtime.runqsteal
         0     0% 98.59%      0.92s 43.19%  runtime.schedule
         0     0% 98.59%      0.44s 20.66%  runtime.semasleep
         0     0% 98.59%      0.95s 44.60%  runtime.semawakeup
         0     0% 98.59%      0.88s 41.31%  runtime.send.goready.func1
         0     0% 98.59%      0.02s  0.94%  runtime.signalM (inline)
         0     0% 98.59%      0.94s 44.13%  runtime.startm
         0     0% 98.59%      0.35s 16.43%  runtime.stealWork
         0     0% 98.59%      0.43s 20.19%  runtime.stopm
         0     0% 98.59%      0.03s  1.41%  runtime.sysUsed (inline)
         0     0% 98.59%      0.03s  1.41%  runtime.sysUsedOS (inline)
         0     0% 98.59%      1.19s 55.87%  runtime.systemstack
         0     0% 98.59%      1.13s 53.05%  runtime.wakep
```
