package main

import (
	"context"
	"github.com/scrapeless-ai/sdk-go/scrapeless"
	scrapeless_actor "github.com/scrapeless-ai/sdk-go/scrapeless/actor"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
	"io"
	"net/http"
	"net/url"
)

type RequestParam struct {
	Url string `json:"url"`
}

var (
	// New scrapeless client
	client *scrapeless.Client
	// New scrapeless actor
	actor *scrapeless_actor.Actor
)

func main() {
	// New scrapeless actor
	actor = scrapeless_actor.New()
	defer actor.Close()
	// New scrapeless client
	client = scrapeless.New(scrapeless.WithDeepSerp())
	defer client.Close()

	var param = &RequestParam{}
	if err := actor.Input(param); err != nil {
		return
	}
	// Get proxy
	proxy, proxyErr := getProxy(context.Background())
	if proxyErr != nil {
		return
	}
	data, requestErr := doRequestWithProxy(param, proxy)
	if requestErr != nil {
		return
	}
	addErr := datasetAddItem(context.Background(), []map[string]interface{}{{"data": data}})
	if addErr != nil {
		return
	}
}

// Use proxy for http access
func doRequestWithProxy(param *RequestParam, proxy *url.URL) ([]byte, error) {
	// Set up proxy using Golang's native HTTP
	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxy)}}
	// get response
	resp, err := client.Get(param.Url)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return body, nil
}
