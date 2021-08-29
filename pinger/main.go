package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"sync"
	"time"

	"github.com/go-ping/ping"
)

type (
	target struct {
		ip     string
		online bool
	}
)

func pingfunc(wg *sync.WaitGroup, host *target) {
	defer wg.Done()

	pinger, err := ping.NewPinger(host.ip)
	if err != nil {
		fmt.Println(err.Error())
		// panic(err)
	}

	pinger.SetPrivileged(true)

	pinger.Count = 3

	pinger.OnRecv = func(stats *ping.Packet) {
		fmt.Print("received: ")
		host.online = true
		fmt.Println(stats)
	}

	pinger.Timeout = time.Second * 5

	err = pinger.Run()
	if err != nil {
		return
	}
}

func checkStatus(ips []*target) {
	var wg sync.WaitGroup

	for i := 0; i < len(ips); i++ {
		wg.Add(1)
		go pingfunc(&wg, ips[i])
	}

	wg.Wait()
}

// This requires sudo privileges in order to run pings
func main() {
	ips := []*target{
		{"8.8.8.8", false},
		{"0.167.1.101", false},
		{"localhost", false},
	}

	checkStatus(ips)

	spew.Dump(ips)

}
