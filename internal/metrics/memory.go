package metrics

import (
	"fmt"
	"runtime"

	"github.com/shirou/gopsutil/mem"
)

func GetSystemMemoryUsage() (*Mem, error) {
	virtualMemory, err := mem.VirtualMemory()
	if err != nil {
		return nil, fmt.Errorf("error al obtener la memoria vitual: %w", err)
	}

	memInfo := &Mem{
		Total:       virtualMemory.Total,
		Available:   virtualMemory.Available,
		Used:        virtualMemory.Used,
		UsedPercent: virtualMemory.UsedPercent,
		Free:        virtualMemory.Free,
	}

	return memInfo, nil
}

func GetGoAppMemoryUsage() uint64 {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	return memStats.Alloc
}
