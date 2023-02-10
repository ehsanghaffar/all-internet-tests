package main

import (
	"github.com/ehsanghaffar/ultimate-internet-test/modules"
)

func main() {

	// flag.Parse()

	// log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	// args := flag.Args()
	// if len(args) > 0 {
	// 	for _, url := range args {
	// 		modules.TestHTTP(url)
	// 	}
	// 	return
	// }

	modules.TestHTTP("http://www.google.com/")
	modules.TestHTTP("https://www.google.com/")
	modules.TestHTTP("https://www.facebook.com/")
	modules.TestHTTP("https://www.youtube.com/")
	modules.TestHTTP("https://leader.ir/")

	modules.CheckSpeed("https://ehsanghaffarii.ir")
	modules.CheckSpeed("https://google.com")

	modules.CheckVPN("http://checkip.dyndns.org/")

	modules.PingCheck("www.google.com")

}
