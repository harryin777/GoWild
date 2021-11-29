/**
*   @Author: yky
*   @File: main
*   @Version: 1.0
*   @Date: 2021-07-14 21:56
 */
package main

import (
	"GoWild/common/ip"
	"GoWild/route"
)

func main() {
	r := route.Route()
	go ip.LocationInstances().ReadLocal()
	r.Run(":8080")
}
