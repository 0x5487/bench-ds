package btree

import (
	"bench-ds/data"
	"fmt"
	"testing"

	"github.com/tidwall/btree"
)

func TestBtree(t *testing.T) {
	tree := &btree.Map[string, []byte]{}

	for i := 0; i < 10; i++ {
		key := string(data.GetKey(i))
		tree.Set(key, data.GetValue())
	}

	tree.Scan(func(key string, value []byte) bool {
		fmt.Println(key, string(value))
		return true
	})
}

func Benchmark_Btree_Set(b *testing.B) {
	tree := &btree.Map[string, []byte]{}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		key := string(data.GetKey(i))
		tree.Set(key, data.GetValue())
	}
}

func Benchmark_Btree_Get(b *testing.B) {
	tree := &btree.Map[string, []byte]{}
	fillData(tree)

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		key := string(data.GetKey(i))
		tree.Get(key)
	}
}

func fillData(tree *btree.Map[string, []byte]) {
	for i := 0; i < 10000; i++ {
		key := string(data.GetKey(i))
		tree.Set(key, data.GetValue())
	}
}

func Benchmark_BtreeG_Set(b *testing.B) {
	tree := NewBTree()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		tree.Set(data.GetKey(i), data.GetValue())
	}
}

func Benchmark_BtreeG_Get(b *testing.B) {
	tree := NewBTree()

	for i := 0; i < 50000; i++ {
		tree.Set(data.GetKey(i), data.GetValue())
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		tree.Get(data.GetKey(i))
	}
}
