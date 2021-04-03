package design

import (
	"encoding/json"
	"time"
)

type Keyword struct {
	word      string
	visit     int
	UpdatedAt *time.Time
}

// Clone 这里使用序列化与反序列化的方式深拷贝
func (k *Keyword) Clone() *Keyword {
	var newKeyword Keyword
	b, _ := json.Marshal(k)
	json.Unmarshal(b, &newKeyword)
	return &newKeyword
}

type Keywords map[string]*Keyword

// Clone 复制一个新的 keywords
// updatedWords: 需要更新的关键词列表，由于从数据库中获取数据常常是数组的方式
func (words Keywords) Clone(updateWords []*Keyword) Keywords {
	newKeywords := Keywords{}
	for k, v := range words {
		// 这里是浅拷贝，直接拷贝了地址
		newKeywords[k] = v
	}
	// 替换需要更新的字段，这里用的深拷贝
	for _, word := range updateWords {
		newKeywords[word.word] = word.Clone()
	}
	return newKeywords
}
