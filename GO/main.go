package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Config struct {
	Vouch struct {
		MyVouchURL  string `json:"MyVouch_URL"`
		RequestDelay int    `json:"Request_Delay"`
	} `json:"Vouch"`
}

func fetchVouchesCount(config Config) int {
	resp, err := http.Get(config.Vouch.MyVouchURL)
	if err != nil {
		fmt.Println("Failed to fetch the vouch count:", err)
		return -1
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("Failed to parse response body:", err)
		return -1
	}

	vouchesText := doc.Find("p.social span:last-child").Text()
	re := regexp.MustCompile(`\d+`)
	vouchesCountStr := re.FindString(vouchesText)
	vouchesCount, _ := strconv.Atoi(vouchesCountStr)
	return vouchesCount
}

func printVouchesCount(config Config) {
	count := fetchVouchesCount(config)
	fmt.Println("Vouch count:", count)
}

func main() {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println("Failed to read config file:", err)
		return
	}

	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println("Failed to parse config file:", err)
		return
	}

	printVouchesCount(config)
	requestDelay := time.Duration(config.Vouch.RequestDelay) * time.Second
	ticker := time.NewTicker(requestDelay)
	defer ticker.Stop()
	for range ticker.C {
		printVouchesCount(config)
	}
}
