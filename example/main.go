package main

import (
	web_server "example/web_test"
	"fmt"
	"net/http"
	"os"
	"time"
)

type sku struct {
	item  string
	price string
}

var items = []sku{
	{"shoes", "46"},
	{"socks", "6"},
	{"sandals", "27"},
	{"clogs", "36"},
	{"pants", "30"},
	{"shorts", "20"},
}

func doQuery(cmd, parms string) {
	resp, err := http.Get("http://localhost:8080/" + cmd + "?" + parms)

	if err == nil {
		defer resp.Body.Close()
		fmt.Fprintf(os.Stderr, "got %s = %d (no err)\n", parms, resp.StatusCode)
	} else if resp != nil {
		defer resp.Body.Close()
		fmt.Fprintf(os.Stderr, "got %s = %d (%v)\n", parms, resp.StatusCode, err)
	} else {
		fmt.Fprintf(os.Stderr, "got err %v\n", err)
	}
}

func runAdds() {
	for {
		for _, s := range items {
			doQuery("create", "item="+s.item+"&price="+s.price)
		}
	}
}

func runUpdates() {
	for {
		for _, s := range items {
			doQuery("update", "item="+s.item+"&price="+s.price)
		}
	}
}

func runDrops() {
	for {
		for _, s := range items {
			doQuery("create", "item="+s.item)
		}
	}
}

func main() {
	go web_server.RunServer()
	go runAdds()
	go runDrops()
	go runUpdates()

	time.Sleep(5 * time.Second)
}
