package stride

import (
	"dat320/lab4/scheduler/cpu"
	"dat320/lab4/scheduler/job"
	"time"
)

type stride struct {
	queue job.Jobs
	cpu   *cpu.CPU
	// TODO(student) add missing fields, if necessary
}

func New(cpus []*cpu.CPU, quantum time.Duration) *stride {
	// TODO(student) construct new stride scheduler
	return nil
}

func (s *stride) Add(job *job.Job) {
	// TODO(student) Add job to queue
}

// Tick runs the scheduled jobs for the system time, and returns
// the number of jobs finished in this tick. Depending on scheduler requirements,
// the Tick method may assign new jobs to the CPU before returning.
func (s *stride) Tick(systemTime time.Duration) int {
	jobsFinished := 0
	// TODO(student) Implement Tick()
	return jobsFinished
}

// reassign assigns a job to the cpu
func (s *stride) reassign() {
	// TODO(student) Implement reassign and use it from Tick()
}

// getNewJob finds a new job to run on the CPU, removes the job from the queue and returns the job
func (s *stride) getNewJob() *job.Job {
	// TODO(student) Implement getNewJob and use it from reassign
	return nil
}

// minPass returns the index of the job with the lowest pass value.
func MinPass(theJobs job.Jobs) int {
	lowest := 0
	// TODO(student) Implement MinPass and use it from getNewJob
	return lowest
}
