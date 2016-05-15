package main

import (
	"fmt"
	"github.com/mxlxm/ipip-go"
)

func main() {
	d, e := ipip.Init("ipip.datx")
	if e != nil {
		panic(e)
	}
	if info, err := d.Find("8.8.8.8"); err == nil {
		fmt.Printf("%+v\n", info)
	}
}
