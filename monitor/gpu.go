package monitor

import (
	"fmt"
	"strings"
)

type GPUInfo struct {
	Index          string    `json:"index"`
	Name           string    `json:"name"`
	DriverVersion  string    `json:"driver_version"`
	MemoryTotal    string    `json:"memory_total"`
	MemoryUsed     string    `json:"memory_used"`
	MemoryFree     string    `json:"memory_free"`
	MemoryUtil     string    `json:"memory_utilization"`
	GpuUtil        string    `json:"gpu_utilization"`
	GpuTemperature string    `json:"gpu_temperature"`
	Processes      []Process `json:"processes"`
}

type Process struct {
	PID           string `json:"index"`
	Name          string `json:"name"`
	UsedGpuMemory string `json:"used_gpu_memory"`
}

func GetGpuInfo() []GPUInfo {
	cmdRe := runCmd("nvidia-smi --query-gpu=index,name,driver_version,memory.total,memory.used,memory.free,utilization.gpu,utilization.memory,temperature.gpu --format=csv")
	gpus := []GPUInfo{}
	for k, v := range cmdRe {
		if k == 0 {
			continue
		}
		arr := strings.Split(v,",")
		if len(arr)!=9{
			continue
		}
		gpu := GPUInfo{}
		gpu.Index = arr[0]
		gpu.Name = arr[1]
		gpu.DriverVersion = arr[2]
		gpu.MemoryTotal = arr[3]
		gpu.MemoryUsed = arr[4]
		gpu.MemoryFree = arr[5]
		gpu.GpuUtil = arr[6]
		gpu.MemoryUtil = arr[7]
		gpu.GpuTemperature = fmt.Sprint(arr[8], "℃")
		gpu.Processes = getProcessInGpu(gpu.Index)
		gpus = append(gpus, gpu)
	}
	return gpus
}

/**
获取gpu运行程序
 */
func getProcessInGpu(gid string) []Process {
	cmdStr := fmt.Sprintf("nvidia-smi -i %s --query-compute-apps=pid,name,used_memory --format=csv", gid)
	cmdRe := runCmd(cmdStr)
	processes := []Process{}
	for k, v := range cmdRe {
		if k == 0 {
			continue
		}
		arr := strings.Split(v,",")
		if len(arr)!=3{
			continue
		}
		p := Process{}
		p.PID = arr[0]
		p.Name = arr[1]
		p.UsedGpuMemory = arr[2]
		processes = append(processes, p)
	}
	return processes
}
