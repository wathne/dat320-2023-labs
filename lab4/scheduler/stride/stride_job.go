package stride

import (
	"dat320/lab4/scheduler/job"
	"time"
)

// NewJob creates a job for stride scheduling.
func NewJob(size, tickets int, estimated time.Duration) *job.Job {
	const numerator = 10_000
	job := job.New(size, estimated)
	job.Stride = int(numerator / tickets)
	job.Pass = int(0)
	job.Tickets = tickets
	return job
}
