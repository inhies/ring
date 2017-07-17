package ring

import "testing"

func TestRing(t *testing.T) {
	// Create a new buffer with a capacity of 5
	r := NewRing(5)
	// Write 1 to 3 int he first three sections
	for i := 1; i <= 3; i++ {
		r.Write(i) //1, 2, 3
	}
	// Read off the first two
	if r.Read() != 1 {
		t.Fail()
	}

	if r.Read() != 2 {
		t.Fail()
	}

	// Add the fourth item
	r.Write(4)

	// Read the third
	if r.Read() != 3 {
		t.Fail()
	}

	// Read the fourth
	if r.Read() != 4 {
		t.Fail()
	}

	// Attempt to read the fifth, but it's nil, so we shouldnt advance
	if r.Read() != nil {
		t.Fail()
	}

	// Add the fifth item and then read it
	r.Write(5)
	if r.Read() != 5 {
		t.Fail()
	}

	// Add the sixth item, which should overwrite the first
	r.Write(6)

	if r.Read() != 6 {
		t.Fail()
	}

	// We should have nil here since the 6th item was the last one put in
	if r.Read() != nil {
		t.Fail()
	}

}
