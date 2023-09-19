package system

import (
	"dat320/lab4/scheduler/cpu"
	"dat320/lab4/scheduler/job"
	"dat320/lab4/scheduler/schedules"
	"dat320/lab4/scheduler/system/systime"
	"fmt"
	"os"
	"sort"
	"testing"
	"text/tabwriter"
	"time"

	"github.com/google/go-cmp/cmp"
)

type System struct {
	systemTime   time.Duration
	jobsFinished int
	scheduler    Scheduler
	cpus         []*cpu.CPU
	w            *tabwriter.Writer
}

func New(sched Scheduler, cpus []*cpu.CPU) *System {
	s := &System{
		cpus:      cpus,
		scheduler: sched,
		w:         tabwriter.NewWriter(os.Stdout, 2, 8, 2, ' ', 0),
	}
	return s
}

func (s *System) Run(schedule []*Entry) {
	sort.Slice(schedule, func(i, j int) bool {
		return schedule[i].Arrival < schedule[j].Arrival
	})
	s.printHeader()
	for s.jobsFinished < len(schedule) {
		for _, e := range schedule {
			if e.Arrival == s.systemTime {
				e.Job.Scheduled(s)
				s.scheduler.Add(e.Job)
			}
		}
		s.tick()
		s.printCPUs()
	}
	if *verbose {
		s.w.Flush()
	}
}

// RunCheck runs the schedule and compare with the expected jobMap.
func (s *System) RunCheck(t *testing.T, schedule []*Entry, jobMap schedules.JobMap) {
	t.Helper()

	sort.Slice(schedule, func(i, j int) bool {
		return schedule[i].Arrival < schedule[j].Arrival
	})
	cnt := 0
	const maxCnt = 200

	for ; s.jobsFinished < len(schedule) && cnt < maxCnt; cnt++ {
		for _, e := range schedule {
			if e.Arrival == s.systemTime {
				e.Job.Scheduled(s)
				s.scheduler.Add(e.Job)
			}
		}
		s.tick()
		if *verbose {
			s.printCPUs()
		}
		gotJobs := s.InspectCPUs()
		wantJobs := jobMap[s.systemTime]
		if !cmp.Equal(wantJobs, gotJobs, cmpJob) {
			if *verbose {
				for i := 0; i < 1; i++ {
					t.Errorf("MISMATCH CPU %d: +got %v -want %v\n", i, gotJobs[i], wantJobs[i])
				}
			}
		}
	}
	if *verbose {
		s.w.Flush()
	}
	if cnt > maxCnt-1 {
		t.Error("max count exceeded")
	}
}

// Now returns the current scheduler/system time.
func (s *System) Now() time.Duration {
	return s.systemTime
}

func (s *System) tick() {
	fmt.Fprintf(s.w, "%v\t", s.systemTime)
	s.jobsFinished += s.scheduler.Tick(s.systemTime)
	s.systemTime += systime.TickDuration
}

func (s *System) printHeader() {
	fmt.Fprint(s.w, "Tick\t")
	for _, cpu := range s.cpus {
		fmt.Fprintf(s.w, "%v\t", cpu.Header())
	}
	fmt.Fprintln(s.w)
}

func (s *System) printCPUs() {
	for _, cpu := range s.cpus {
		fmt.Fprintf(s.w, "%v\t", cpu)
	}
	fmt.Fprintln(s.w)
}

// InspectCPUs returns the jobs currently running on the CPUs.
// The returned map is indexed by the CPU's ID.
func (s *System) InspectCPUs() map[int]job.Job {
	cpuJobs := make(map[int]job.Job)
	for _, cpu := range s.cpus {
		curJob := cpu.CurrentJob()
		if curJob != nil {
			cpuJobs[cpu.ID()] = curJob.Clone()
		}
	}
	return cpuJobs
}

var cmpJob = cmp.AllowUnexported(job.Job{})
