package _022_06

import "fmt"

// TinyURL 的加密与解密
// leetcode: https://leetcode.cn/problems/encode-and-decode-tinyurl/
type Codec struct {
	m  map[string]string
	id int
}

func Constructor() Codec {
	return Codec{m: map[string]string{}}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	url := fmt.Sprintf("http://tinyurl.com/%d", this.id)
	this.m[url] = longUrl
	return url
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	return this.m[shortUrl]
}
