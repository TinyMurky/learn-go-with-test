package synctestpractice

import (
	"fmt"
	"reflect"
	"testing"
	"testing/synctest"
	"time"
)

// func TestAsyncUpdate_1(t *testing.T) {
// 	shared := 0
// 	asyncUpdate(&shared)
// 	// assertEqual(t, shared, 1) // 有這行就會 got 0, want 1
// 	time.Sleep(5 * time.Microsecond)
// 	assertEqual(t, shared, 2)
// }

func TestAsyncUpdate_2(t *testing.T) {
	synctest.Test(t, func(*testing.T) {
		fmt.Printf("Before test: %v\n", time.Now()) // Before test: 2000-01-01 08:00:00 +0800 CST
		shared := 0
		asyncUpdate(&shared)

		// 第一次 Wait，goroutine 已經執行 shared = 1，並卡在 Sleep()
		synctest.Wait()
		fmt.Printf("After first test: %v\n", time.Now()) // After first test: 2000-01-01 08:00:00 +0800 CST
		assertEqual(t, shared, 1)

		time.Sleep(time.Microsecond - 1*time.Nanosecond)
		fmt.Printf("After 999 nanosecond sleep: %v\n", time.Now()) // After 999 nanosecond sleep: 2000-01-01 08:00:00.000000999 +0800 CST

		synctest.Wait()
		fmt.Printf("After Second test: %v\n", time.Now()) // After Second test: 2000-01-01 08:00:00.000000999 +0800 CST
		assertEqual(t, shared, 1)

		time.Sleep(time.Nanosecond)
		fmt.Printf("After 1 nanosecond sleep: %v\n", time.Now()) // After 1 nanosecond sleep: 2000-01-01 08:00:00.000001 +0800 CST

		synctest.Wait()
		fmt.Printf("After Third test: %v\n", time.Now()) // After Third test: 2000-01-01 08:00:00.000001 +0800 CST
		assertEqual(t, shared, 2)
	})
}

func assertEqual[T any](t testing.TB, got, want T) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
