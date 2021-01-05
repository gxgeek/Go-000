package main

import (
	"container/ring"
	"fmt"
	"sync"
	"time"
)

type Bucket struct {
	Value int
}

type RollWindow struct {
	mux sync.RWMutex
	//size    int64 // buckets记录时间长度
	now int64			// 当前窗口最近的桶的秒数，每个桶1s，该数据作为索引查找桶
	buckets *ring.Ring	//链表

}

func NewRollWindow(count int) *RollWindow {
	r := & RollWindow{
		buckets: ring.New(count),
	}
	clearBuckets(r)
	return r
}

func clearBuckets(r *RollWindow) {
	for i := 0; i < r.buckets.Len(); i++ {
		if r.buckets.Value == nil {
			r.buckets.Value = &Bucket{}
		}else {
			r.buckets.Value.(*Bucket).Value = 0

		}
		r.buckets = r.buckets.Next()
	}
}


func (r *RollWindow) IncrValue()  {
	r.updateTime()

	r.mux.Lock()
	defer r.mux.Unlock()
	r.buckets.Value.(*Bucket).Value++
}

func (r *RollWindow) GetCounts() int {
	r.updateTime()

	r.mux.Lock()
	defer r.mux.Unlock()
	//r.counts.Success++
	var cur = r.buckets
	var count = 0
	for i := 0; i < cur.Len(); i++ {
		count += cur.Value.(*Bucket).Value
		cur = cur.Next()
	}
	return count
}

func (r *RollWindow) updateTime() {
	now := time.Now().Unix()
	r.mux.Lock()
	defer r.mux.Unlock()
	//超出窗口 clear
	if (now - r.now) >= int64(r.buckets.Len()){
		clearBuckets(r)
	}else if now - r.now > int64(0) {
		n := now - r.now
		for i := int64(0); i < n ;i++ {
			r.buckets = r.buckets.Next()
			r.buckets.Value.(*Bucket).Value = 0
		}
	}
	r.now = now

}

func main()  {
	window := NewRollWindow(5)
	for i := 0; i < 10 ; i++ {
		window.IncrValue()
		fmt.Println(window.GetCounts())
		time.Sleep(time.Second)
	}
}
