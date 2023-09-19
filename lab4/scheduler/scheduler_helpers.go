package scheduler

import (
	"dat320/lab4/scheduler/job"
	"dat320/lab4/scheduler/stride"
	"dat320/lab4/scheduler/system"
	"time"
)

// j is a helper function to create a scheduling entry with a job and its arrival time.
// With this helper, the job zero size and the given estimated running time.
var j = func(estimated, arrival time.Duration) *system.Entry {
	return &system.Entry{Job: job.New(0, estimated), Arrival: arrival}
}

var k = func(estimated, arrival time.Duration, tickets int) *system.Entry {
	return &system.Entry{Job: stride.NewJob(0, tickets, estimated), Arrival: arrival}
}

var jc = func(size int, estimated, arrival time.Duration) *system.Entry {
	return &system.Entry{Job: job.New(size, estimated), Arrival: arrival}
}

func GetJobSchedule() []*system.Entry {
	job.ResetJobCounter()
	return []*system.Entry{
		j(t020, 0), // a
		j(t015, 0), // b
	}
}

func GetJobSchedule_Idle() []*system.Entry {
	job.ResetJobCounter()
	return []*system.Entry{
		j(t010, 0),    // a
		j(t010, 0),    // b
		j(t010, t030), // c
		j(t005, t030), // d
		j(t008, t030), // e
	}
}

func GetJobSchedule_ManyJobs() []*system.Entry {
	job.ResetJobCounter()
	return []*system.Entry{
		j(t010, 0),    // a
		j(t003, 0),    // b
		j(t010, 0),    // c
		j(t003, 0),    // d
		j(t010, 0),    // e
		j(t005, 0),    // f
		j(t010, 0),    // g
		j(t008, 0),    // h
		j(t005, t007), // i
		j(t008, t009), // j
	}
}

func GetStrideJobSchedule_2jobs() []*system.Entry {
	job.ResetJobCounter()
	return []*system.Entry{
		k(t020, 0, 100), // a
		k(t010, 0, 50),  // b
	}
}

func GetStrideJobSchedule() []*system.Entry {
	job.ResetJobCounter()
	return []*system.Entry{
		k(t020, 0, 100), // a
		k(t020, 0, 50),  // b
		k(t020, 0, 250), // c
	}
}

func GetStrideJobSchedule_ManyJobs() []*system.Entry {
	job.ResetJobCounter()
	return []*system.Entry{
		k(t020, 0, 100),  // a
		k(t020, 0, 50),   // b
		k(t010, 0, 250),  // c
		k(t020, 0, 10),   // e
		k(t020, 0, 50),   // d
		k(t020, 0, 75),   // f
		k(t020, 0, 100),  // g
		k(t005, 0, 1000), // h
		k(t020, 0, 250),  // i
	}
}
