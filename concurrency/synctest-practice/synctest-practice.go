package synctestpractice

import "time"

func asyncUpdate(shared *int) {
	go func() {
		*shared = 1 // 第一次更新
		time.Sleep(1 * time.Microsecond)
		*shared = 2 // 第二次更新
	}()
}
