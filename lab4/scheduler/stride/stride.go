package stride

import (
	"dat320/lab4/scheduler/cpu"
	"dat320/lab4/scheduler/job"
	"dat320/lab4/scheduler/system/systime"
	"time"
)

type stride struct {
	queue job.Jobs
	cpu   *cpu.CPU
	// Add missing fields.
	quantum time.Duration
	remaining time.Duration
}

func New(cpus []*cpu.CPU, quantum time.Duration) *stride {
	// Construct new stride scheduler.
	if len(cpus) != 1 {
		panic("stride scheduler supports only a single CPU")
	}
	return &stride{
		cpu: cpus[0],
		quantum: quantum,
		queue: make(job.Jobs, 0),
		remaining: time.Duration(0),
	}
}

func (s *stride) Add(job *job.Job) {
	// Add job to queue.
	s.queue = append(s.queue, job)
}

// Tick runs the scheduled jobs for the system time, and returns
// the number of jobs finished in this tick. Depending on scheduler requirements,
// the Tick method may assign new jobs to the CPU before returning.
func (s *stride) Tick(systemTime time.Duration) int {
	jobsFinished := 0
	// Implement Tick().
	if s.cpu.IsRunning() {
		if s.cpu.Tick() {
			jobsFinished++
		}
	}
	if s.remaining <= time.Duration(0) {
		s.remaining = s.quantum
		s.reassign()
	}
	s.remaining -= systime.TickDuration
	return jobsFinished
}

// reassign assigns a job to the cpu
func (s *stride) reassign() {
	// Implement reassign and use it from Tick().
	if s.cpu.IsRunning() {
		currentJob := s.cpu.CurrentJob()
		// Put the current job back in the queue if it has remaining time.
		if currentJob.Remaining() > time.Duration(0) {
			s.Add(currentJob)
		}
	}
	nextJob := s.getNewJob()
	s.cpu.Assign(nextJob)
}

// getNewJob finds a new job to run on the CPU, removes the job from the queue and returns the job
func (s *stride) getNewJob() *job.Job {
	// Implement getNewJob and use it from reassign.
	if len(s.queue) == 0 {
		return nil
	}
	index := MinPass(s.queue)
	removedJob := s.queue[index]
	removedJob.Pass += removedJob.Stride
	s.queue = append(s.queue[:index], s.queue[index+1:]...)
	return removedJob
}

// minPass returns the index of the job with the lowest pass value.
func MinPass(theJobs job.Jobs) int {
	lowest := 0
	// Implement MinPass and use it from getNewJob.
	for i, job := range theJobs {
		lowestJob := theJobs[lowest]
		if job.Pass < lowestJob.Pass {
			lowest = i
		}
		// If multiple pass values are the same, use the job with the lowest stride value instead.
		if job.Pass == lowestJob.Pass {
			if job.Stride < lowestJob.Stride {
				lowest = i
			}
		}
	}
	return lowest
}
