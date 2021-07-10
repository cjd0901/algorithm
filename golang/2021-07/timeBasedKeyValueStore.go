package _021_07

import "sort"

// Time Based Key-Value Store
// leetcode: https://leetcode-cn.com/problems/time-based-key-value-store/
type Pair struct {
	timestamp int
	value     string
}

type TimeMap struct {
	m map[string][]Pair
}

func Constructor() TimeMap {
	return TimeMap{m: map[string][]Pair{}}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.m[key] = append(this.m[key], Pair{
		timestamp,
		value,
	})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	pairs := this.m[key]
	i := sort.Search(len(pairs), func(i int) bool {
		return pairs[i].timestamp > timestamp
	})
	if i > 0 {
		return pairs[i-1].value
	}
	return ""
}
