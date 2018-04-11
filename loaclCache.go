// loaclCache project loaclCache.go
package loaclCache

import "sync"
import "time"

type Item struct {
	v     interface{}
	timer int64
}

type Cache struct {
	cache map[string]Item
	mu    sync.RWMutex
}

var c Cache

func init() {

	c.cache = make(map[string]Item)
	go reclaim()

}
func reclaim() {

	for {

		time.Sleep(time.Second * 30)
		c.mu.Lock()
		for k, v := range c.cache {
			if v.timer < time.Now().Unix() {

				delete(c.cache, k)

			}

		}
		c.mu.Unlock()

	}

}

func Get(k string) (interface{}, bool) {
	c.mu.RLock()
	x, ok := c.cache[k]
	c.mu.RUnlock()
	if ok {
		if x.timer > time.Now().Unix() {
			return x.v, ok
		}

	}
	return nil, ok

}
func Set(k string, v interface{}, t int64) {
	x := time.Now().Unix() + t
	c.mu.Lock()
	c.cache[k] = Item{v, x}
	c.mu.Unlock()

}
