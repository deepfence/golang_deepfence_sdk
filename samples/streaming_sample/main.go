package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Invalid usage.")
		fmt.Println("Correct usage is:" + os.Args[0] + " <mgmt_console_ip> <port> <auth_key>")
		fmt.Println("e.g.:" + os.Args[0] +
			"134.211.211.211 443 default:6b5e9f7a-e3a2-402c-b9e3-41698366da05")
		return
	}
	SampleStreamer(os.Args[1], os.Args[2], os.Args[3])
}
