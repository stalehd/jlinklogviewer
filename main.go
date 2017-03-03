package main

import (
	"fmt"
	"net"
	"time"
)

const endpoint = "127.0.0.1:19021"

func main() {
	fmt.Printf("JLinkLogViewer reading from %s\n", endpoint)
	for {
		jlinkExeAddr, err := net.ResolveTCPAddr("tcp4", endpoint)
		if err != nil {
			<-time.After(2 * time.Second)
			continue
		}

		conn, err := net.DialTCP("tcp4", nil, jlinkExeAddr)
		if err != nil {
			<-time.After(2 * time.Second)
			continue
		}

		data := make([]byte, 1024)
		fmt.Printf("\n\n===============================================================================\n")
		for err == nil {
			n, err := conn.Read(data)
			if err != nil {
				conn.Close()
				<-time.After(2 * time.Second)
				break
			}
			fmt.Printf("%s", string(data[0:n]))
		}
	}
}
