package modules

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"

	"github.com/ehsanghaffar/ultimate-internet-test/utils"
	"github.com/go-ping/ping"
)

func PingCheck(domain string) {
	pinger, err := ping.NewPinger(domain)
	if err != nil {
		panic(err)
	}
	// Listen for Ctrl-C.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			pinger.Stop()
		}
	}()

	pinger.OnRecv = func(pkt *ping.Packet) {
		log.Printf("%d bytes from %s: icmp_seq=%d time=%v\n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt)

		// TODO: find best practice for this
		if pkt.Seq >= 5 {
			pinger.Stop()
		}
	}

	pinger.OnDuplicateRecv = func(pkt *ping.Packet) {
		log.Printf("%d bytes from %s: icmp_seq=%d time=%v ttl=%v (DUP!)\n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt, pkt.Ttl)
	}

	pinger.OnFinish = func(stats *ping.Statistics) {
		fmt.Printf("\n--- %s ping statistics ---\n", stats.Addr)
		fmt.Printf("%d packets transmitted, %d packets received, %v%% packet loss\n",
			stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
		fmt.Printf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
			stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
		// TODO
		saveToJson(&utils.PingTest{
			URL:         stats.Addr,
			Transmitted: stats.PacketsSent,
			Received:    stats.PacketsRecv,
			Loss:        stats.PacketLoss,
		})
	}

	fmt.Printf("PING %s (%s):\n", pinger.Addr(), pinger.IPAddr())
	err = pinger.Run()
	if err != nil {
		panic(err)
	}

}

// Save ping result to the json file
// Note: this is ugly code, need to improve
// TODO: Find Best Practice and Improve it

func saveToJson(pingRes *utils.PingTest) {

	jsonFields := utils.Tests{}

	vpnFields, err := ioutil.ReadFile("data.json")
	if err != nil {
		os.Exit(1)
	}
	unmarshalErr := json.Unmarshal(vpnFields, &jsonFields)
	if unmarshalErr != nil {
		fmt.Println("Json unmarshalErr: ", unmarshalErr)
	}

	jsonFields.PingTest = *pingRes

	pingTestResult, marshalErr := json.MarshalIndent(jsonFields, "", "  ")
	if marshalErr != nil {
		fmt.Println("Json marshalErr: ", marshalErr)
	}

	ioutil.WriteFile("data.json", pingTestResult, 0644)

}
