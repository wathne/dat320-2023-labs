package sjf

import (
	"dat320/lab4/scheduler/cpu"
	"dat320/lab4/scheduler/job"
	"sort"
	"time"
)

type sjf struct {
	queue job.Jobs
	cpu   *cpu.CPU
}

func New(cpus []*cpu.CPU) *sjf {
	// Construct new SJF scheduler.
	if len(cpus) != 1 {
		panic("sjf scheduler supports only a single CPU")
	}
	return &sjf{
		cpu:   cpus[0],
		queue: make(job.Jobs, 0),
	}
}

func (s *sjf) Add(job *job.Job) {
	// Add job to queue.
	q := &s.queue
	*q = append(*q, job)
	// Sort queue by least estimated time.
	// The Jobs type implements the sort interface.
	sort.Stable(*q)
	// Sort queue by least estimated time.
	//sort.SliceStable(*q, q.Less)
	// Sort queue by least remaining time.
	/*
	sort.SliceStable(*q, func(i, j int) bool {
		return (*q)[i].Remaining() < (*q)[j].Remaining()
	})
	*/
}

// Tick runs the scheduled jobs for the system time, and returns
// the number of jobs finished in this tick. Depending on scheduler requirements,
// the Tick method may assign new jobs to the CPU before returning.
func (s *sjf) Tick(systemTime time.Duration) int {
	jobsFinished := 0
	// Implement Tick.
	if s.cpu.IsRunning() {
		if s.cpu.Tick() {
			jobsFinished++
			s.reassign()
		}
	} else {
		// CPU is idle, find new job in own queue
		s.reassign()
	}
	return jobsFinished
}

// reassign assigns a job to the cpu
func (s *sjf) reassign() {
	// Implement reassign and use it from Tick.
	nxtJob := s.getNewJob()
	s.cpu.Assign(nxtJob)
}

// getNewJob finds a new job to run on the CPU, removes the job from the queue and returns the job
func (s *sjf) getNewJob() *job.Job {
	// Implement getNewJob and use it from reassign.
	if len(s.queue) == 0 {
		return nil
	}
	removedJob := s.queue[0]
	s.queue = s.queue[1:]
	return removedJob
}
