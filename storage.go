package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/chai2010/webp"
	"github.com/scrapeless-ai/sdk-go/scrapeless/log"
	"github.com/scrapeless-ai/sdk-go/scrapeless/services/storage/kv"
	"image/png"
	"net/http"
	"time"
)

// setKv Set kv
func setKv(ctx context.Context, k, v string) error {
	value, err := actor.SetValue(ctx, k, v, 3600)
	if err != nil {
		return fmt.Errorf("set kv failed: %w", err)
	}
	if !value {
		return fmt.Errorf("set kv failed")
	}
	log.Infof("set kv success, key: %v, value: %v", k, v)
	return nil
}

// bulkSetKv Set kv in batch
func bulkSetKv(ctx context.Context, items []kv.BulkItem) error {
	count, err := actor.BulkSetValue(ctx, items)
	if err != nil {
		return fmt.Errorf("bulk set kv failed: %w", err)
	}
	if count <= 0 {
		return fmt.Errorf("bulk set kv success count is 0")
	}
	log.Infof("bulk set kv success, count: %v", count)
	return nil
}

// objectPut Save object
func objectPut(ctx context.Context) error {
	pngBytes, err := downloadWebpAsPngBytes("https://banner2.cleanpng.com/20180408/vae/avgpocfjw.webp")
	if err != nil {
		log.Warnf("download image failed: %v", err)
	}
	value, putErr := actor.PutObject(ctx, "demo.png", pngBytes)
	if putErr != nil {
		log.Warnf("save object failed: %v", putErr)
		return putErr
	}
	log.Infof("save object success, object: %v", value)
	return nil
}

func datasetAddItem(ctx context.Context, items []map[string]interface{}) error {
	isAdd, addErr := actor.AddItems(context.Background(), items)
	if addErr != nil {
		log.Error(addErr.Error())
		return addErr
	}
	if !isAdd {
		log.Error("add item failed")
		return fmt.Errorf("add item failed")
	}
	log.Info("add item success")
	return nil
}

// downloadWebpAsPngBytes download webp as png bytes
func downloadWebpAsPngBytes(url string) ([]byte, error) {
	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	// 下载 .webp 图片
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("http get error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status: %s", resp.Status)
	}

	// 解码 WebP -> image.Image
	img, err := webp.Decode(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("webp decode error: %w", err)
	}

	// 编码为 PNG 并写入 buffer
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return nil, fmt.Errorf("png encode error: %w", err)
	}

	return buf.Bytes(), nil
}
