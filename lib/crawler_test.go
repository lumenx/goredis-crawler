package crawler

import (
	"fmt"
	"testing"
)

func TestGeneral(t *testing.T) {

	crawl, err := New("http://rpiai.com/")
	if err != nil {
		t.Error(err)
	}

	crawl.RedisURL = "localhost"
	crawl.RedisPort = "6379"
	crawl.Verbose = true

	err = crawl.Init()
	if err != nil {
		t.Error(err)
	}

	urls, err := crawl.scrapeLinks("http://rpiai.com")
	if err != nil {
		t.Error(err)
	}
	if len(urls) < 15 {
		t.Error("%v", urls)
	}

	err = crawl.Crawl()
	if err != nil {
		t.Error(err)
	}
}

func TestProxy(t *testing.T) {
	crawl, err := New("http://rpiai.com/")
	if err != nil {
		t.Error(err)
	}
	crawl.RedisURL = "localhost"
	crawl.RedisPort = "6378"
	crawl.Verbose = true
	crawl.Init()
	ip, err := crawl.getIP()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ip)
	urls, err := crawl.scrapeLinks("http://rpiai.com")
	if err != nil {
		t.Error(err)
	}
	if len(urls) < 15 {
		t.Error("%v", urls)
	}
}
