package paging

// findFreeFrames returns indices for n free frames.
// If there are not enough free frames available, an error is returned.
func (fl *freeList) findFreeFrames(n int) ([]int, error) {
	if n < 0 {
		return make([]int, 0), errArgOutOfBounds
	}
	if n == 0 {
		return make([]int, 0), nil
	}
	freeFrames := make([]int, n)
	total := 0
	for i, frame := range fl.freeList {
		if frame {
			freeFrames[total] = i
			total++
		}
		if total == n {
			break
		}
	}
	if total < n {
		return freeFrames[0:total:total], errOutOfMemory
	}
	return freeFrames, nil
}
