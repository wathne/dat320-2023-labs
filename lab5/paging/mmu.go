package paging

import (
	"math"
)

// MMU is the structure for the simulated memory management unit.
type MMU struct {
	frames    [][]byte // contains memory content in form of frames[frameIndex][offset]
	frameSize int
	freeList                     // tracks free physical frames
	processes map[int]*PageTable // contains page table for each process (key=pid)
}

// OffsetLookupTable gives the bit mask corresponding to a virtual address's offset of length n,
// where n is the table index. This table can be used to find the offset mask needed to extract
// the offset from a virtual address. It supports up to 32-bit wide offset masks.
//
// OffsetLookupTable[0] --> 0000 ... 0000
// OffsetLookupTable[1] --> 0000 ... 0001
// OffsetLookupTable[2] --> 0000 ... 0011
// OffsetLookupTable[3] --> 0000 ... 0111
// OffsetLookupTable[8] --> 0000 ... 1111 1111
// etc.
var OffsetLookupTable = []int{
	// 0000, 0001, 0011, 0111, 1111, etc.
	0x0000000, 0x00000001, 0x00000003, 0x00000007,
	0x000000f, 0x0000001f, 0x0000003f, 0x0000007f,
	0x00000ff, 0x000001ff, 0x000003ff, 0x000007ff,
	0x0000fff, 0x00001fff, 0x00003fff, 0x00007fff,
	0x000ffff, 0x0001ffff, 0x0003ffff, 0x0007ffff,
	0x00fffff, 0x001fffff, 0x003fffff, 0x007fffff,
	0x0ffffff, 0x01ffffff, 0x03ffffff, 0x07ffffff,
	0xfffffff, 0x1fffffff, 0x3fffffff, 0x7fffffff, 0xffffffff,
}

// NewMMU creates a new MMU with a memory of memSize bytes.
// memSize should be >= 1 and a multiple of frameSize.
func NewMMU(memSize, frameSize int) *MMU {
	// Task 2: initialize the MMU object
	numFrames := memSize / frameSize
	frames := make([][]byte, numFrames)
	// A trick to improve memory locality.
	bytes := make([]byte, numFrames*frameSize)
	for i := 0; i < numFrames; i++ {
		// Full slice expressions: http://golang.org/ref/spec#Slice_expressions
		// a[low : high : max], it controls the resulting slice's capacity by
		// setting it to max - low. In our case, the capacity of each frame is:
		// ((i+1) * frameSize) - (i * frameSize) = frameSize
		frames[i] = bytes[i*frameSize : (i+1)*frameSize : (i+1)*frameSize]
	}
	return &MMU{
		frames:    frames,
		frameSize: frameSize,
		freeList:  newFreeList(numFrames),
		processes: make(map[int]*PageTable, 0),
	}
}

// Alloc allocates n bytes of memory for process pid.
// The allocated memory is added to the process's page table.
// The process is given a page table if it doesn't already have one,
// unless an out of memory error occurred.
func (mmu *MMU) Alloc(pid, n int) error {
	// Task 2: implement memory allocation
	// Suggested approach:
	// - calculate #frames needed to allocate n bytes, error if not enough free frames
	// - if process pid has no page table, create one for it
	// - determine which frames to allocate to the process
	// - add the frames to the process's (identified by pid) page table and
	// - update the free list
	if n < 1 {
		return errNothingToAllocate
	}
	framesRequired := int(math.Ceil(float64(n) / float64(mmu.frameSize)))
	if framesRequired > mmu.freeList.calculateNumFreeFrames() {
		return errOutOfMemory
	}
	_, ok := mmu.processes[pid]
	if !ok {
		mmu.processes[pid] = &PageTable{
			frameIndices: make([]int, 0),
		}
	}
	freeFrames, err1 := mmu.freeList.findFreeFrames(framesRequired)
	if err1 != nil {
		return err1
	}
	mmu.processes[pid].Append(freeFrames)
	err2 := mmu.freeList.removeFrames(freeFrames)
	if err2 != nil {
		return err2
	}
	return nil
}

// Write writes content to the given process's address space starting at virtualAddress.
func (mmu *MMU) Write(pid, virtualAddress int, content []byte) error {
	// Task 3: implement writing
	// Suggested approach:
	// - check valid pid (must have a page table)
	// - translate the virtual address
	// - check if the memory must be extended in order to write the content
	//   from the given starting address
	// - attempt to allocate more memory if necessary to complete the write
	// - sequentially write content into the known-to-be-valid address space
	vpn, offset, err1 := mmu.translateAndCheck(pid, virtualAddress)
	if err1 != nil {
		return err1
	}
	frames := &mmu.frames
	frameSize := mmu.frameSize
	pageTable := mmu.processes[pid]
	bytesAvailable := (pageTable.Len()-vpn)*frameSize - offset
	bytesRequired := len(content)
	if bytesRequired > bytesAvailable {
		err2 := mmu.Alloc(pid, bytesRequired-bytesAvailable)
		if err2 != nil {
			return err2
		}
	}
	pfn, err3 := pageTable.Lookup(vpn)
	if err3 != nil {
		return err3
	}
	for _, byte := range content {
		if offset >= frameSize {
			offset = 0
			vpn++
			pfn, err3 = pageTable.Lookup(vpn)
			if err3 != nil {
				return err3
			}
		}
		(*frames)[pfn][offset] = byte
		offset++
	}
	return nil
}

// Read returns content of size n bytes from the given process's address space starting at virtualAddress.
func (mmu *MMU) Read(pid, virtualAddress, n int) (content []byte, err error) {
	// Task 3: implement reading
	// Suggested approach:
	// - check valid pid (must have a page table)
	// - translate the virtual address
	// - (optional) determine if it's possible to read the requested number
	//   of bytes before starting to read the memory content
	// - read and return the requested memory content
	if n < 1 {
		return nil, errNothingToRead
	}
	vpn, offset, err1 := mmu.translateAndCheck(pid, virtualAddress)
	if err1 != nil {
		return nil, err1
	}
	frames := &mmu.frames
	frameSize := mmu.frameSize
	pageTable := mmu.processes[pid]
	bytesAvailable := (pageTable.Len()-vpn)*frameSize - offset
	if n > bytesAvailable {
		return nil, errAddressOutOfBounds
	}
	pfn, err2 := pageTable.Lookup(vpn)
	if err2 != nil {
		return nil, err2
	}
	content = make([]byte, n)
	for i := range content {
		if offset >= frameSize {
			offset = 0
			vpn++
			pfn, err2 = pageTable.Lookup(vpn)
			if err2 != nil {
				return nil, err2
			}
		}
		content[i] = (*frames)[pfn][offset]
		offset++
	}
	return content, nil
}

// Free is called by a process's Free() function to free some of its allocated memory.
func (mmu *MMU) Free(pid, n int) error {
	// Task 4: implement freeing of memory
	// Suggested approach:
	// - check valid pid (must have a page table)
	// - check if there are at least n entries in the page table of pid
	// - free n pages
	// - set all the bytes in the freed memory to the value 0
	// - re-add the freed frames to the free list
	if n < 1 {
		return errNothingToAllocate
	}
	pageTable, ok := mmu.processes[pid]
	if !ok {
		return errInvalidProcess
	}
	freedFrames, err1 := pageTable.Free(n)
	if err1 != nil {
		return err1
	}
	var i int // Physical frame number.
	var j int // Offset within the frame.
	for k := range freedFrames {
		i = freedFrames[k]
		for j = range mmu.frames[i] {
			mmu.frames[i][j] = 0
		}
	}
	err2 := mmu.freeList.addFrames(freedFrames)
	if err2 != nil {
		return err2
	}
	return nil
}

// extract returns the virtual page number and offset for the given virtual address,
// and the number of bits in the offset n.
func extract(virtualAddress, n int) (vpn, offset int) {
	// Implement virtual address translation as described in
	// the Virtual Addresses section of the README.
	// The procedure is described in detail in Chapter 18.1 of the textbook.
	// HINT: It can be solved quite easily with bitwise operators.
	// (see https://golang.org/ref/spec#Arithmetic_operators ).
	// You might also find the provided log2 function and the OffsetLookupTable
	// table useful for this purpose.
	vpn = virtualAddress >> n
	offset = virtualAddress & OffsetLookupTable[n]
	//frameSize := uint32(1) << n
	return vpn, offset
}

// translateAndCheck returns the virtual page number and offset for the given virtual address.
// If the virtual address is invalid for process pid, an error is returned.
func (mmu *MMU) translateAndCheck(pid, virtualAddress int) (vpn, offset int, err error) {
	// Implement virtual address translation as described in
	// the Virtual Addresses section of the README.
	// The procedure is described in detail in Chapter 18.1 of the textbook.
	// It is expected that this method calls the extract function above
	// to compute the VPN and offset to be returned from this function after
	// checking that the process has access to the returned VPN.
	// You might also find the provided log2 function useful to calculate one
	// of the inputs to the extract function.
	n := log2(mmu.frameSize)
	vpn, offset = extract(virtualAddress, n)
	pageTable, ok := mmu.processes[pid]
	if !ok {
		return vpn, offset, errInvalidProcess
	}
	_, err = pageTable.Lookup(vpn)
	return vpn, offset, err
}

// NOTE(wathne): Returns PFN instead of VPN.
func (mmu *MMU) translateAndCheckPFN(pid, virtualAddress int) (pfn, offset int, err error) {
	n := log2(mmu.frameSize)
	var vpn int
	vpn, offset = extract(virtualAddress, n)
	pageTable, ok := mmu.processes[pid]
	if !ok {
		return NoEntry, offset, errInvalidProcess
	}
	pfn, err = pageTable.Lookup(vpn)
	return pfn, offset, err
}

// log2 calculates m given n = 2^m.
func log2(n int) int {
	exp := 0
	for {
		if n%2 == 0 && n > 0 {
			exp++
		} else {
			return exp
		}
		n /= 2
	}
}
