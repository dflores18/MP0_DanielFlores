package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("To: ")
	text, _ := reader.ReadString('\n')

	reader1 := bufio.NewReader(os.Stdin)
	fmt.Print("From: ")
	text1, _ := reader1.ReadString('\n')

	reader2 := bufio.NewReader(os.Stdin)
	fmt.Print("Title: ")
	text2, _ := reader2.ReadString('\n')

	reader3 := bufio.NewReader(os.Stdin)
	fmt.Print("Content: ")
	text3, _ := reader3.ReadString('\n')

	newEmail := "To: " + text + "\n" + "From: " + text1 + "\n" + "Title: " + text2 + "\n" + "Content: " + text3 + "\n"

	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		io.WriteString(conn, fmt.Sprint(newEmail, time.Now(), "\n"))
		conn.Close()
		fmt.Println("TCP client exiting...")
		return
	}
}