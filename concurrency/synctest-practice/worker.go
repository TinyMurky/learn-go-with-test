package synctestpractice

import "sync"

type Result struct {
	Output any
	Error  error
}

type Job interface {
	Run()
	GetResult() Result
}

type WorkerPool struct {
	jobCh     chan Job
	resultCh  chan Result
	workerCnt int
	wg        sync.WaitGroup
}

func NewWorkerPool(workerCnt int) *WorkerPool {
	buf := workerCnt
	wp := &WorkerPool{
		jobCh:     make(chan Job),
		resultCh:  make(chan Result, buf),
		workerCnt: workerCnt,
	}

	return wp
}

func (p *WorkerPool) Push(j Job) {
	p.jobCh <- j
}

func (p *WorkerPool) Close() {
	close(p.jobCh)
}

// Run 必須在「送工作之前」先呼叫
func (p *WorkerPool) Run() {
	for range p.workerCnt {
		p.wg.Add(1)
		go func() {
			defer p.wg.Done()
			// range 在close 的時候停下
			for job := range p.jobCh {
				job.Run()
				p.resultCh <- job.GetResult()
			}

		}()
	}
	go func() {
		p.wg.Wait()
		close(p.resultCh)
	}()
}

func (p *WorkerPool) GetResultCh() <-chan Result {
	return p.resultCh
}
