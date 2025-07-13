package metrics

import (
	"fmt"
	"sync"

	"github.com/shirou/gopsutil/load"
)

func GetSystemMetrics() SystemMetrics {
	var waitGroup sync.WaitGroup
	metrics := SystemMetrics{}

	errChan := make(chan error, 4)

	waitGroup.Add(4)

	go func() {
		defer waitGroup.Done()
		cpuUsage, err := GetCPUUsage()
		if err != nil {
			errChan <- fmt.Errorf("error al obtener uso de CPU %w", err)
			return
		}
		metrics.CPUUsage = cpuUsage
	}()

	go func() {
		defer waitGroup.Done()
		memUsage, err := GetSystemMemoryUsage()
		if err != nil {
			errChan <- fmt.Errorf("error al obtener el uso de memoria %w", err)
			return
		}
		metrics.MemoryUsage = memUsage
	}()

	go func() {
		defer waitGroup.Done()
		loadAvg, err := load.Avg()
		if err != nil {
			errChan <- fmt.Errorf("error al obtener carga promedio: %w", err)
			return
		}
		metrics.LoadAvg = loadAvg
	}()

	go func() {
		defer waitGroup.Done()
		diskIO, err := GetDiskIOUsage()
		if err != nil {
			errChan <- fmt.Errorf("error al obtener métricas de E/S de disco: %w", err)
			return
		}
		metrics.DiskIO = diskIO
	}()

	metrics.Goroutines = GetGoroutinesCount()

	waitGroup.Wait()
	close(errChan)

	for err := range errChan {
		fmt.Println("Error de métrica: ", err)
	}

	return metrics
}
