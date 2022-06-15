package main

import (
	"fmt"
	"github.com/insomniacslk/dhcp/dhcpv4/client4"
)

func main() {
	c := client4.NewClient()

	res, err := c.Exchange("en0")

	fmt.Println("moo")
	_, _ = res, err

}
