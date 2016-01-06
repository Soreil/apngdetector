package apngdetector

import (
	"io/ioutil"
	"testing"
)

func TestDetector(t *testing.T) {
	tests := []struct {
		filename string
		expected bool
	}{
		{"infinitecomplex.png", true},
		{"still.png", false},
		{"loop.gif", false},
		{"infinite2frame.png", true},
	}
	for _, test := range tests {
		file, err := ioutil.ReadFile(test.filename)
		if err != nil {
			panic(err)
		}
		if Detect(file) != test.expected {
			t.Fatal(test)
		}
	}
}
