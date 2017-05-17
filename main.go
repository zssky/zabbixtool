package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/zssky/log"
	"github.com/zssky/tc/http"
)

var (
	url = flag.String("url", "", "Request url")
)

func main() {

	flag.Parse()
	if *url == "" {
		log.Errorf("url can not be null")
		return
	}

	data, _, err := http.SimpleGet(*url, time.Second*5, time.Second*15)
	if err != nil {
		log.Errorf("err%v", err)
		return
	}

	fmt.Printf("%s", string(data))
}
