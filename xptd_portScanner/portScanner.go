// v0.1 test tcp connt
// package main

// import (
// 	"fmt"
// 	"net"
// 	"os"
// )

// func checkErr(err error) {
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "fatal err:%v\r\n", err)
// 		os.Exit(1)
// 	}
// }
// func main() {
// 	fmt.Fprintf(os.Stdout, "hello xptd!\r\n")
// 	var ipAddr string = "192.168.2.19"

// 	for i := 20; i < 1023; i++ {
// 		target := fmt.Sprintf("%s:%d", ipAddr, i)
// 		conn, err := net.Dial("tcp", target)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "port:%d closed\r\n", i)
// 			continue
// 		}
// 		fmt.Fprintf(os.Stdout, "port:%d opened\r\n", i)
// 		defer func() {
// 			err := conn.Close()
// 			checkErr(err)
// 		}()
// 	}
// }

// v0.2 test ping
// package main

// import (
// 	"fmt"
// 	"net"
// 	"os"
// 	"time"

// 	"github.com/go-ping/ping"
// )

// const (
// 	helpInfo  string = `Usage:xptdPing.exe ipAddr`
// 	ICMPCOUNT int    = 30
// 	PINGTIME  int64  = 200
// )

// func checkErr(err error) {
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "fatal err:\r\n%s\r\n", err.Error())
// 		os.Exit(1)
// 	}
// }

// func pingTest(target string) bool {
// 	ip, err := net.ResolveIPAddr("ip", target)
// 	checkErr(err)
// 	fmt.Fprintf(os.Stdout, "ip:%v\r\n", ip)
// 	pinger, err := ping.NewPinger(os.Args[1])
// 	checkErr(err)
// 	pinger.Count = ICMPCOUNT
// 	pinger.SetPrivileged(true)
// 	pinger.Timeout = time.Duration(500 * time.Millisecond)

// 	pinger.OnRecv = func(pkt *ping.Packet) {
// 		fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v\n",
// 			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt)
// 	}

// 	pinger.OnDuplicateRecv = func(pkt *ping.Packet) {
// 		fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v ttl=%v (DUP!)\n",
// 			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt, pkt.Ttl)
// 	}

// 	pinger.OnFinish = func(stats *ping.Statistics) {
// 		fmt.Printf("\n--- %s ping statistics ---\n", stats.Addr)
// 		fmt.Printf("%d packets transmitted, %d packets received, %v%% packet loss\n",
// 			stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
// 		fmt.Printf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
// 			stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
// 	}
// 	err = pinger.Run()
// 	checkErr(err)
// 	stat := pinger.Statistics()
// 	if stat.PacketsRecv >= 1 {
// 		return true
// 	}
// 	return false
// }

// func main() {
// 	fmt.Fprintf(os.Stdout, "hell xptd test ping\r\n")
// 	if len(os.Args) < 2 {
// 		fmt.Fprintf(os.Stderr, "param err:\r\n%s\r\n", helpInfo)
// 		os.Exit(1)
// 	}
// 	if pingTest(os.Args[1]) {
// 		fmt.Fprintf(os.Stdout, "ok\r\n")
// 	} else {
// 		fmt.Fprintf(os.Stdout, "err\r\n")
// 	}
// 	os.Exit(0)
// }

// //v0.3 ip format
// package main

// import (
// 	"fmt"
// 	"net"
// 	"os"
// )

// func checkErr(err error) {
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "fatal err\r\n%v\r\n", err)
// 		os.Exit(1)
// 	}
// }

// func getIP(input string) (retIp net.IP) {
// 	if input == "" {
// 		return nil
// 	}
// 	ip, err := net.ResolveIPAddr("ip", input)
// 	checkErr(err)
// 	ret := net.ParseIP(ip.IP.String())
// 	if ret == nil {
// 		fmt.Fprintf(os.Stderr, "input err")
// 		os.Exit(1)
// 	}
// 	retIp = ret
// 	return retIp
// }

// func main() {
// 	fmt.Fprintf(os.Stdout, "xptdIPTest\r\n")

// 	if len(os.Args) < 2 {
// 		fmt.Fprintf(os.Stderr, "param err!\r\nUsage :%s ipAddr\r\n", os.Args[0])
// 		os.Exit(1)
// 	}
// 	ip := getIP(os.Args[1])
// 	if ip != nil {
// 		fmt.Fprintf(os.Stdout, "ip:%v", ip)
// 		os.Exit(0)
// 	} else {
// 		fmt.Fprintf(os.Stderr, "ip err\r\n")
// 		os.Exit(0)
// 	}

// }

//https://github.com/go-ping/ping/blob/master/cmd/ping/ping.go
//https://cloud.tencent.com/developer/article/1778187

//v.04