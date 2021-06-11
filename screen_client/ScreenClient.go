package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

const (
	WBUFLEN int = 512
)

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "fatial err %v\r\n", err)
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "param err\r\n Usage:%s server ip:port\r\n", os.Args[0])
		os.Exit(1)
	}
	serverAddr, err := net.ResolveTCPAddr("tcp4", os.Args[1])
	checkErr(err)
	connTcp, err := net.DialTCP("tcp", nil, serverAddr)
	checkErr(err)
	cnt, err := connTcp.Write([]byte("xptd golang socket client test"))
	checkErr(err)
	fmt.Fprintf(os.Stdout, "%d bytes writen\r\n", cnt)
	buf := make([]byte, WBUFLEN)
	reader := bufio.NewReader(connTcp)
	ret := ""
	for {
		count, err := reader.Read(buf)
		if count == 0 {
			break
		}
		if err != nil && err != io.EOF {
			break
		}
		ret += string(buf[:count])
	}
	fmt.Fprintf(os.Stdout, "rev:%s\r\n", ret)
	os.Exit(1)
}
