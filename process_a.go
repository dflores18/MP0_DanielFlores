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
	//took in command line/user input for each field
	//i.e To: , From:, Title: , Content:
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

	//concatenate the email/message fields to be sent through TCP channel as one message
	newEmail := "To: " + text + "\n" + "From: " + text1 + "\n" + "Title: " + text2 + "\n" + "Content: " + text3 + "\n"

	ln, err := net.Listen("tcp", ":9000") //the server is listening
	if err != nil {
		//handle an error
		panic(err)
	}
	defer ln.Close() //close the listener

	for {
		conn, err := ln.Accept() //Accept function returns reader and writer
		if err != nil {
			panic(err)
		}
		io.WriteString(conn, fmt.Sprint(newEmail, time.Now(), "\n")) //writes the contents of the string (fmt.Sprint) to (conn), which accepts a slice of bytes
		conn.Close()                                                 //closes the connect between server and client
		fmt.Println("TCP client exiting...")
		return
	}
}
