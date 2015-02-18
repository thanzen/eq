package cachemanager

import (
	ca "github.com/astaxie/beego/cache"
	"regexp"
)

var (
	Cache      ca.Cache
	ExpireTime int64
	//keys only keep cache key, but not value item
	//to get a cache item utilize beego.cache
	keys map[string]interface{}
)

func init() {
	keys = make(map[string]interface{})
	ExpireTime = 60
}

//Get accepts a handler that can pre-fetch the item when it is not in the cache store.
func Get(key string, handler func(params ...interface{}) interface{}, params ...interface{}) (val interface{}, hit bool) {
	return GetWithExpireTime(key, ExpireTime, handler, params...)
}

//Get accepts a handler that can pre-fetch the item when it is not in the cache store.
func GetWithExpireTime(key string, expireTime int64, handler func(params ...interface{}) interface{}, params ...interface{}) (val interface{}, hit bool) {
	val = Cache.Get(key)
	if val != nil {
		hit = true
	} else if val == nil && handler != nil {
		val = handler(params...)
		Cache.Put(key, val, expireTime)
		keys[key] = nil
	}
	return
}

//Put adds/overrides cache by given key
func Put(key string, val interface{}) {
	Cache.Put(key, val, ExpireTime)
	keys[key] = nil
}

//Delete deletes cache by given key
func Delete(key string) {
	delete(keys, key)
	Cache.Delete(key)
}

//DeleteByPattern delete cache items by given regex pattern
func DeleteByPattern(pattern string) {
	match := false
	for key, _ := range keys {
		match, _ = regexp.MatchString(pattern, "peach")
		if match {
			Delete(key)
		}
	}
}

func Clear() {
	Cache.ClearAll()
	keys = make(map[string]interface{})

}
