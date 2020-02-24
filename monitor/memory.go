package monitor

import (
	"github.com/shirou/gopsutil/mem"
	"strings"
)

type Memory struct {
	Memtotal        uint64  `json:"memory_total"`
	Memused         uint64  `json:"memory_used"`
	MemUsedPercent  float64 `json:"memory_utilization"`
	MemAvailable    uint64  `json:"memory_available"`
	MemFree         uint64  `json:"memory_free"`
	SwapTotal       uint64  `json:"swap_total"`
	SwapFree        uint64  `json:"swap_free"`
	SwapUsed        uint64  `json:"swap_used"`
	SwapUsedPercent float64 `json:"swap_utilization"`
}

func MemInfo() interface{} {
	v, _ := mem.VirtualMemory()
	memory := Memory{}
	memory.Memtotal = v.Total / (1024 * 1024)
	memory.Memused = v.Used / (1024 * 1024)
	memory.MemUsedPercent = v.UsedPercent
	memory.MemAvailable = v.Available / (1024 * 1024)
	memory.MemFree = v.Free / (1024 * 1024)
	sw, _ := mem.SwapMemory()
	memory.SwapFree = sw.Free / (1024 * 1024)
	memory.SwapTotal = sw.Total / (1024 * 1024)
	memory.SwapUsed = sw.Used / (1024 * 1024)
	memory.SwapUsedPercent = sw.UsedPercent
	return memory
}

/**
监控内存
*/
func memoryMonitor() interface{} {
	cmdRe := runCmd("free -h")
	mapInstances := map[string]interface{}{}
	for k, v := range cmdRe {
		if k == 0 {
			continue
		} else {
			arr := strings.Fields(v)
			if k == 1 {
				mapInstances["memTotal"] = arr[1]
				mapInstances["memUsed"] = arr[2]
				mapInstances["memFree"] = arr[3]
				mapInstances["memShared"] = arr[4]
				mapInstances["memBuffOrCache"] = arr[5]
				mapInstances["memAvailable"] = arr[6]
			} else if k == 2 {
				mapInstances["swapTotal"] = arr[1]
				mapInstances["swapUsed"] = arr[2]
				mapInstances["swapFree"] = arr[3]
			} else {
				break
			}
		}
	}
	return mapInstances
}
