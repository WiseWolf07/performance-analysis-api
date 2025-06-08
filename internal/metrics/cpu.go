package metrics

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

func GetCPUUsage() ([]float64, error) {
	cpuPercent, err := cpu.Percent(time.Second, true)
	if err != nil {
		return nil, fmt.Errorf("Error al obtener el uso de CPU: %w", err)
	}
	return cpuPercent, nil
}
