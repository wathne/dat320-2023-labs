package scheduler

import (
	"dat320/lab4/scheduler/cpu"
	"dat320/lab4/scheduler/fifo"
	"dat320/lab4/scheduler/job"
	"dat320/lab4/scheduler/rr"
	"dat320/lab4/scheduler/schedules"
	"dat320/lab4/scheduler/sjf"
	"dat320/lab4/scheduler/stride"
	"dat320/lab4/scheduler/system"
	"testing"
)

func TestFIFOScheduler(t *testing.T) {
	job.ResetJobCounter()
	schedule := []*system.Entry{
		j(t003, 0),    // a
		j(t003, 0),    // b
		j(t004, 0),    // c
		j(t002, 0),    // d
		j(t004, 0),    // e
		j(t004, 0),    // f
		j(t005, t007), // g
		j(t003, t009), // h
	}
	cpus := cpu.NewCPUs(numCPUs)
	scheduler := fifo.New(cpus)
	if scheduler == nil {
		t.Fatalf("fifo scheduler not implemented")
	}
	sys := system.New(scheduler, cpus)
	sys.Run(schedule)
}

func TestShortestJobFirstScheduler(t *testing.T) {
	job.ResetJobCounter()
	schedule := []*system.Entry{
		j(t003, 0),    // a
		j(t003, 0),    // b
		j(t004, 0),    // c
		j(t002, 0),    // d
		j(t004, 0),    // e
		j(t004, 0),    // f
		j(t005, t007), // g
		j(t003, t009), // h
	}
	cpus := cpu.NewCPUs(numCPUs)
	scheduler := sjf.New(cpus)
	if scheduler == nil {
		t.Fatalf("sjf scheduler not implemented")
	}
	sys := system.New(scheduler, cpus)
	sys.Run(schedule)
}

func TestRoundRobinScheduler(t *testing.T) {
	job.ResetJobCounter()
	schedule := []*system.Entry{
		j(t002, 0),    // a
		j(t003, 0),    // b
		j(t004, 0),    // c
		j(t002, 0),    // d
		j(t004, 0),    // e
		j(t004, 0),    // f
		j(t005, t007), // g
		j(t003, t009), // h
	}
	cpus := cpu.NewCPUs(numCPUs)
	scheduler := rr.New(cpus, t002)
	if scheduler == nil {
		t.Fatalf("rr scheduler not implemented")
	}
	sys := system.New(scheduler, cpus)
	sys.Run(schedule)
}

func TestRoundRobinExample(t *testing.T) {
	job.ResetJobCounter()
	schedule := []*system.Entry{
		j(t003, 0), // 1
		j(t003, 0), // 2
		j(t003, 0), // 3
	}
	cpus := cpu.NewCPUs(numCPUs)
	scheduler := rr.New(cpus, t005)
	if scheduler == nil {
		t.Fatalf("rr scheduler not implemented")
	}
	sys := system.New(scheduler, cpus)
	sys.Run(schedule)
}

func TestStrideScheduler(t *testing.T) {
	job.ResetJobCounter()
	schedule := []*system.Entry{
		k(t010, 0, 100), // a
		k(t010, 0, 50),  // b
		k(t010, 0, 250), // c
	}
	if schedule[0].Job == nil {
		t.Fatal("stride.NewJob not implemented")
	}
	cpus := cpu.NewCPUs(numCPUs)
	scheduler := stride.New(cpus, t001)
	if scheduler == nil {
		t.Fatalf("stride scheduler not implemented")
	}
	sys := system.New(scheduler, cpus)
	sys.Run(schedule)
}

func TestShortestJobFirst(t *testing.T) {
	tests := []struct {
		name     string
		jobs     []*system.Entry
		schedule schedules.JobMap
	}{
		{
			name:     "Sjf_2_jobs",
			jobs:     GetJobSchedule(),
			schedule: schedules.Job_ticks_Sjf_2_jobs,
		},
		{
			name:     "Sjf_Idle_time",
			jobs:     GetJobSchedule_Idle(),
			schedule: schedules.Job_ticks_Sjf_idle_time,
		},
		{
			name:     "Sjf_Many_jobs",
			jobs:     GetJobSchedule_ManyJobs(),
			schedule: schedules.Job_ticks_Sjf_many_jobs,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cpus := cpu.NewCPUs(numCPUs)
			scheduler := sjf.New(cpus)
			if scheduler == nil {
				t.Fatalf("sjf scheduler not implemented")
			}
			sys := system.New(scheduler, cpus)
			sys.RunCheck(t, test.jobs, test.schedule)
		})
	}
}

func TestRoundRobin(t *testing.T) {
	tests := []struct {
		name     string
		jobs     []*system.Entry
		schedule schedules.JobMap
	}{
		{
			name:     "Rr_2_jobs",
			jobs:     GetJobSchedule(),
			schedule: schedules.Job_ticks_Rr_2_jobs,
		},
		{
			name:     "Rr_Idle_time",
			jobs:     GetJobSchedule_Idle(),
			schedule: schedules.Job_ticks_Rr_idle_time,
		},
		{
			name:     "Rr_Many_jobs",
			jobs:     GetJobSchedule_ManyJobs(),
			schedule: schedules.Job_ticks_Rr_many_jobs,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cpus := cpu.NewCPUs(numCPUs)
			scheduler := rr.New(cpus, t005)
			if scheduler == nil {
				t.Fatalf("rr scheduler not implemented")
			}
			sys := system.New(scheduler, cpus)
			sys.RunCheck(t, test.jobs, test.schedule)
		})
	}
}

func TestStride(t *testing.T) {
	tests := []struct {
		name     string
		jobs     []*system.Entry
		schedule schedules.JobMap
	}{
		{
			name:     "Stride_2_jobs",
			jobs:     GetStrideJobSchedule_2jobs(),
			schedule: schedules.Job_ticks_Stride_2_jobs,
		},
		{
			name:     "Stride_example",
			jobs:     GetStrideJobSchedule(),
			schedule: schedules.Job_ticks_Stride_example,
		},
		{
			name:     "Stride_Many_jobs",
			jobs:     GetStrideJobSchedule_ManyJobs(),
			schedule: schedules.Job_ticks_Stride_many_jobs,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.jobs[0].Job == nil {
				t.Fatal("stride.NewJob not implemented")
			}
			cpus := cpu.NewCPUs(numCPUs)
			scheduler := stride.New(cpus, t005)
			if scheduler == nil {
				t.Fatalf("stride scheduler not implemented")
			}
			sys := system.New(scheduler, cpus)
			sys.RunCheck(t, test.jobs, test.schedule)
		})
	}
}
