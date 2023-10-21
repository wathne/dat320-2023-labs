package paging

// NoEntry is produced when no entry matching a request exists
const NoEntry = -1

// PageTable is a per-process data structure which holds translations from virtual page numbers to physical frame numbers
type PageTable struct {
	frameIndices []int // maps virtual page number (index) to physical frame number (content)
}

// Append adds pages to a page table
func (pt *PageTable) Append(pages []int) {
	// Implement append functionality for the pagetable
	pt.frameIndices = append(pt.frameIndices, pages...)
}

// Free removes the n last pages from the page table and returns the removed entries
func (pt *PageTable) Free(n int) ([]int, error) {
	// Implement free functionality for the pagetable
	if n < 0 {
		return make([]int, 0, 0), errArgOutOfBounds
	}
	if n == 0 {
		return make([]int, 0, 0), nil
	}
	a := &pt.frameIndices
	if n > len(*a) {
		return make([]int, 0, 0), errFreeOutOfBounds
	}
	freedFrames := make([]int, n, n)
	copy(freedFrames, (*a)[len(*a) - n : len(*a) : len(*a)])
	*a = (*a)[0 : len(*a) - n : len(*a) - n]
	return freedFrames, nil
}

// Lookup returns the mapping of a virtual page number to a physical frame number, or an error if it does not exist.
func (pt *PageTable) Lookup(virtualPageNum int) (frameIndex int, err error) {
	// Implement lookup functionality for the pagetable
	if virtualPageNum < 0 {
		return NoEntry, errArgOutOfBounds
	}
	if virtualPageNum >= len(pt.frameIndices) {
		return NoEntry, errAddressOutOfBounds
	}
	return pt.frameIndices[virtualPageNum], nil
}

// Len returns the length of the page table
func (pt *PageTable) Len() int {
	return len(pt.frameIndices)
}
