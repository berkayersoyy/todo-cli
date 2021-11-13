package cli

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
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
	ex := "No argument such as 'todo a',for more help 'todo -h'"
	rd := NewR("todo a")
	sc := bufio.NewReader(rd)
	rstd := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	_, _, err := ReadConsole(sc)
	if err != nil {
		t.Error(err)
	}
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rstd
	outs := string(out)
	t.Errorf("s %s", out)
	if outs != ex {
		t.Errorf("TestReadConsole: got %s, want %s %s", outs, ex, out)
	}

}
