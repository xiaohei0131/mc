package common

import (
	"log"
	"os"
)

var MCLOG *log.Logger

func init() {
	file, err := os.OpenFile("mc.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("fail to create mc.log file!")
	}
	MCLOG = log.New(file, "", log.Llongfile)
	MCLOG.SetFlags(log.LstdFlags)
}
