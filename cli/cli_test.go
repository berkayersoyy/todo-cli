package cli

import (
	"bufio"
	"io"
	"testing"
)

type R struct {
	Data string
	done bool
}

func (r *R) Read(p []byte) (n int, err error) {
	copy(p, []byte(r.Data))
	if r.done {
		return 0, io.EOF
	}
	r.done = true
	return len([]byte(r.Data)), nil
}

func NewR(data string) *R {
	return &R{data, false}
}

func TestReadConsole(t *testing.T) {
	rd := NewR("todo -h")
	sc := bufio.NewReader(rd)
	_, _, err := ReadConsole(sc)
	if err != nil {
		t.Error(err)
	}

}
