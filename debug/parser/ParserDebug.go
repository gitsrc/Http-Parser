package main

import (
	"Parser"
	"fmt"
)

func main() {
	debug := Parser.GetNewParser()
	HttpRaw := "POST http://baidu.com HTTP/1.1\r\nAccept-Encoding: gzip,deflate,sdch\r\nHost: 50.23.145.133:20100\r\n\r\n12313132"
	err := debug.GetHttpStructFromRaw(HttpRaw)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(debug.Url)
		fmt.Println(string(debug.BodyBytes))
		fmt.Println(debug.RequestType)
	}
}
