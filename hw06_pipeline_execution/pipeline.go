package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	pipe := in

	consumer := func(in In) {
		for range in {
			continue
		}
	}

	worker := func(in In, done In) Out {
		out := make(Bi)
		go func() {
			defer close(out)
			for {
				select {
				case <-done:
					go consumer(in)
					return
				case v, ok := <-in:
					if !ok {
						return
					}
					out <- v
				}
			}
		}()
		return out
	}

	for _, stage := range stages {
		pipe = stage(worker(pipe, done))
	}

	return pipe
}
