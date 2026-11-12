package pipeline

import "context"

func Collect[R any](ctx context.Context, stream <-chan R, buffer int) []R {
	list := make([]R, 0, buffer)
	for item := range stream {
		select {
		case <-ctx.Done():
			return list
		default:
		}
		list = append(list, item)
	}
	return list
}
