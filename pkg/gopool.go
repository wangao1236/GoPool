package pkg

import (
	"sync"
)

type Task interface {
	Run()
}

type Executor interface {
	Execute(task Task)
	Shutdown()
	Wait() chan struct{}
}

type executor struct {
	lock             sync.Mutex
	waitingTasks     []chan struct{}
	activeTasks      int64
	concurrencyLimit int64
	finish           chan struct{}
}

func (ex *executor) Execute(task Task) {
	ex.start(task)
}

func (ex *executor) Wait() chan struct{} {
	return ex.finish
}

func (ex *executor) Shutdown() {
	<-ex.finish
}

func (ex *executor) start(task Task) {
	startCh := make(chan struct{})
	stopCh := make(chan struct{})

	go startTask(startCh, stopCh, task)
	ex.enqueue(startCh)
	go ex.waitDone(stopCh)

}

func NewExecutor(concurrencyLimit int64) Executor {
	ex := &executor{
		waitingTasks:     make([]chan struct{}, 0),
		activeTasks:      0,
		concurrencyLimit: concurrencyLimit,
		finish:           make(chan struct{}),
	}
	return ex
}

func startTask(startCh, stopCh chan struct{}, task Task) {
	defer close(stopCh)

	<-startCh
	task.Run()
}

func (ex *executor) enqueue(startCh chan struct{}) {
	ex.lock.Lock()
	defer ex.lock.Unlock()

	if ex.concurrencyLimit == 0 || ex.activeTasks < ex.concurrencyLimit {
		close(startCh)
		ex.activeTasks++
	} else {
		ex.waitingTasks = append(ex.waitingTasks, startCh)
	}
}

func (ex *executor) waitDone(stopCh chan struct{}) {
	<-stopCh

	ex.lock.Lock()
	defer ex.lock.Unlock()

	if len(ex.waitingTasks) == 0 {
		ex.activeTasks--
		if ex.activeTasks == 0 {
			close(ex.finish)
		}
	} else {
		close(ex.waitingTasks[0])
		ex.waitingTasks = ex.waitingTasks[1:]
	}
}
