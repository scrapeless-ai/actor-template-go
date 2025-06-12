package main

import (
	"context"
	scrapeless_actor "github.com/scrapeless-ai/sdk-go/scrapeless/actor"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
	"github.com/scrapeless-ai/sdk-go/scrapeless/services/proxies"
	"io"
	"net/http"
	"net/url"
)

type RequestParam struct {
	Url string `json:"url"`
}

func main() {
	// New scrapeless actor
	actor := scrapeless_actor.New()
	defer actor.Close()
	var param = &RequestParam{}
	if err := actor.Input(param); err != nil {
		log.Error(err.Error())
		return
	}
	// Get proxy
	proxy, err := actor.Proxy.Proxy(context.TODO(), proxies.ProxyActor{
		Country:         "us",
		SessionDuration: 10,
	})

	if err != nil {
		log.Error(err.Error())
		return
	}
	parse, err := url.Parse(proxy)
	if err != nil {
		log.Error(err.Error())
		return
	}
	// Set up proxy using Golang's native HTTP
	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(parse)}}
	// get response
	resp, err := client.Get(param.Url)
	if err != nil {
		log.Error(err.Error())
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error(err.Error())
		return
	}
	// use dataset
	items, err := actor.AddItems(context.Background(), []map[string]interface{}{
		{"body": string(body)},
	})
	if err != nil {
		log.Error(err.Error())
		return
	}
	log.Info(items)
}
