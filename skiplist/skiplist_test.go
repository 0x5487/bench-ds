package skiplist

import (
	"bench-ds/data"
	"fmt"
	"testing"

	"github.com/huandu/skiplist"
)

func TestSkiplist(t *testing.T) {
	list := skiplist.New(skiplist.BytesDesc)

	for i := 0; i < 10; i++ {
		list.Set(data.GetKey(i), i)
	}

	element := list.Front()
	for element != nil {
		fmt.Println(element.Value)
		element = element.Next()
	}
}

func Benchmark_Skiplist_Set(b *testing.B) {
	// Create a skip list with int key.
	list := skiplist.New(skiplist.BytesDesc)

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		list.Set(data.GetKey(i), data.GetValue())
	}
}

func Benchmark_Skiplist_Get(b *testing.B) {
	// Create a skip list with int key.
	list := skiplist.New(skiplist.BytesDesc)
	fillData(list)

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		list.Get(data.GetKey(i))
	}
}

func fillData(list *skiplist.SkipList) {
	for i := 0; i < 10000; i++ {
		list.Set(data.GetKey(i), data.GetValue())
	}
}

func Benchmark_Skiplist_StringGet(b *testing.B) {
	// Create a skip list with int key.
	list := skiplist.New(skiplist.StringAsc)
	for i := 0; i < 50000; i++ {
		key := string(data.GetKey(i))
		list.Set(key, data.GetValue())
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		key := string(data.GetKey(i))
		list.Get(key)
	}
}
