package barrier

import (
	"fmt"
	"sync"
	"testing"
)

func TestBarrier(t *testing.T) {
	b := New(100)
	wg := &sync.WaitGroup{}
	m := &sync.Mutex{}
	tmp := ""

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(b *Barrier) {
			defer wg.Done()
			m.Lock()
			tmp += fmt.Sprintf("%s", "1")
			m.Unlock()
			b.BarrierWait()
			m.Lock()
			tmp += fmt.Sprintf("%s", "2")
			m.Unlock()
		}(b)
	}

	wg.Wait()
	fmt.Println(tmp, len(tmp))
	for i := 0; i < 100; i++ {
		if tmp[i] != byte('1') || tmp[100+i] != byte('2') {
			t.Error("error 1")
		}
	}
}
