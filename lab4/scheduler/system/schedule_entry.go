package system

import (
	"dat320/lab4/scheduler/job"
	"time"
)

type Schedule []*Entry

type Entry struct {
	Job     *job.Job
	Arrival time.Duration
}
