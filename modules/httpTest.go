package modules

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// This module create by @fffaraz
// https://github.com/fffaraz

func TestHTTP(url string) {
	log.Println("URL:", url)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println("Error creating request:", url, err)
		return
	}

	transport := http.Transport{}
	client := http.Client{
		Transport: &transport,
		Timeout:   5 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			log.Println("Redirect:", req.URL)
			return http.ErrUseLastResponse
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request:", url, err)
		return
	}
	defer resp.Body.Close()

	log.Println("Response status:", resp.Status, resp.Proto)

	if resp.TLS != nil {
		log.Println("Response TLS version:", resp.TLS.Version)
		log.Println("Response TLS cipher suite:", resp.TLS.CipherSuite)
		log.Println("Response TLS server name:", resp.TLS.ServerName)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response:", url, err)
		return
	}

	for k, v := range resp.Header {
		log.Println("Response header:", k, v)
	}

	log.Println("Response length:", len(body))
	fmt.Println("------------------------------------------------------------")
}
