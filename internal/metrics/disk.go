package metrics

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/disk"
)

func GetDiskIOUsage() (*DiskMetrics, error) {
	initialCounters, err := disk.IOCounters()
	if err != nil {
		return nil, fmt.Errorf("error al obtener contadores iniciales del disco: %w", err)
	}

	time.Sleep(1 * time.Second)

	finalCounters, err := disk.IOCounters()
	if err != nil {
		return nil, fmt.Errorf("error al obtener los contadores finales del disco: %w", err)
	}

	var readBytesDiff, writeBytesDiff, readOpsDiff, writeOpsDiff uint64
	for _, initial := range initialCounters {
		for _, final := range finalCounters {
			if initial.Name == final.Name {
				readBytesDiff += final.ReadBytes - initial.ReadBytes
				writeBytesDiff += final.WriteBytes - initial.WriteBytes
				readOpsDiff += final.ReadCount - initial.ReadCount
				writeOpsDiff += final.WriteCount - initial.WriteCount
				break
			}
		}
	}

	totalIops := readOpsDiff + writeOpsDiff

	return &DiskMetrics{
		ReadBytes:  readBytesDiff,
		WriteBytes: writeBytesDiff,
		ReadOps:    readOpsDiff,
		WriteOps:   writeOpsDiff,
		Iops:       totalIops,
	}, nil
}
