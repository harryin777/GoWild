/**
*   @Author: yky
*   @File: main
*   @Version: 1.0
*   @Date: 2021-07-14 21:56
 */
package main

import (
	"GoWild/route"
)

func main(){
	r := route.Route()

	r.Run(":8080")
}
