package pokeCache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://somethingsomething.example.com",
			val: []byte("someTestData"),
		},
		{
			key: "https://comethingother.example.com",
			val: []byte("someOtherTestData"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)

			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find a key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})

	}
}

func TestReapLoop(t *testing.T) {
	baseTime := 5 * time.Millisecond
	waitTime := 2 * baseTime

	cache := NewCache(baseTime)
	cache.Add("https://somethingsomething.example.com", []byte("someTestData"))

	_, ok := cache.Get("https://somethingsomething.example.com")
	if !ok {
		t.Errorf("expected to find value")
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://somethingsomething.example.com")
	if ok {
		t.Errorf("expected not to find value")
	}
}
