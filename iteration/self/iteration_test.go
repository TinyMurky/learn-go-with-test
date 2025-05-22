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
// go test -bench=.
// 6047209是執行次數 ns 是耐秒 b 是 "bytes" allocs 是操作幾次記憶體配置

// goarch: amd64
// pkg: example.com/iteration
// cpu: AMD Ryzen 7 7735HS with Radeon Graphics
// BenchmarkRepeat3-16      6047209        190.0 ns/op       16 B/op        4 allocs/op
// PASS
// ok   example.com/iteration 1.313s

// 會慢的原因：
// string是immutable 所以每次concatenation 都會複製整個舊字串並重新配置記憶體
var result string

func BenchmarkRepeat3(b *testing.B) {
	// 教學的作法
	// for i := 0; i < b.N; i++ {
	// 	Repeat3("a")
	// }

	// lsp推薦的作法
	var r string
	for b.Loop() {
		r = Repeat3("a")
	}
	result = r // 可以避免 compiler 認為測試中計算結果沒有被使用或儲存 而略過執行
}

// Strings builder minimize memorry copy
// goarch: amd64
// pkg: example.com/iteration
// cpu: AMD Ryzen 7 7735HS with Radeon Graphics
// BenchmarkRepeat5-16     22450646         61.02 ns/op        8 B/op        1 allocs/op
// PASS
// ok   example.com/iteration 1.517s
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
