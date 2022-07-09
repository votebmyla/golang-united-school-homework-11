package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	var wg sync.WaitGroup
	var mutx sync.Mutex
	sem := make(chan struct{}, pool)
	for i := 0; i < int(n); i++ {
		wg.Add(1)
		sem <- struct{}{}
		go func(id int64) {
			defer wg.Done()
			oneUser := getOne(id)
			mutx.Lock()
			res = append(res, oneUser)
			mutx.Unlock()
			<-sem
		}(int64(i))
	}
	wg.Wait()
	return res
}
