package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main(){
	li, err := net.Listen("tcp", ":8080")
	if err != nil{
		log.Panic(err)
	}
	for{
		conn, err := li.Accept()
		if err != nil{
			log.Println(err)
		}
		go handle(conn)
	}
}
//1EYDHYLAHY1E
//1EYDHYLAHY1HHggg
//BONNA8808DDDD
func handle(conn net.Conn){
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil{
		log.Println("CONN TIMEOUT")
	}

	scanner := bufio.NewScanner(conn)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan(){
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard you said: %s\n", ln)
	}


	defer conn.Close()
	fmt.Println("*** CODE CAN GET HERE ***")
}