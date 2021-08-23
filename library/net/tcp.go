package net

import (
	"fmt"
	"github.com/gogf/gf/net/gtcp"
)

//type TypeReq struct {
//	Ho
//
//}

func Tcp(host string) {
	gtcp.NewServer(host, func(conn *gtcp.Conn) {
		defer conn.Close()
		for {
			data, err := conn.Recv(-1)
			if len(data) > 0 {
				if err := conn.Send(append([]byte("> "), data...)); err != nil {
					fmt.Println(err)
				}
			}
			if err != nil {
				break
			}
		}
	}).Run()
}
