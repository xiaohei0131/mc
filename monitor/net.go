package monitor

import (
	"mc/common"
	"net"
)

/**
此方法遇到多个IP接口时会出现不确定性
 */
func GetLocalIP() string {
	defer func() {
		if err := recover(); err != nil {
			common.MCLOG.Println("IP采集失败", err)
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

func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		common.MCLOG.Println("IP采集失败", err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}