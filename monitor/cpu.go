package monitor
import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
	"time"
)

func CpuInfo() interface{} {
	mapInstances := map[string]interface{}{}
	res, err := cpu.Times(false) // false是展示全部总和 true是分别展示
	if err != nil {
		return mapInstances
	}
	if len(res) == 1 {
		// CPU使用率
		percent, _ := cpu.Percent(time.Second, false)
		mapInstances["user"]=res[0].User
		mapInstances["system"]=res[0].System
		mapInstances["idle"]=res[0].Idle
		mapInstances["nice"]=res[0].Nice
		mapInstances["iowait"]=res[0].Iowait
		mapInstances["irq"]=res[0].Irq
		mapInstances["softirq"]=res[0].Softirq
		mapInstances["percent"]= percent[0]/100
		mapInstances["load"]= getCpuLoad()
	}
	return mapInstances
}

// cpu info
func getCpuInfo() {
	cpuInfos, err := cpu.Info()
	if err != nil {
		fmt.Printf("get cpu info failed, err:%v", err)
	}
	for _, ci := range cpuInfos {
		fmt.Println(ci)
	}
	// CPU使用率
	for {
		percent, _ := cpu.Percent(time.Second, false)
		fmt.Printf("cpu percent:%v\n", percent)
	}
}

/**
cpu负载
 */
func getCpuLoad() *load.AvgStat{
	avgStat, _ := load.Avg()
	return avgStat
}
