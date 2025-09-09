package pokecache

import (
	"testing"
	"time"
)

func TestCacheAdd(t *testing.T) {
	sut := NewCache(5 * time.Second)
	sut.Add("Key", []byte("data"))
	if len(sut.data) != 1 {
		t.Errorf("Actual length of %v doesn't match expected length of %v", len(sut.data), 1)
	}
}

func TestCacheGet(t *testing.T) {
	sut := NewCache(5 * time.Second)
	sut.Add("Key",[]byte("data"))
	val, ok := sut.Get("Key")
	if !ok {
		t.Errorf("Cache.Get returned no data")
	}
	if string(val) != "data" {
		t.Errorf("Cache did not return the correct data")
	}
}

func TestReaping(t * testing.T) {
	sut := NewCache(1 * time.Microsecond)
	sut.Add("Key", []byte("data"))
	time.Sleep(15 * time.Microsecond)
	_, ok := sut.Get("Key")
	if ok {
		t.Errorf("Reaping did not remove the key")
	}
}
