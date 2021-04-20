package monitor

import (
	"mc/common"
)

func GetSysInfo() string {
	defer func() {
		if err := recover(); err != nil {
			common.MCLOG.Println("操作系统数据采集失败", err)
		}
	}()
	cmdRe := runCmd("uname -a")
	if len(cmdRe) > 0 {
		return cmdRe[0]
	}
	return ""
}
