package main

import (
	"encoding/base64"
	"fmt"

	gmc "github.com/bradfitz/gomemcache/memcache"
)

type sampleCacheEntry struct {
	name     string
	key      string
	b64Value string
}

var sampleCacheEntries = []sampleCacheEntry{
	{
		name: "basic model",
	},
}

var BootstrapMemcachedCommentCache = func(entries []sampleCacheEntry) func(client *gmc.Client) error {
	return func(client *gmc.Client) error {
		for _, entry := range sampleCacheEntries {
			value, err := base64.StdEncoding.DecodeString(entry.b64Value)
			if err != nil {
				return fmt.Errorf("decoding string: %w", err)
			}

			if err := client.Add(&gmc.Item{
				Key:   entry.key,
				Value: value,
			}); err != nil {
				return fmt.Errorf("inserting item: %w", err)
			}
		}
		return nil
	}
}
