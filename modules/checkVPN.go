package modules

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"regexp"
)

func CheckVPN(ipCheker string) {
	resp, err := http.Get(ipCheker)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	re := regexp.MustCompile(`IP Address: (\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})`)
	externalIP := re.FindStringSubmatch(string(body))[1]
	localIP, err := net.LookupHost("localhost")
	if err != nil {
		log.Println(err)
		return
	}
	if externalIP == localIP[0] {
		log.Println("Not using VPN or proxy.")
		fmt.Println("------------------------------------------------------------")
	} else {
		log.Println("Using VPN or proxy.")
		fmt.Println("------------------------------------------------------------")
	}
}
