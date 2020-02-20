package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func handle(src net.Conn) {
	dst, err := net.Dial("tcp", "cybersec.kz:80")
	if err != nil {
		fmt.Println("Couldn't dial")
		log.Fatalln(err)
	}
	defer dst.Close()

	go func() {
		if _, err := io.Copy(dst, src); err != nil {
			fmt.Println("Couldn't copy forward")
			log.Fatalln(err)
		}
	}()

	if _, err := io.Copy(src, dst); err != nil {
		fmt.Println("Couldn't copy backward")
		log.Fatalln(err)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Couldn't listen")
		log.Fatalln(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Couldn't accept")
			log.Fatalln(err)
		}
		go handle(conn)
	}
}
