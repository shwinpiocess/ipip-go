package main

import (
	"fmt"
	"github.com/mxlxm/ipip-go"
)

func main() {
	ipip.Init("ipip.datx")
	if info, err := ipip.Find("8.8.8.8"); err == nil {
		fmt.Printf("%v\n", info)
	}
}
