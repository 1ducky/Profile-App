package utils

func Retry(maxRetry int, fn func() error, delayMS int) error {
	for i := 0; i < maxRetry; i++ {
		if err := fn(); err != nil {
			return err
		}
		break
	}
	return nil
}
