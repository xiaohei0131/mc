package monitor

import (
	"fmt"
	"github.com/shirou/gopsutil/mem"
	"strings"
)

func MemInfo() interface{} {
	v, _ := mem.VirtualMemory()
	mem := map[string]interface{}{}
	/*mem["total"] = fmt.Sprint(v.Total/(1024*1024),"M")
	mem["used"] = fmt.Sprint(v.Used/(1024*1024),"M")
	mp := fmt.Sprintf("%.2f",v.UsedPercent)
	mem["usedPercent"] = fmt.Sprint(mp,"%")
	mem["available"] = fmt.Sprint(v.Available/(1024*1024),"M")
	mem["free"] = fmt.Sprint(v.Free/(1024*1024),"M")*/
	mem["total"] = v.Total / (1024 * 1024)
	mem["used"] = v.Used / (1024 * 1024)
	mp := fmt.Sprintf("%.2f", v.UsedPercent)
	mem["usedPercent"] = fmt.Sprint(mp, "%")
	mem["available"] = v.Available / (1024 * 1024)
	mem["free"] = v.Free / (1024 * 1024)
	return mem
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
