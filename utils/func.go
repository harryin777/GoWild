package utils

import (
	"net"
	"sync"
)

var clientIPOnce sync.Once
var clientIP string

func GetDeviceClientIp() string {
	clientIPOnce.Do(func() {
		addrs, err := net.InterfaceAddrs()
		if err != nil {
			return
		}

		for _, address := range addrs {
			if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					clientIP = ipnet.IP.String()
					break
				}

			}
		}

	})
	return clientIP
}
