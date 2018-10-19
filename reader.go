package slowio

import (
	"io"
	"time"
)

// reader is a slow reader.
type reader struct {
	delay time.Duration
	r     io.Reader
}

// NewReader instantiates a reader that reads 1 byte during during delay from a given reader.
func NewReader(delay time.Duration, r io.Reader) io.Reader {
	if delay < 0 {
		panic("delay must be non-negative")
	}
	return &reader{
		delay: delay,
		r:     r,
	}
}

func (r *reader) Read(data []byte) (int, error) {
	time.Sleep(r.delay)
	n, err := r.r.Read(data[:1])
	return n, err
}
