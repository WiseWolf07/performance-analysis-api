package metrics

import (
	"github.com/shirou/gopsutil/load"
)

type DiskMetrics struct {
	ReadBytes  uint64 `json:"read_bytes"`
	WriteBytes uint64 `json:"write_bytes"`
	ReadOps    uint64 `json:"read_ops"`
	WriteOps   uint64 `json:"write_ops"`
	Iops       uint64 `json:"iops"`
}

type SystemMetrics struct {
	CPUUsage    []float64     `json:"cpu_usage"`
	MemoryUsage *Mem          `json:"memory_usage"`
	LoadAvg     *load.AvgStat `json:"load_average"`
	DiskIO      *DiskMetrics  `json:"disk_io"`
	Goroutines  int           `json:"goroutines_count"`
	LatencyMs   float64       `json:"latency_ms"`
}

type Mem struct {
	Total       uint64  `json:"total"`
	Available   uint64  `json:"available"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"used_percent"`
	Free        uint64  `json:"free"`
}
