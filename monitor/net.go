package monitor

import (
	"log"
	"net"
)

func GetLocalIP(logger *log.Logger) string {
	defer func() {
		if err := recover(); err != nil {
			logger.Println("IP采集失败", err)
		}
	}()
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "unknown"
	}
	for _, addr := range addrs {
		ipAddr, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.IsLoopback() {
			continue
		}
		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}
		return ipAddr.IP.String()
	}
	return "unknown"
}