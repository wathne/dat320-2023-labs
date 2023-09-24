package job

import (
	"dat320/lab4/scheduler/system/systime"
	"time"
)

func (j *Job) Scheduled(s systime.SystemTime) {
	// How can job.Job call Now()? Rhetorical question.
	// systime.SystemTime interface is embedded in job.Job struct.
	// systime.SystemTime interface is implemented by system.System struct.
	// system.System.Run() calls job.Job.Scheduled() and pass a system.System
	// (itself) as the argument. See below, as we set the embedded interface.
	j.SystemTime = s
	// Implemented task 2.1
	j.arrival = j.Now()
}

func (j *Job) Started(cpuID int) {
	// Implemented task 2.2
	if j.start == NotStartedYet {
		j.start = j.Now()
	}
}

func (j Job) TurnaroundTime() time.Duration {
	// Implemented task 2.3
	tt := j.finished - j.arrival
	return tt
}

func (j Job) ResponseTime() time.Duration {
	// Implemented task 2.3
	rt := j.start - j.arrival
	return rt
}
