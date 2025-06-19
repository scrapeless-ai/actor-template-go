package main

import (
	"context"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
	"github.com/scrapeless-ai/sdk-go/scrapeless/services/deepserp"
)

func deepserpTrend(ctx context.Context) []byte {
	params := map[string]interface{}{
		"q":         "Mercedes-Benz,BMW X5",
		"data_type": "interest_over_time",
		"date":      "today 1-m",
		"hl":        "en-sg",
		"tz":        "-480",
	}

	result, err := client.DeepSerp.Scrape(ctx, deepserp.DeepserpTaskRequest{
		Actor:        "scraper.google.trends",
		Input:        params,
		ProxyCountry: "US",
	})
	if err != nil {
		log.Warnf("deepserp failed: %v", err)
	}
	log.Infof("deepserp trend data:\n%s", string(result))
	return result
}
