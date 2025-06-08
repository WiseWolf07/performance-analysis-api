package metrics

import "runtime"

func GetGoroutinesCount() int {
	return runtime.NumGoroutine()
}
