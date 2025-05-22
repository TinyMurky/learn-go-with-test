package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat1(t *testing.T) {
	repeated := Repeat1("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestRepeat2(t *testing.T) {
	repeated := Repeat2("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestRepeat3(t *testing.T) {
	repeated := Repeat3("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

// ==== Benchmark ====
func BenchmarkRepeat3(b *testing.B) {
	for b.Loop() {
		Repeat3("a")
	}
}

func BenchmarkRepeat5(b *testing.B) {
	for b.Loop() {
		Repeat5("a")
	}
}

func TestRepeat6(t *testing.T) {
	repeatTimes := 6
	expected := "aaaaaa"

	repeated := Repeat6("a", repeatTimes)
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func ExampleRepeat6() {
	repeated := Repeat6("a", 6)
	fmt.Println(repeated)
	// Output: aaaaaa
}
