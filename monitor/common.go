package monitor

import (
	"fmt"
	"os/exec"
	"strings"
)

func runCmd(cmdstr string)  []string{
	cmd := exec.Command("/bin/bash", "-c", cmdstr)
	out,err := cmd.Output()

	if err != nil{
		fmt.Println(err)
		return nil
	}else {
		result := strings.Split(string(out),"\n")
		return result
	}
}
