package monitor

import (
	"log"
	"strings"
)

// disk info
/*func GetDiskInfo() {
	parts, err := disk.Partitions(true)
	if err != nil {
		fmt.Printf("get Partitions failed, err:%v\n", err)
		return
	}
	for _, part := range parts {
		fmt.Printf("part:%v\n", part.String())
		diskInfo, _ := disk.Usage(part.Mountpoint)
		fmt.Printf("disk info:used:%v free:%v\n", diskInfo.UsedPercent, diskInfo.Free)
	}

	ioStat, _ := disk.IOCounters()
	for k, v := range ioStat {
		fmt.Printf("%v:%v\n", k, v)
	}
}*/

/**
监控磁盘
*/
func DiskMonitor(logger *log.Logger) interface{} {
	defer func() {
		if err := recover(); err != nil {
			logger.Println("磁盘（disk）采集失败", err)
		}
	}()
	//cmdRe := runCmd("df -hT")
	cmdRe := runCmd("df -mT")
	mapInstances := []map[string]interface{}{}
	for k, v := range cmdRe {
		if k == 0 || v == "" {
			continue
		} else {
			arr := strings.Fields(v)
			instance := map[string]interface{}{}
			instance["device"] = arr[0]
			instance["fstype"] = arr[1]
			instance["total"] = convertToUnit(arr[2])
			instance["used"] = convertToUnit(arr[3])
			instance["available"] = convertToUnit(arr[4])
			instance["usedPercent"] = parsePercent(arr[5]) / 100
			instance["mountpoint"] = arr[6]
			mapInstances = append(mapInstances, instance)
		}
	}
	return mapInstances
}
