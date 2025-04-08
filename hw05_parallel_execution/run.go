package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	var wg sync.WaitGroup
	var errCount int32
	taskChan := make(chan Task)

	// исполнитель
	worker := func() {
		defer wg.Done()
		for task := range taskChan {
			if int(errCount) >= m {
				return
			}
			if err := task(); err != nil {
				atomic.AddInt32(&errCount, 1)
			}
		}
	}

	for i := 0; i < n; i++ {
		wg.Add(1)
		go worker()
	}

	// очередь задач кладем в канал, следим за количеством ошибок
	for _, task := range tasks {
		if int(errCount) >= m {
			break
		}
		taskChan <- task
	}
	close(taskChan)

	// ожидание завершения всех горутин
	wg.Wait()

	if int(errCount) >= m {
		return ErrErrorsLimitExceeded
	}

	return nil
}
