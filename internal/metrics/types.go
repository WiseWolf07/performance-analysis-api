package metrics

import (
	"github.com/shirou/gopsutil/load"
)

type SystemMetrics struct {
	CPUUsage    []float64     `json:"cpu_usage"`
	MemoryUsage *Mem          `json:"memory_usage"`
	LoadAvg     *load.AvgStat `json:"load_average"`
	Goroutines  int           `json:"goroutines_count"`
}

type Mem struct {
	Total       uint64  `json:"total"`
	Available   uint64  `json:"available"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"used_percent"`
	Free        uint64  `json:"free"`
}
