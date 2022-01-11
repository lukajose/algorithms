package hashtable

import (
	"collections/linkedlist"
	"math"
)

type HashTable struct {
	bucketsSize int
	elements    int
	buckets     []linkedlist.LinkedList
}

type BucketData struct {
	key   string
	value interface{}
}

func NewHashTable(size int) *HashTable {
	return &HashTable{
		bucketsSize: size,
		buckets:     make([]linkedlist.LinkedList, size),
	}
}

func NewBucketData(key string, value interface{}) *BucketData {
	return &BucketData{
		key:   key,
		value: value,
	}
}

func (bd BucketData) CompareTo(item interface{}) int {
	switch v := item.(type) {
	case string:
		if bd.key == v {
			return 0
		}
	case BucketData:
		if v == bd {
			return 0
		}
	default:
		return -1
	}
	return -1
}

func hash(s string) int {
	h := 0
	for pos, char := range s {
		h += int(char) * int(math.Pow(31, float64(len(s)-pos+1)))
	}
	return h
}

func (ht *HashTable) index(s string) int {
	h := hash(s)
	return h % ht.bucketsSize
}

func (ht *HashTable) Set(key string, value interface{}) {
	index := ht.index(key)
	data := NewBucketData(key, value)
	ht.buckets[index].Insert(data)
}

func (ht *HashTable) Get(key string) interface{} {
	index := ht.index(key)
	v := ht.buckets[index].Search(key)
	return v
}
