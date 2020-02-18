package monitor

import (
	//"github.com/shirou/gopsutil/disk"
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
func DiskMonitor()  interface{}{
	cmdRe := runCmd("df -hT")
	mapInstances := []map[string]interface{}{}
	for k,v:= range cmdRe{
		if k==0 || v == ""{
			continue
		}else {
			arr := strings.Fields(v)
			instance := map[string]interface{}{}
			instance["device"] = arr[0]
			instance["fstype"] = arr[1]
			instance["total"] = arr[2]
			instance["used"] = arr[3]
			instance["available"] = arr[4]
			instance["usedPercent"] = arr[5]
			instance["mountpoint"] = arr[6]
			mapInstances = append(mapInstances, instance)
		}
	}
	return mapInstances
}