package main

import (
	"fmt"
	"io/ioutil"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9000") //dials/connects to the same server (tcp)
	if err != nil {
		panic(err)
	} //handles error
	defer conn.Close() //closes connection

	bs, _ := ioutil.ReadAll(conn)
	fmt.Println(string(bs))

}
