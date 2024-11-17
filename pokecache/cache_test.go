package pokecache

import (
	"fmt"
	"testing"
)

func TestAddToCache(t *testing.T) {
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for index, element := range cases {
		t.Run(fmt.Sprintf("Test case %v", index), func(t *testing.T) {
			cacheInstance := CacheInstance()
			cacheInstance.Add(element.key, element.val, 10)
			val, ok := cacheInstance.Get(element.key)

			if !ok {
				t.Errorf("expected to find key")
				return
			}

			if string(val) != string(element.val) {
				t.Errorf("values differs")
				return
			}
		})
	}
}

func TestDeleteFromCache(t *testing.T) {
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for index, element := range cases {
		t.Run(fmt.Sprintf("Test case %v", index), func(t *testing.T) {
			cacheInstance := CacheInstance()
			cacheInstance.Add(element.key, element.val, 10)

			cacheInstance.Delete(element.key)

			_, found := cacheInstance.Get(element.key)
			if found {
				t.Errorf("expected to not find key")
				return
			}
		})
	}
}
