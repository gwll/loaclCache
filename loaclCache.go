// loaclCache project loaclCache.go
package loaclCache

import "sync"
import "time"

type Item struct { //缓存项
	v     interface{}
	timer int64
}

type Cache struct { //缓存池
	cache map[string]Item
	mu    sync.RWMutex
}

var c Cache

func init() {

	c.cache = make(map[string]Item)
	//go Reclaim()

}
func Reclaim() { //定期清理过期缓存内容

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

func Get(k string) (interface{}, bool) { //获取缓存内容
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
func Set(k string, v interface{}, t int64) { //设置缓存内容
	x := time.Now().Unix() + t
	c.mu.Lock()
	c.cache[k] = Item{v, x}
	c.mu.Unlock()

}
