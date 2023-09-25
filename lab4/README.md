# Lab 4: Scheduling and Metrics

| Lab 4: | Scheduling and Metrics |
| ---------------------    | --------------------- |
| Subject:                 | DAT320 Operating Systems and Systems Programming |
| Deadline:                | **October 7, 2023 23:59** |
| Expected effort:         | 20-30 hours |
| Grading:                 | Pass/fail |
| Submission:              | Group Presentation |

## Table of Contents

- [Lab 4: Scheduling and Metrics](#lab-4-scheduling-and-metrics)
  - [Table of Contents](#table-of-contents)
  - [Introduction](#introduction)
  - [Scheduling](#scheduling)
  - [Description of the Scheduling Simulator](#description-of-the-scheduling-simulator)
    - [Task 1: Implement Scheduling Algorithms](#task-1-implement-scheduling-algorithms)
      - [Testing the Various Schedulers](#testing-the-various-schedulers)
    - [Task 2: Implement Scheduling Metrics for Scheduled Jobs](#task-2-implement-scheduling-metrics-for-scheduled-jobs)
    - [Expected Output for Task 1 Schedulers](#expected-output-for-task-1-schedulers)

## Introduction

In this lab, you will build a job scheduler able to schedule jobs according to different scheduling policies.

<span style="color:red">Please note that in order to obtain approval for this lab, as well as all subsequent group labs, it is necessary to deliver your code, in group, to one of the student assistants.</span>

## Scheduling

From the lectures you have learned about scheduling.
In this exercise, you will implement three different job schedulers in Go.

| Policy                    | Description                                                                                                                        |
| ------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| First In First Out (FIFO) | Schedules jobs in the order of arrival.                                                                                            |
| Shortest Job First (SJF)  | Schedules jobs based on the estimated execution time; runs the shortest jobs first.                                                |
| Round Robin (RR)          | Schedules jobs in the FIFO order, but only for some given time quantum, giving each job a fair share of the processor.             |
| Stride Scheduling (SS)    | Schedules jobs using the Stride Scheduling algorithm in Chapter 9.6, giving each job an exact proportional share of the processor. |

We provide the FIFO scheduler as a template for implementing the more advanced schedulers.
Each scheduler should be able to schedule a list of jobs according to the different scheduling policies given in the table above.

The Round Robin algorithm will need to keep track of the remaining time for each job, since it might not have completed, when its time slice is exhausted.
Note that for Round Robin, job switching happens only at synchronized time intervals, determined by the system time and the `quantum` parameter.
This means that if a job finishes before the time quantum expires, the CPU running that job should become **Idle**.

Stride Scheduling works by giving jobs a *pass* and *stride* value based on each job's allocated *tickets*.
Similar to Round Robin, jobs are scheduled for some given time quantum.
If multiple pass values are the same, use the job with the lowest *stride* value instead.
Note that this is different from the textbook, which say that then the choice is arbitrary.
However, arbitrary choices are not conducive to unit tests.

## Description of the Scheduling Simulator

The assignments build on a *simulated* scheduling framework.
The simulator consists of separate packages for different entities, such as the CPU, Job, and the System itself.
The system receives a schedule of job entries via its `Run` method, and is responsible for scheduling these jobs according to their arrival times.
The system can be given any type of scheduler as long as it implements the `system.Scheduler` interface.

```golang
type Scheduler interface {
    Add(*job.Job)
    Tick(time.Duration) int
}
```

The `Add` method takes a job and adds it to the scheduler's queue.

- The `Tick` method is where the main scheduling logic should be implemented.
- The `Tick` method is called by the simulator every clock tick of the system, defined via the `systime.TickDuration` constant.
- The `Tick` method receives the current *system time* as input.
  That is, the number of ticks since the start of the simulation.
- The `Tick` method should tick all CPUs that has a running job, unless the scheduler specifies otherwise.
- The `Tick` method returns the number of jobs that finished in the current tick.
  That is, `Tick` may return 0 or 1 for a single-CPU system, and in a multi-CPU system it may return a value between 0 and the number of CPUs whose job finished in the current clock tick.
- The `Tick` method *may assign* new jobs to the different CPUs before returning.

The `System` for driving the scheduling will only use the `Scheduler` interface (the `Add` and `Tick` methods) to interact with a scheduler.

### Task 1: Implement Scheduling Algorithms

The code is organized in several files.
Study and familiarize yourself with the code.

The different schedulers that you implement must satisfy the `Scheduler` interface in `system/scheduler_iface.go`.
Schedulers may need additional inputs and fields in their respective structs, such as the time slice, or `quantum`, as used by RR and SS.

The `Scheduler`'s `Tick()` method is responsible for ticking the system forward.
`Tick()` should return the number of jobs finished at each tick and may reassign a new job to the cpu if the `quantum` is over.

The `Scheduler`'s `reassign()` method is responsible for correctly assigning a new job to the cpu.
We suggest using `GetNewJob()` to find the next job in the queue to assign in the `reassign()` method.

The `Tick()` method of the `cpu` struct is responsible for actually executing the job.

Finally, since the stride scheduler works with `tickets`, `stride`, and `pass` values per `job`, you will have to implement the `stride.NewJob()` constructor function that returns a `*job.Job` with the relevant stride fields set.

#### Testing the Various Schedulers

You can do preliminary testing of your schedulers locally before pushing to GitHub and QuickFeed for testing.
The test will only print out what is on the cpu at each timestep.
Therefore you must examine the output from the tests and compare to the expected output [Expected Output](#expected-output-for-task-1-schedulers)
That is, to run the various tests in the `system_test.go` file, first cd into the `scheduler` directory:

```console
cd lab4/scheduler
```

Then use one of these commands:

```console
go test -v -run TestFIFOScheduler
go test -v -run TestShortestJobFirstScheduler
go test -v -run TestRoundRobinScheduler
go test -v -run TestStrideScheduler
```

You can add the `-verbose` flag to see the scheduler's decisions, e.g.:

```console
go test -v -run TestFIFOScheduler -verbose
```

Note that before you can test the stride scheduler you must implement the `stride.NewJob()` constructor function.
You can test your function with:

```console
cd lab4/scheduler/stride
go test -v -run TestStrideNewJob
```

We also provide a test for the `MinPass()` function that belongs to the stride scheduler.
To only run tests for `MinPass()` use this command:

```console
go test -v -run TestMinPass
```

If you want to run all tests in one go, use this command:

```console
go test -v ./...
```

The `-v` flag is used to print verbose output, that is, print a message for each test that passes or fails.
Without the `-v` flag, the command will only print when a test fails.
The `./...` part tells the `go` command to run all tests in all subdirectories.

### Task 2: Implement Scheduling Metrics for Scheduled Jobs

In this assignment, you will implement functions for computing two useful scheduling metrics.
For reference, Chapter 7 in the textbook describes the scheduling metrics.

All functions **must** be implemented in the provided files `job/job_metrics.go` and `system/schedule_metrics.go`.
You may wish to study the code in `job/job.go`.
Observe that you may access the current (simulated) system time using the `Now()` method available on the `Job` type.

If you cannot get the tests to pass, you will still get points for the correct parts.

1. Implement the function `Scheduled`.

2. Implement the function `Started`.
   Note that this function may be called multiple times during an execution.
   However, you should only update the start time of the job, when first started.
   You may find the `NotStartedYet` constant useful.

3. Implement the functions `TurnaroundTime` and `ResponseTime` to compute the turnaround time and response time of a single job, respectively.

   You can run the following command to check if implemented correctly.

   ```shell
   cd system
   go test -v -run TestSingleJobMetrics/fifo/book_schedule
   ```

   You can also run with the `verbose` flag to show the schedule execution.

   ```shell
   go test -v -run TestSingleJobMetrics/fifo/book_schedule -verbose
   ```

   *Once you have accomplished this subtask, make sure to commit your changes locally.*

4. Implement:

   (a) A single `Avg` function to compute the average turnaround time and average response time of a schedule:

   ```golang
   func (sch Schedule) Avg(f func(*job.Job) time.Duration) time.Duration
   ```

   (b) Two separate functions `AvgTurnaroundTime` and `AvgResponseTime` to compute the average turnaround time and average response time of a schedule:

   ```golang
   func (sch Schedule) AvgTurnaroundTime() time.Duration
   func (sch Schedule) AvgResponseTime() time.Duration
   ```

   **Hint:** Each job's turnaround time and response time should *only be included once* in the average, even though the job appears many times in the schedule.

   You can run the following command to check if implemented correctly.

   ```shell
   go test -v -run TestAverageMetrics/fifo/book_schedule
   ```

   *Once you have accomplished this subtask, make sure to commit your changes locally.*

   NOTE: For this task, the tests expect the results to be rounded up.
   Thus you need to round the values before computing the `time.Duration` to be returned.
   Also, note that `math.Round` expects a `float64` type, so you may need to convert the types.

5. What can you say about the impact of the time quantum on the two schedulers?

   (You should be prepared to discuss this with your approver.)

### Expected Output for Task 1 Schedulers

The number in first parenthesis after the job name is the remaining time for the job.
The `1x` in the second parenthesis is the CPUs speed and is always 1 in this assignment.

```text
=== RUN   TestFIFOScheduler
Tick  CPU0
0s    A( 3ms)(1x)
1ms   A( 2ms)(1x)
2ms   A( 1ms)(1x)
3ms   B( 3ms)(1x)
4ms   B( 2ms)(1x)
5ms   B( 1ms)(1x)
6ms   C( 4ms)(1x)
7ms   C( 3ms)(1x)
8ms   C( 2ms)(1x)
9ms   C( 1ms)(1x)
10ms  D( 2ms)(1x)
11ms  D( 1ms)(1x)
12ms  E( 4ms)(1x)
13ms  E( 3ms)(1x)
14ms  E( 2ms)(1x)
15ms  E( 1ms)(1x)
16ms  F( 4ms)(1x)
17ms  F( 3ms)(1x)
18ms  F( 2ms)(1x)
19ms  F( 1ms)(1x)
20ms  G( 5ms)(1x)
21ms  G( 4ms)(1x)
22ms  G( 3ms)(1x)
23ms  G( 2ms)(1x)
24ms  G( 1ms)(1x)
25ms  H( 3ms)(1x)
26ms  H( 2ms)(1x)
27ms  H( 1ms)(1x)
28ms  Idle
```

```text
=== RUN   TestShortestJobFirstScheduler
Tick  CPU0
0s    D( 2ms)(1x)
1ms   D( 1ms)(1x)
2ms   A( 3ms)(1x)
3ms   A( 2ms)(1x)
4ms   A( 1ms)(1x)
5ms   B( 3ms)(1x)
6ms   B( 2ms)(1x)
7ms   B( 1ms)(1x)
8ms   C( 4ms)(1x)
9ms   C( 3ms)(1x)
10ms  C( 2ms)(1x)
11ms  C( 1ms)(1x)
12ms  H( 3ms)(1x)
13ms  H( 2ms)(1x)
14ms  H( 1ms)(1x)
15ms  E( 4ms)(1x)
16ms  E( 3ms)(1x)
17ms  E( 2ms)(1x)
18ms  E( 1ms)(1x)
19ms  F( 4ms)(1x)
20ms  F( 3ms)(1x)
21ms  F( 2ms)(1x)
22ms  F( 1ms)(1x)
23ms  G( 5ms)(1x)
24ms  G( 4ms)(1x)
25ms  G( 3ms)(1x)
26ms  G( 2ms)(1x)
27ms  G( 1ms)(1x)
28ms  Idle
```

```text
=== RUN   TestRoundRobinScheduler
Tick  CPU0
0s    A( 2ms)(1x)
1ms   A( 1ms)(1x)
2ms   B( 3ms)(1x)
3ms   B( 2ms)(1x)
4ms   C( 4ms)(1x)
5ms   C( 3ms)(1x)
6ms   D( 2ms)(1x)
7ms   D( 1ms)(1x)
8ms   E( 4ms)(1x)
9ms   E( 3ms)(1x)
10ms  F( 4ms)(1x)
11ms  F( 3ms)(1x)
12ms  B( 1ms)(1x)
13ms  Idle
14ms  C( 2ms)(1x)
15ms  C( 1ms)(1x)
16ms  G( 5ms)(1x)
17ms  G( 4ms)(1x)
18ms  H( 3ms)(1x)
19ms  H( 2ms)(1x)
20ms  E( 2ms)(1x)
21ms  E( 1ms)(1x)
22ms  F( 2ms)(1x)
23ms  F( 1ms)(1x)
24ms  G( 3ms)(1x)
25ms  G( 2ms)(1x)
26ms  H( 1ms)(1x)
27ms  Idle
28ms  G( 1ms)(1x)
29ms  Idle
```

```text
=== RUN   TestStrideScheduler
Tick  CPU0
0s    C(10ms)(1x)
1ms   A(10ms)(1x)
2ms   B(10ms)(1x)
3ms   C( 9ms)(1x)
4ms   C( 8ms)(1x)
5ms   A( 9ms)(1x)
6ms   C( 7ms)(1x)
7ms   C( 6ms)(1x)
8ms   C( 5ms)(1x)
9ms   A( 8ms)(1x)
10ms  B( 9ms)(1x)
11ms  C( 4ms)(1x)
12ms  C( 3ms)(1x)
13ms  A( 7ms)(1x)
14ms  C( 2ms)(1x)
15ms  C( 1ms)(1x)
16ms  A( 6ms)(1x)
17ms  B( 8ms)(1x)
18ms  A( 5ms)(1x)
19ms  A( 4ms)(1x)
20ms  B( 7ms)(1x)
21ms  A( 3ms)(1x)
22ms  A( 2ms)(1x)
23ms  B( 6ms)(1x)
24ms  A( 1ms)(1x)
25ms  B( 5ms)(1x)
26ms  B( 4ms)(1x)
27ms  B( 3ms)(1x)
28ms  B( 2ms)(1x)
29ms  B( 1ms)(1x)
30ms  Idle
```
