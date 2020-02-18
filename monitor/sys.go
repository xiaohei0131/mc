package monitor

func GetSysInfo() string {
	cmdRe := runCmd("uname -a")
	if len(cmdRe)>0{
		return cmdRe[0]
	}
	return ""
}
