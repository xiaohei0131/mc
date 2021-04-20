package monitor

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
	"mc/common"
	"time"
)

func CpuInfo() interface{} {
	defer func() {
		if err := recover(); err != nil {
			common.MCLOG.Println("CPU采集失败", err)
		}
	}()
	mapInstances := map[string]interface{}{}
	res, err := cpu.Times(false) // false是展示全部总和 true是分别展示
	if err != nil {
		return mapInstances
	}
	if len(res) == 1 {
		// CPU使用率
		percent, _ := cpu.Percent(time.Second, false)
		mapInstances["user"]=res[0].User/res[0].Total()
		mapInstances["system"]=res[0].System/res[0].Total()
		mapInstances["idle"]=res[0].Idle/res[0].Total()
		mapInstances["nice"]=res[0].Nice/res[0].Total()
		mapInstances["iowait"]=res[0].Iowait/res[0].Total()
		mapInstances["irq"]=res[0].Irq/res[0].Total()
		mapInstances["softirq"]=res[0].Softirq/res[0].Total()
		mapInstances["steal"]=res[0].Steal/res[0].Total()
		mapInstances["percent"]= percent[0]/100
		mapInstances["load"]= getCpuLoad()
	}
	return mapInstances
}

// cpu info
func getCpuInfo() {
	cpuInfos, err := cpu.Info()
	if err != nil {
		common.MCLOG.Printf("get cpu info failed, err:%v", err)
	}
	for _, ci := range cpuInfos {
		common.MCLOG.Println(ci)
	}
	// CPU使用率
	for {
		percent, _ := cpu.Percent(time.Second, false)
		common.MCLOG.Printf("cpu percent:%v\n", percent)
	}
}

/**
cpu负载
 */
func getCpuLoad() *load.AvgStat{
	avgStat, _ := load.Avg()
	return avgStat
}
