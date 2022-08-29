package store

import (
	mystore "github.com/dog/bBookStoreApi/store"
	factory "github.com/dog/bBookStoreApi/store/factory"
	"sync"
)

func init() {
	factory.Register("mem", &MemStore{
		books: make(map[string]*mystore.Book),
	})
}

type MemStore struct {
	sync.RWMutex
	books map[string]*mystore.Book
}
