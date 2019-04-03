package main

import (
	"io"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil{
		log.Fatal(err)
	}
	defer lis.Close()

	for{
		conn, err := lis.Accept()
		if err != nil{
			log.Println(err)
			continue
		}

		//write to an open connection
		io.WriteString(conn, "Connection Established")

		conn.Close()
	}
}
