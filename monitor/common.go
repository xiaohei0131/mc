package monitor

import (
	"context"
	"errors"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func runCmdBlock(cmdstr string) []string {
	cmd := exec.Command("/bin/bash", "-c", cmdstr)
	out, err := cmd.Output()

	if err != nil {
		panic(err)
	} else {
		result := strings.Split(string(out), "\n")
		return result
	}
}

/**
执行命令10秒超时
 */
func runCmd(cmdstr string) []string {
	ctxt, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctxt, "/bin/bash", "-c", cmdstr)
	//当经过Timeout时间后，程序依然没有运行完，则会杀掉进程，ctxt也会有err信息
	if out, err := cmd.Output(); err != nil {
		//检测报错是否是因为超时引起的
		if ctxt.Err() != nil && ctxt.Err() == context.DeadlineExceeded {
			panic(errors.New("command timeout"))
		}else{
			panic(err)
		}
	} else {
		result := strings.Split(string(out), "\n")
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
	ns := strings.TrimSpace(ov)
	if (strings.HasSuffix(ov, "%")) {
		ns = ns[:len(ns)-1]
	}
	i, _ := strconv.ParseFloat(ns, 64)
	return i
}
