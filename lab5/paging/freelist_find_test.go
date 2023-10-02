package paging

import "testing"

func TestFindFreeList(t *testing.T) {

	// Create a new free list with 10 frames
	fl := newFreeList(10)

	// Find 5 free frames
	indices, err := fl.findFreeFrames(5)
	if err != nil {
		t.Fatal(err)
	}

	// Check that we got 5 indices
	if len(indices) != 5 {
		t.Fatalf("Expected 5 indices, got %d", len(indices))
	}

	// Check that the indices are in the range [0, 10)
	for _, i := range indices {
		if i < 0 || i >= 10 {
			t.Fatalf("Index %d is out of range [0, 10)", i)
		}
	}

	// Check that the indices are unique
	for i, j := range indices {
		for k, l := range indices {
			if i != k && j == l {
				t.Fatalf("Index %d is not unique", j)
			}
		}
	}

	// Check that the indices are marked as used
	for _, i := range indices {
		if fl.freeList[i] != true {
			t.Fatalf("Index %d is not marked as used", i)
		}
	}

	// Check that we can't find more than 10 free frames
	_, err = fl.findFreeFrames(11)
	if err == nil {
		t.Fatal("Expected an error when trying to find more than 10 free frames")
	}

}
