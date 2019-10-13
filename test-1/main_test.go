package main

import "testing"

var input string = "omama"
var expected string = "aaomm"

func TestSplitChar(t *testing.T) {
	splitChar(input)
	t.Logf("Word : %s\n", input)
	t.Logf("Output : %s\n", output)
	t.Logf("Expected : %s\n", expected)

	if output != expected {
		t.Errorf("Error! expected (%s) got (%s)", expected, output)
	}

}

func BenchmarkSplitChar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		splitChar(input)
	}
}
