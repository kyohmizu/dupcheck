package main

import (
    "bytes"
	"testing"
	"os"
	"io"
	"strings"
)

var buffer *bytes.Buffer

func TestDupcheck(t *testing.T) {
    r, w, err := os.Pipe()
    if err != nil {
        t.Fatal(err)
    }
    stdout := os.Stdout
    os.Stdout = w

    dupcheck("test")

    os.Stdout = stdout
    w.Close()

    var buf bytes.Buffer
	io.Copy(&buf, r)
	s := strings.TrimRight(buf.String(), "\n")
	
	want := "aaa"
    if s != want {
        t.Errorf("dupcheck(\"test\")) = %s, want %s", s, want)
    }
}
