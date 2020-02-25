package monitor

import "log"

func GetSysInfo(logger *log.Logger) string {
	defer func() {
		if err := recover(); err != nil {
			logger.Println("操作系统数据采集失败", err)
		}
	}()
	cmdRe := runCmd("uname -a")
	if len(cmdRe)>0{
		return cmdRe[0]
	}
	return ""
}
