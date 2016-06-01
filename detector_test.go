package apngdetector

import (
	"os"
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
		file, err := os.Open(test.filename)
		if err != nil {
			panic(err)
		}

		isApng, err := Detect(file)
		if err != nil {
			panic(err)
		}
		if isApng != test.expected {
			t.Fatal(test)
		}
	}
}
