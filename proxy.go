package main

import (
	"context"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
	"github.com/scrapeless-ai/sdk-go/scrapeless/services/proxies"
	"net/url"
)

func getProxy(ctx context.Context) (*url.URL, error) {
	// Get proxy
	proxy, err := actor.Proxy.Proxy(ctx, proxies.ProxyActor{
		Country:         "us",
		SessionDuration: 10,
	})

	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	parse, err := url.Parse(proxy)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return parse, nil
}
