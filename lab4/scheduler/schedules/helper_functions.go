package schedules

import (
	"dat320/lab4/scheduler/job"
	"time"
)

// cpuJobMap maps a CPU's ID to a job.
type cpuJobMap = map[int]job.Job

// JobMap maps a scheduled duration for a job to run on a CPU;
// the CPU's ID is the key to the secondary map that holds the job as value.
type JobMap map[time.Duration]cpuJobMap

var jb = func(id int, estimated, remaining time.Duration) job.Job {
	return job.NewTestJob(id, estimated, remaining)
}
