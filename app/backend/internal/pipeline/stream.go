package pipeline

import "context"

func Stream[J, R any](ctx context.Context, inputStream <-chan J, buffer int, fn func(context.Context, J) (R, error)) <-chan R {
	outchan := make(chan R, buffer)
	go func() {
		defer close(outchan)
		for job := range inputStream {
			select {
			case <-ctx.Done():
				return
			default:
			}

			res, err := fn(ctx, job)
			if err != nil {
				continue
			}

			select {
			case <-ctx.Done():
				return
			case outchan <- res:
			}

		}
	}()
	return outchan
}

func ProduceJob[J any](ctx context.Context, Job []J, buffer int) <-chan J {
	inchan := make(chan J, buffer)
	go func() {
		defer close(inchan)
		for _, job := range Job {
			select {
			case <-ctx.Done():
				return
			default:
			}

			select {
			case <-ctx.Done():
				return
			case inchan <- job:
			}

		}
	}()
	return inchan
}
