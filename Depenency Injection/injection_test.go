package main

import (
	"bytes"
	"testing"
)



func TestInjection(t *testing.T){
	buffer := bytes.Buffer{}
	Greet(&buffer , "chris")

	got := buffer.String()
	want := "Hello, chris"
	if got != want{
			t.Errorf("got %q, want %q", got, want)
		}
	
}


