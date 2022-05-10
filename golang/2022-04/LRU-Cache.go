package _022_04

// leetcode: https://leetcode-cn.com/problems/lru-cache/
type LRUCache struct {
	count      int                // 当前缓存数
	capacity   int                // 缓存容量
	cacheDict  map[int]*CacheLink // 保存缓存数据的map
	head, tail *CacheLink         // 双向链表
}

type CacheLink struct {
	key, value int
	prev, next *CacheLink
}

func Constructor(capacity int) LRUCache {
	cacheDict := make(map[int]*CacheLink)
	head := initCacheLink(0, 0)
	tail := initCacheLink(0, 0)
	head.next = tail
	tail.prev = head
	return LRUCache{
		capacity:  capacity,
		cacheDict: cacheDict,
		head:      head,
		tail:      tail,
	}
}

func initCacheLink(key, value int) *CacheLink {
	return &CacheLink{
		key:   key,
		value: value,
	}
}

func (this *LRUCache) Get(key int) int {
	value, ok := this.cacheDict[key]
	if !ok {
		return -1
	}
	this.MoveHead(value)
	return value.value
}

func (this *LRUCache) Put(key int, value int) {
	cache, ok := this.cacheDict[key]
	if !ok {
		this.count++
		cache = initCacheLink(key, value)
		this.cacheDict[key] = cache
	}
	cache.value = value
	this.MoveHead(cache)
	if this.count > this.capacity {
		end := this.tail.prev
		end.next.prev = end.prev
		end.prev.next = end.next
		delete(this.cacheDict, end.key)
		this.count--
	}
}

func (this *LRUCache) MoveHead(cache *CacheLink) {
	if cache.prev != nil && cache.next != nil {
		cache.next.prev = cache.prev
		cache.prev.next = cache.next
	}
	cache.next = this.head.next
	cache.prev = this.head
	this.head.next.prev = cache
	this.head.next = cache
}
