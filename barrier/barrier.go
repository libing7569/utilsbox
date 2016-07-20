package barrier

import "sync"

type Barrier struct {
	curCnt int
	maxCnt int
	cond   *sync.Cond
}

func New(maxCnt int) *Barrier {
	mutex := new(sync.Mutex)
	cond := sync.NewCond(mutex)
	curCnt := maxCnt
	return &Barrier{curCnt: curCnt, maxCnt: maxCnt, cond: cond}
}

func (barrier *Barrier) BarrierWait() {
	barrier.cond.L.Lock()
	if barrier.curCnt--; barrier.curCnt > 0 {
		barrier.cond.Wait()
	} else {
		barrier.cond.Broadcast()
		barrier.curCnt = barrier.maxCnt
	}
	barrier.cond.L.Unlock()
}
