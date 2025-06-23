package synctestpractice

import (
	"slices"
	"testing"
	"testing/synctest"
	"time"
)

const jobSleep = time.Second

type testJob struct {
	num int
}

func (t *testJob) Run() {
	time.Sleep(jobSleep)
	t.num = 2 * t.num
}

func (t *testJob) GetResult() Result {
	return Result{
		Output: t.num,
	}
}

func TestWorkerPool(t *testing.T) {
	workers := 2
	totalJobs := 5

	synctest.Test(t, func(*testing.T) {
		// start := time.Now() // 記錄「虛擬」起始時間
		pool := NewWorkerPool(workers)

		// pool 必須再放入job前就要起起來, 因為job沒有buffer
		// 另外 用 goroutine 不要塞住主線
		go pool.Run()

		// 收集結果
		var got []int
		done := make(chan struct{})

		go func() {
			for r := range pool.GetResultCh() {
				got = append(got, r.Output.(int))
			}

			close(done)
		}()

		for i := range totalJobs {
			pool.Push(&testJob{num: i})
		}

		// 通知做完了
		pool.Close()

		// time.Sleep(time.Duration(totalJobs) * jobSleep)
		// synctest.Wait()
		<-done

		slices.Sort(got)
		want := []int{0, 2, 4, 6, 8}
		if !slices.Equal(got, want) {
			t.Fatalf("結果錯誤：got %v, want %v", got, want)
		}
	})
}
