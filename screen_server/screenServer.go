// package main

// import (
// 	"fmt"
// 	"net"
// 	"os"
// )

// func main() {
// 	if len(os.Args) < 2 {
// 		fmt.Fprintf(os.Stderr, "param err:Usage:%s ip addr\r\n", os.Args[0])
// 		os.Exit(1)
// 	}
// 	name := os.Args[1]
// 	addr := net.ParseIP(name)
// 	if addr == nil {
// 		fmt.Fprintf(os.Stderr, "%s:ip addr param err\r\n", os.Args[1])
// 		os.Exit(1)
// 	} else {
// 		fmt.Fprintf(os.Stdout, "ip addr:%v\r\n", addr)
// 	}

// }

// package main

// import (
// 	"fmt"
// 	"net"
// 	"os"
// )

// func main() {
// 	if len(os.Args) < 2 {
// 		fmt.Fprintf(os.Stderr, "param err:%s tcp_addr:tcp_port\r\n", os.Args[0])
// 		os.Exit(1)
// 	}
// 	tcpAddr, err := net.ResolveTCPAddr("tcp", os.Args[1])
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "tcp_addr resolve err\r\n")
// 		os.Exit(1)
// 	} else {
// 		fmt.Fprintf(os.Stdout, "tcpaddr:%v,tcpport:%v\r\n", tcpAddr.IP, tcpAddr.Port)
// 	}

// }

package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "fatail err:%s\r\n", err.Error())
		os.Exit(1)
	}
}
func connWorker(con net.Conn) {
	defer func() {
		if err := con.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "fatail err:%s\r\n", err.Error())
			os.Exit(1)
		}
	}()
	now := time.Now().String()
	con.Write([]byte(now))
	con.Close()
}
func main() {
	service := ":12345"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkErr(err)
	tcpListener, err := net.ListenTCP("tcp", tcpAddr)
	checkErr(err)
	for {
		connect, err := tcpListener.Accept()
		if err != nil {
			continue
		}
		connWorker(connect)
	}

}
