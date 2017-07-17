package ring

import "container/ring"

// Ring represents a container/ring with seperate reader and writer interfaces.
// It will only read non-nil values and will overwrite the oldest values on
// subsequent Write()'s. Originally designed as a FIFO packet buffer.
type Ring struct {
	reader *ring.Ring
	writer *ring.Ring
}

// Return a new Ring with the specified capactiy.
func NewRing(size int) *Ring {
	r := &Ring{
		reader: ring.New(size),
	}
	r.writer = r.reader
	return r
}

// Read the next value from the ring. If the value is nil, we will sit here
// until there is a non-nil value to return, then advance for the next Read().
func (r *Ring) Read() interface{} {
	s := r.reader.Value
	// If we have read nil do not advance to the next item
	if s == nil {
		return nil
	}
	// Clear this item and advance to the next one
	r.reader.Value = nil
	r.reader = r.reader.Next()
	return s
}

// Write the value to the Ring. Write's will always succeed and will overwrite
// the oldest elements in the Ring first.
func (r *Ring) Write(value interface{}) {
	r.writer.Value = value
	r.writer = r.writer.Next()
}
