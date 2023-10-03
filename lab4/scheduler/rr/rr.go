package rr

import (
	"dat320/lab4/scheduler/cpu"
	"dat320/lab4/scheduler/job"
	"dat320/lab4/scheduler/system/systime"
	"time"
)

type roundRobin struct {
	queue job.Jobs
	cpu   *cpu.CPU
	// Add missing fields.
	quantum   time.Duration
	remaining time.Duration
}

func New(cpus []*cpu.CPU, quantum time.Duration) *roundRobin {
	// Construct new RR scheduler.
	if len(cpus) != 1 {
		panic("rr scheduler supports only a single CPU")
	}
	return &roundRobin{
		cpu:       cpus[0],
		quantum:   quantum,
		queue:     make(job.Jobs, 0),
		remaining: time.Duration(0),
	}
}

func (rr *roundRobin) Add(job *job.Job) {
	// Add job to queue.
	rr.queue = append(rr.queue, job)
}

// Tick runs the scheduled jobs for the system time, and returns
// the number of jobs finished in this tick. Depending on scheduler requirements,
// the Tick method may assign new jobs to the CPU before returning.
func (rr *roundRobin) Tick(systemTime time.Duration) int {
	jobsFinished := 0
	// Implement Tick.
	if rr.cpu.IsRunning() {
		if rr.cpu.Tick() {
			jobsFinished++
		}
	}
	if rr.remaining <= time.Duration(0) {
		rr.remaining = rr.quantum
		rr.reassign()
	}
	rr.remaining -= systime.TickDuration
	return jobsFinished
}

// reassign assigns a job to the cpu
func (rr *roundRobin) reassign() {
	// Implement reassign and use it from Tick.
	if rr.cpu.IsRunning() {
		currentJob := rr.cpu.CurrentJob()
		// Put the current job back in the queue if it has remaining time.
		if currentJob.Remaining() > time.Duration(0) {
			rr.Add(currentJob)
		}
	}
	nextJob := rr.getNewJob()
	rr.cpu.Assign(nextJob)
}

// getNewJob finds a new job to run on the CPU, removes the job from the queue and returns the job
func (rr *roundRobin) getNewJob() *job.Job {
	// Implement getNewJob and use it from reassign.
	if len(rr.queue) == 0 {
		return nil
	}
	removedJob := rr.queue[0]
	rr.queue = rr.queue[1:]
	return removedJob
}
