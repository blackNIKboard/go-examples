package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

func main() {
	var (
		repeat int
		times  []time.Duration
		host   string
		port   int
	)

	//host = "8.8.8.8"
	host = "194.226.199.82"
	port = 33434
	port = 14567
	repeat = 4

	for i := 0; i < repeat; i++ {
		times = append(times, ping(i, host, port))

		time.Sleep(time.Second)
	}

	fmt.Printf("---- Stats for %s:%d ----\n", host, port)
	var sum int64
	for _, duration := range times {
		sum += duration.Microseconds()
	}
	fmt.Printf("avg time: %.2f, transmitted %d packages\n", float64(sum)/float64(len(times))/1000, repeat)
}

func sendUDP(host string, port int) {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{Port: 1234})
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	dst := &net.UDPAddr{
		IP:   net.ParseIP(host),
		Port: port,
		//Port: 14567,
	}

	sendHello(conn, dst)
}

func ping(seq int, host string, port int) time.Duration {
	var (
		sendTime    time.Time
		receiveTime time.Time
	)

	listenPacket, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		log.Fatalf("listen err, %s", err)
	}
	defer listenPacket.Close()

	sendTime = time.Now()

	sendUDP(host, port)

	rb := make([]byte, 1500)
	n, _, err := listenPacket.ReadFrom(rb)
	if err != nil {
		log.Fatal(err)
	}

	receiveTime = time.Now()

	rm, err := icmp.ParseMessage(ipv4.ICMPTypeEchoReply.Protocol(), rb[:n])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Host replied with %v in %dms, seq %d\n", rm.Type, receiveTime.Sub(sendTime).Milliseconds(), seq)

	return receiveTime.Sub(sendTime)
}

func sendHello(conn *net.UDPConn, addr *net.UDPAddr) {
	n, err := conn.WriteTo([]byte("hello"), addr)
	if err != nil {
		log.Fatal("Write:", err)
	}
	fmt.Println("Sent", n, "bytes", conn.LocalAddr(), "->", addr)
}
