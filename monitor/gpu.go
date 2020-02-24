package monitor

import (
	"fmt"
	"strconv"
	"strings"
)

type GPUInfo struct {
	Index          string    `json:"index"`
	Name           string    `json:"name"`
	DriverVersion  string    `json:"driver_version"`
	MemoryTotal    int    `json:"memory_total"`
	MemoryUsed     int    `json:"memory_used"`
	MemoryFree     int    `json:"memory_free"`
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
		gpu.Name = strings.Replace(arr[1], " ", "", -1)
		gpu.DriverVersion = strings.Replace(arr[2], " ", "", -1)
		gpu.MemoryTotal = getMibValue(arr[3])
		gpu.MemoryUsed = getMibValue(arr[4])
		gpu.MemoryFree = getMibValue(arr[5])
		gpu.GpuUtil = strings.Replace(arr[6], " ", "", -1)
		gpu.MemoryUtil = strings.Replace(arr[7], " ", "", -1)
		gpu.GpuTemperature = fmt.Sprint(strings.Replace(arr[8], " ", "", -1), "℃")
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
		p.Name = strings.Replace(arr[1], " ", "", -1)
		p.UsedGpuMemory = strings.Replace(arr[2], " ", "", -1)
		processes = append(processes, p)
	}
	return processes
}

func getMibValue(ov string) int  {
	arr := strings.Split(strings.TrimSpace(ov)," ")
	if len(arr) == 2{
		i, _ := strconv.Atoi(arr[0])
		return i
	}
	return 0
}