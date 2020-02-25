package monitor

import (
	"fmt"
	"strings"
)

type GPUInfo struct {
	Index          uint64    `json:"index"`
	Name           string    `json:"name"`
	DriverVersion  string    `json:"driver_version"`
	MemoryTotal    uint64    `json:"memory_total"`
	MemoryUsed     uint64    `json:"memory_used"`
	MemoryFree     uint64    `json:"memory_free"`
	MemoryUtil     float64   `json:"memory_utilization"`
	GpuUtil        float64   `json:"gpu_utilization"`
	GpuTemperature float64   `json:"gpu_temperature"`
	Processes      []Process `json:"processes"`
}

type Process struct {
	PID           uint64 `json:"pid"`
	Name          string `json:"name"`
	UsedGpuMemory uint64 `json:"used_gpu_memory"`
}

func GetGpuInfo() []GPUInfo {
	cmdRe := runCmd("nvidia-smi --query-gpu=index,name,driver_version,memory.total,memory.used,memory.free,utilization.gpu,utilization.memory,temperature.gpu --format=csv,noheader,nounits")
	gpus := []GPUInfo{}
	for _, v := range cmdRe {
		arr := strings.Split(v, ",")
		if len(arr) != 9 {
			continue
		}
		gpu := GPUInfo{}
		gpu.Index = convertToUnit(arr[0])
		gpu.Name = strings.Replace(arr[1], " ", "", -1)
		gpu.DriverVersion = strings.Replace(arr[2], " ", "", -1)
		gpu.MemoryTotal = convertToUnit(arr[3])
		gpu.MemoryUsed = convertToUnit(arr[4])
		gpu.MemoryFree = convertToUnit(arr[5])
		gpu.GpuUtil = convertToFloat(arr[6]) / 100
		gpu.MemoryUtil = convertToFloat(arr[7]) / 100
		gpu.GpuTemperature = convertToFloat(arr[8])
		gpu.Processes = getProcessInGpu(gpu.Index)
		gpus = append(gpus, gpu)
	}
	return gpus
}

/**
获取gpu运行程序
 */
func getProcessInGpu(gid uint64) []Process {
	cmdStr := fmt.Sprintf("nvidia-smi -i %d --query-compute-apps=pid,name,used_memory --format=csv,noheader,nounits", gid)
	cmdRe := runCmd(cmdStr)
	processes := []Process{}
	for _, v := range cmdRe {
		arr := strings.Split(v, ",")
		if len(arr) != 3 {
			continue
		}
		p := Process{}
		p.PID = convertToUnit(arr[0])
		p.Name = strings.Replace(arr[1], " ", "", -1)
		p.UsedGpuMemory = convertToUnit(arr[2])
		processes = append(processes, p)
	}
	return processes
}
