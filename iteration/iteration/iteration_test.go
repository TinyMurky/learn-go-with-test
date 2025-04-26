package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("It should repeat 5 times", func(t *testing.T) {

		repeated := Repeat("a")

		expect := "aaaaa"

		assertEqual(t, repeated, expect)
	})
}

// use `go test -bench=.` to run
// goos: linux                          // 作業系統是 Linux
// goarch: amd64                        // CPU 架構是 amd64
// pkg: example.come/iteration/iteration // 測試的 package 名稱
// cpu: AMD Ryzen 7 7735HS with Radeon Graphics // 實際執行測試的 CPU 型號

// BenchmarkRepeat-16                   // 基準測試的名稱（這個 `-16` 表示用了 16 個執行緒，也就是 GOMAXPROCS=16）
//     2210092                          // 總共跑了 2,210,092 次 benchmark 函式
//     453.3 ns/op                      // 每次執行平均耗時 453.3 奈秒
//     120 B/op                         // 每次執行平均分配了 120 bytes 記憶體
//     7 allocs/op                      // 每次執行平均會做 7 次記憶體配置（allocation）

// PASS                                 // 測試通過
// ok   example.come/iteration/iteration  1.558s  // 整個 benchmark 測試耗時 1.558 秒
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

func assertEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()

	if got != want {
		t.Errorf("Expect \"%v\", but got \"%v\"", want, got)
	}
}
