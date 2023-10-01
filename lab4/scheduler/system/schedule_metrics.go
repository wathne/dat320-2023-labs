package system

import (
	"dat320/lab4/scheduler/job"
	"math"
	"time"
)

// Avg returns the average of a metric defined by the function f.
func (sch Schedule) Avg(f func(*job.Job) time.Duration) time.Duration {
	// Implemented task 2.4(a)
	sum := time.Duration(0)
	var i int
	var entry *Entry
	for i, entry = range sch {
		sum += f(entry.Job)
	}
	return time.Duration(math.Round(float64(sum) / float64(i+1)))
}

func (sch Schedule) AvgResponseTime() time.Duration {
	// Implemented task 2.4(b)
	return sch.Avg(func(j *job.Job) time.Duration {
		return j.ResponseTime()
	})
}

func (sch Schedule) AvgTurnaroundTime() time.Duration {
	// Implemented task 2.4(b)
	return sch.Avg(func(j *job.Job) time.Duration {
		return j.TurnaroundTime()
	})
}
