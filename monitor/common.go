package monitor

import (
	"os/exec"
	"strconv"
	"strings"
)

func runCmd(cmdstr string)  []string{
	cmd := exec.Command("/bin/bash", "-c", cmdstr)
	out,err := cmd.Output()

	if err != nil{
		panic(err)
	}else {
		result := strings.Split(string(out),"\n")
		return result
	}
}

func convertToUnit(ov string) uint64 {
	i, _ := strconv.ParseUint(strings.TrimSpace(ov), 10, 64)
	return i
}

func convertToFloat(ov string) float64 {
	i, _ := strconv.ParseFloat(strings.TrimSpace(ov), 64)
	return i
}

func parsePercent(ov string) float64 {
	ns:=strings.TrimSpace(ov)
	if(strings.HasSuffix(ov,"%")){
		ns = ns[:len(ns)-1]
	}
	i, _ := strconv.ParseFloat(ns, 64)
	return i
}
