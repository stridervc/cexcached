package cexcached

import (
	"github.com/stridervc/cex-go"
	"time"
)

type CexCached struct {
	APIWait    int64 // minimum delay in seconds between API requests
	CacheValid int64 // time in seconds that cached data is valid for

	lastAPICall int64 // timestamp of last API call

	// store queried currency data based on source and target
	cache map[string]cex.CurrencyData

	// timestamp of last query for each source,target
	cacheTime map[string]int64
}

func NewCexCached() CexCached {
	c := CexCached{}

	// set default values
	c.APIWait = 1
	c.CacheValid = 300

	// create maps
	c.cache = make(map[string]cex.CurrencyData)
	c.cacheTime = make(map[string]int64)

	return c
}

func (c *CexCached) ExchangeRate(source, target string) (cex.CurrencyData, bool, error) {
	now := time.Now().Unix()
	key := source + "+" + target

	// if we have cached data that's valid, return it
	if now-c.cacheTime[key] <= c.CacheValid {
		return c.cache[key], true, nil
	}

	// wait until we're allowed to query
	for now-c.lastAPICall < c.APIWait {
		time.Sleep(time.Second)
		now = time.Now().Unix()
	}

	// query and cache result
	data, err := cex.ExchangeRate(source, target)
	if err != nil {
		return data, false, err
	}

	c.lastAPICall = time.Now().Unix()
	c.cacheTime[key] = c.lastAPICall
	c.cache[key] = data

	// return result
	return data, false, nil
}
