package pkg

var Pools = NewPool(100)

type Pool struct {
	work chan func()
	sem  chan struct{}
}

func NewPool(size int) *Pool {
	return &Pool{
		make(chan func()),
		make(chan struct{}, size),
	}
}

func (p *Pool) Schedule(task func()) {
	select {
	case p.work <- task:
	case p.sem <- struct{}{}:
		go p.worker(task)
	}
}

func (p *Pool) worker(task func()) {
	defer func() {
		<-p.sem
	}()
	for {
		task()
		task = <-p.work
	}
}
