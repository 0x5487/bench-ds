package btree

import (
	"bytes"

	"github.com/tidwall/btree"
)

type Item struct {
	key   []byte
	value []byte
}

type BTree struct {
	btree *btree.BTreeG[*Item]
}

func NewBTree() *BTree {
	return &BTree{
		btree: btree.NewBTreeG[*Item](func(a, b *Item) bool {
			return bytes.Compare(a.key, b.key) == -1
		}),
	}
}

func (bt *BTree) Set(key []byte, v []byte) {
	bt.btree.Set(&Item{key: key, value: v})
}

func (bt *BTree) Get(key []byte) ([]byte, bool) {
	item, ok := bt.btree.Get(&Item{key: key})
	if ok {
		return item.value, ok
	}
	return nil, ok
}
