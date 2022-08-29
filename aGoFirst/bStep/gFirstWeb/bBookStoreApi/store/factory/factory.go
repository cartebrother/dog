package factory

import (
	"fmt"
	"github.com/dog/bBookStoreApi/store"
	"sync"
)

var (
	providersMu sync.RWMutex
	providers   = make(map[string]store.Store)
)

func Register(name string, p store.Store) {
	providersMu.Lock()
	defer providersMu.Unlock()
	if p == nil {
		panic("store :register provider is nil")
	}

	if _, dup := providers[name]; dup {
		panic("store: register called twice for provider " + name)
	}

	providers[name] = p
}

func New(providerName string) (store.Store, error) {
	providersMu.RLock()
	p, ok := providers[providerName]
	providersMu.RUnlock()

	if !ok {
		return nil, fmt.Errorf("store:unkown provider %s", providerName)
	}
	return p, nil
}
