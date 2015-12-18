package cexcached

import "testing"

func TestCaching(t *testing.T) {
	cc := NewCexCached()

	// first query should be uncached
	c, cached, err := cc.ExchangeRate("USD", "ZAR")
	if err != nil {
		t.Error(err)
	} else if cached {
		t.Error("Got a cached result on first query")
	} else {
		t.Log("Got uncached result on first query")
		t.Log("1 USD =", c.Amount, "ZAR")
	}

	// second query should be cached
	c, cached, err = cc.ExchangeRate("USD", "ZAR")
	if err != nil {
		t.Error(err)
	} else if !cached {
		t.Error("Got an uncached result on second query")
	} else {
		t.Log("Got a cached result on second query")
		t.Log("1 USD =", c.Amount, "ZAR")
	}
}
