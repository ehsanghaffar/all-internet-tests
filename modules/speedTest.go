package modules

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func CheckSpeed(url string) {
	startTime := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	elapsedTime := time.Since(startTime)
	speed := float64(len(body)) / elapsedTime.Seconds()
	log.Println("URL:", url)
	log.Printf("Download speed: %.2f Mbps\n", speed/1000000*8)
	log.Printf("Elapsed time: %s\n", elapsedTime)
	fmt.Println("------------------------------------------------------------")
}
