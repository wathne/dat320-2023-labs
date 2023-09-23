package systime

import "time"

const TickDuration = 1 * time.Millisecond

type SystemTime interface {
	Now() time.Duration
}
