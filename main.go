package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel *Channel `xml:"channel"`
}

type Channel struct {
	Title    string `xml:"title"`
	ItemList []Item `xml:"item"`
}

type Item struct {
	Title     string `xml:"title"`
	Link      string `xml:"link"`
	Traffic   string `xml:"approx_traffic"`
	NewsItems []News `xml:"news_item"`
}

type News struct {
	Headline     string `xml:"news_item_title"`
	HeadlineLink string `xml:"news_item_url"`
}

func main() {
	var r RSS

	data := readGoogleTrends()

	err := xml.Unmarshal(data, &r)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Println("Below are all the google trends for today")
	fmt.Println("--------------------")
	for i := range r.Channel.ItemList {
		rank := i + 1
		fmt.Println("Rank: ", rank)
		fmt.Println("Title: ", r.Channel.ItemList[i].Title)
		fmt.Println("Link: ", r.Channel.ItemList[i].Link)
		fmt.Println("Traffic: ", r.Channel.ItemList[i].Traffic)
		fmt.Println("First result: ", r.Channel.ItemList[i].NewsItems[0].Headline)
		fmt.Println("First result link: ", r.Channel.ItemList[i].NewsItems[0].HeadlineLink)
		fmt.Println("--------------------")
	}
}

func getGoogleTrends(l string) *http.Response {
	url := "https://trends.google.com.br/trends/trendingsearches/daily/rss?geo=" + l
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	return resp
}

func readGoogleTrends() []byte {
	resp := getGoogleTrends("BR")

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	return data
}
