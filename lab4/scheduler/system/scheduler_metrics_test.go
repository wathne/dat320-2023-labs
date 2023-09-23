package system

import (
	"dat320/lab4/scheduler/cpu"
	"dat320/lab4/scheduler/fifo"
	"dat320/lab4/scheduler/job"
	"dat320/lab4/scheduler/rr"
	"reflect"
	"testing"
	"time"
)

const (
	t001   = 1 * time.Millisecond
	t002   = 2 * time.Millisecond
	t010   = 10 * time.Millisecond
	t020   = 20 * time.Millisecond
	t030   = 30 * time.Millisecond
	t090   = 90 * time.Millisecond
	t100   = 100 * time.Millisecond
	t110   = 110 * time.Millisecond
	t120   = 120 * time.Millisecond
	t070   = 70 * time.Millisecond
	t028   = 28 * time.Millisecond
	t029   = 29 * time.Millisecond
	t059   = 59 * time.Millisecond
	t05967 = 59*time.Millisecond + 666667*time.Nanosecond
	t06333 = 63*time.Millisecond + 333333*time.Nanosecond
	t10333 = 103*time.Millisecond + 333333*time.Nanosecond
	t00333 = 333333 * time.Nanosecond
)

type SchedulerParams struct {
	name    string
	cpus    int
	quantum time.Duration
}

func (sp SchedulerParams) NewScheduler(cpus []*cpu.CPU) (scheduler Scheduler) {
	switch sp.name {
	case "fifo":
		scheduler = fifo.New(cpus)
	case "rr":
		scheduler = rr.New(cpus, sp.quantum)
	}
	return
}

// j is a helper function to create a scheduling entry with a job and its arrival time.
// With this helper, the job zero size and the given estimated running time.
var j = func(estimated, arrival time.Duration) *Entry {
	return &Entry{Job: job.New(0, estimated), Arrival: arrival}
}

// BookSchedule1 returns a schedule with three jobs with arrival time 0 and completion time 10 ms.
func BookSchedule1() Schedule {
	job.ResetJobCounter()
	return []*Entry{
		j(t010, 0), // a
		j(t010, 0), // b
		j(t010, 0), // c
	}
}

// BookSchedule2 returns a schedule with three jobs with arrival time 0 and completion time 10 ms, except the first job which takes 100 ms.
func BookSchedule2() Schedule {
	job.ResetJobCounter()
	return []*Entry{
		j(t100, 0), // a
		j(t010, 0), // b
		j(t010, 0), // c
	}
}

// BookSchedule3 returns a schedule with three jobs the first one with arrival time 0, and the last two arrive at 10 ms.
// The completion time of the first job is 100 ms, while the last two takes 10 ms.
func BookSchedule3() Schedule {
	job.ResetJobCounter()
	return []*Entry{
		j(t100, 0),    // a
		j(t010, t010), // b
		j(t010, t010), // c
	}
}

var metricsTests = []struct {
	name     string
	params   SchedulerParams
	schedule Schedule
	tt       []time.Duration
	rt       []time.Duration
}{
	{
		name:     "fifo/book_schedule1",
		params:   SchedulerParams{name: "fifo", cpus: 1},
		schedule: BookSchedule1(),
		tt:       []time.Duration{t010, t020, t030},
		rt:       []time.Duration{0, t010, t020},
	},
	{
		name:     "fifo/book_schedule2",
		params:   SchedulerParams{name: "fifo", cpus: 1},
		schedule: BookSchedule2(),
		tt:       []time.Duration{t100, t110, t120},
		rt:       []time.Duration{0, t100, t110},
	},
	{
		name:     "fifo/book_schedule3",
		params:   SchedulerParams{name: "fifo", cpus: 1},
		schedule: BookSchedule3(),
		tt:       []time.Duration{t100, t100, t110},
		rt:       []time.Duration{0, t090, t100},
	},
	{
		name:     "rr/book_schedule1/q=1ms",
		params:   SchedulerParams{name: "rr", cpus: 1, quantum: t001},
		schedule: BookSchedule1(),
		tt:       []time.Duration{t028, t029, t030},
		rt:       []time.Duration{0, t001, t002},
	},
	{
		name:     "rr/book_schedule2/q=1ms",
		params:   SchedulerParams{name: "rr", cpus: 1, quantum: t001},
		schedule: BookSchedule2(),
		tt:       []time.Duration{t120, t029, t030},
		rt:       []time.Duration{0, t001, t002},
	},
	{
		name:     "rr/book_schedule3/q=1ms",
		params:   SchedulerParams{name: "rr", cpus: 1, quantum: t001},
		schedule: BookSchedule3(),
		tt:       []time.Duration{t120, t028, t029},
		rt:       []time.Duration{0, 0, t001},
	},
}

func TestSingleJobMetrics(t *testing.T) {
	for _, test := range metricsTests {
		t.Run(test.name, func(t *testing.T) {
			cpus := cpu.NewCPUs(test.params.cpus)
			scheduler := test.params.NewScheduler(cpus)
			if reflect.ValueOf(scheduler).IsNil() {
				t.Fatalf("%s scheduler not implemented", test.params.name)
			}

			sys := New(scheduler, cpus)
			sys.Run(test.schedule)
			for i, job := range test.schedule {
				gotTT := job.Job.TurnaroundTime()
				if test.tt[i] != gotTT {
					t.Errorf("TurnaroundTime() = %v, expected %v", gotTT, test.tt[i])
				}
				gotRT := job.Job.ResponseTime()
				if test.rt[i] != gotRT {
					t.Errorf("ResponseTime() = %v, expected %v", gotRT, test.rt[i])
				}
			}
		})
	}
}

var avgTests = []struct {
	name     string
	params   SchedulerParams
	schedule Schedule
	avgTT    time.Duration
	avgRT    time.Duration
}{
	{
		name:     "fifo/book_schedule1",
		params:   SchedulerParams{name: "fifo", cpus: 1},
		schedule: BookSchedule1(),
		avgTT:    t020, // textbook chapter 7, page 3.
		avgRT:    t010, // response_time = (0+10+20)/3 = 10
	},
	{
		name:     "fifo/book_schedule2",
		params:   SchedulerParams{name: "fifo", cpus: 1},
		schedule: BookSchedule2(),
		avgTT:    t110, // textbook chapter 7, page 3.
		avgRT:    t070, // response_time = (0+100+110)/3 = 70
	},
	{
		name:     "fifo/book_schedule3",
		params:   SchedulerParams{name: "fifo", cpus: 1},
		schedule: BookSchedule3(),
		avgTT:    t10333, // turnaround_time = (100+(110-10)+(120-10))/3 = 103.33
		avgRT:    t06333, // response_time = (0+(100-10)+(110-10))/3 = 63.33
	},
	{
		name:     "rr/book_schedule1/q=1ms",
		params:   SchedulerParams{name: "rr", cpus: 1, quantum: t001},
		schedule: BookSchedule1(),
		avgTT:    t029, // turnaround_time = (28+29+30)/3 = 29
		avgRT:    t001, // response_time = (0+1+2)/3 = 1
	},
	{
		name:     "rr/book_schedule2/q=1ms",
		params:   SchedulerParams{name: "rr", cpus: 1, quantum: t001},
		schedule: BookSchedule2(),
		avgTT:    t05967, // turnaround_time = (120+29+30)/3 = 59.67
		avgRT:    t001,   // response_time = (0+1+2)/3 = 1
	},
	{
		name:     "rr/book_schedule3/q=1ms",
		params:   SchedulerParams{name: "rr", cpus: 1, quantum: t001},
		schedule: BookSchedule3(),
		avgTT:    t059,   // turnaround_time = (120+38-10+39-10)/3 = 59
		avgRT:    t00333, // response_time = (0+0+1)/3 = 0.333
	},
}

func TestAverageMetrics(t *testing.T) {
	for _, test := range avgTests {
		t.Run(test.name, func(t *testing.T) {
			cpus := cpu.NewCPUs(test.params.cpus)
			scheduler := test.params.NewScheduler(cpus)
			if reflect.ValueOf(scheduler).IsNil() {
				t.Fatalf("%s scheduler not implemented", test.params.name)
			}

			sys := New(scheduler, cpus)
			sys.Run(test.schedule)
			gotTT := test.schedule.Avg(func(j *job.Job) time.Duration {
				return j.TurnaroundTime()
			})
			if test.avgTT != gotTT {
				t.Errorf("AvgTT() = %v, expected %v", gotTT, test.avgTT)
			}
			gotTT = test.schedule.AvgTurnaroundTime()
			if test.avgTT != gotTT {
				t.Errorf("AvgTT() = %v, expected %v", gotTT, test.avgTT)
			}
			gotRT := test.schedule.Avg(func(j *job.Job) time.Duration {
				return j.ResponseTime()
			})
			if test.avgRT != gotRT {
				t.Errorf("AvgRT() = %v, expected %v", gotRT, test.avgRT)
			}
			gotRT = test.schedule.AvgResponseTime()
			if test.avgRT != gotRT {
				t.Errorf("AvgRT() = %v, expected %v", gotRT, test.avgRT)
			}
		})
	}
}
