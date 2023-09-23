package job

import (
	"dat320/lab4/scheduler/system/systime"
	"time"
)

func (j *Job) Scheduled(s systime.SystemTime) {
	j.SystemTime = s
	// TODO(student) implement task 2.1
}

func (j *Job) Started(cpuID int) {
	// TODO(student) implement task 2.2
}

func (j Job) TurnaroundTime() time.Duration {
	// TODO(student) implement task 2.3
	return 0
}

func (j Job) ResponseTime() time.Duration {
	// TODO(student) implement task 2.3
	return 0
}
