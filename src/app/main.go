package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

var Revision string

type Site struct {
	Number       int
	URL          string
	HTTPCode     int
	ResponseTime time.Duration
	StatusMsg    string
}

func CheckSiteStatus(site *Site) {
	start := time.Now()
	resp, err := http.Get(site.URL)
	elapsed := time.Since(start)

	if err != nil {
		site.HTTPCode = -1
		site.StatusMsg = "NOK"
	} else {
		site.HTTPCode = resp.StatusCode
		site.StatusMsg = strings.Split(resp.Status, " ")[1]
	}

	site.ResponseTime = elapsed.Round(time.Millisecond)
}

func main() {
	fmt.Println("Server Status Checker:")
	fmt.Println()
	fmt.Println("No.\tURLs\t\t\tCode\tTime\tStatus")

	sites := []Site{
		{URL: "https://google.com"},
		{URL: "https://apple.com"},
		{URL: "https://yahoo.com"},
		{URL: "https://freecoder.dev"},
	}

	for i, site := range sites {
		site.Number = i + 1
		CheckSiteStatus(&site)
		fmt.Printf("%d\t%s\t%d\t%.2fs\t%s\n",
			site.Number,
			site.URL,
			site.HTTPCode,
			site.ResponseTime.Seconds(),
			site.StatusMsg)
	}
}
