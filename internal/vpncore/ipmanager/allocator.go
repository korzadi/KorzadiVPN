package ipmanager

import "fmt"

var nextIP = 2

func AllocateIP() string {

	ip := fmt.Sprintf(
		"10.0.0.%d/32",
		nextIP,
	)

	nextIP++

	return ip
}
