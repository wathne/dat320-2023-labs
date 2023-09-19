package system

import (
	"dat320/lab4/scheduler/job"
	"time"
)

// Avg returns the average of a metric defined by the function f.
func (sch Schedule) Avg(f func(*job.Job) time.Duration) time.Duration {
	// TODO implement task 2.4(a)
	sum := time.Duration(0)
	return sum
}

func (sch Schedule) AvgResponseTime() time.Duration {
	// TODO implement task 2.4(b)
	sum := time.Duration(0)
	return sum
}

func (sch Schedule) AvgTurnaroundTime() time.Duration {
	// TODO implement task 2.4(b)
	sum := time.Duration(0)
	return sum
}
