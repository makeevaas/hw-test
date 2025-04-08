package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	var wg sync.WaitGroup
	var mu sync.Mutex

	var errCount int32
	taskChan := make(chan Task)

	// исполнитель
	worker := func() {
		defer wg.Done()
		for task := range taskChan {
			mu.Lock()
			if int(errCount) >= m {
				mu.Unlock()
				return
			}
			mu.Unlock()
			if err := task(); err != nil {
				mu.Lock()
				errCount++
				mu.Unlock()
			}
		}
	}

	for i := 0; i < n; i++ {
		wg.Add(1)
		go worker()
	}

	// очередь задач кладем в канал, следим за количеством ошибок
	for _, task := range tasks {
		mu.Lock()
		if int(errCount) >= m {
			mu.Unlock()
			break
		}
		mu.Unlock()
		taskChan <- task
	}
	close(taskChan)

	// ожидание завершения всех горутин
	wg.Wait()

	mu.Lock()
	defer mu.Unlock()
	if int(errCount) >= m {
		return ErrErrorsLimitExceeded
	}

	return nil
}
