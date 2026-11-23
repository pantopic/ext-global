package wazero_global

import (
	"os"
)

type HostModule interface {
	LayerAdd(fn Lookup)
	Set(key string, val uint64)
	Del(key string)
}

type Option func(*hostModule)

type Lookup func(key string) (val string, exists bool)

func WithOverlay(fn Lookup) Option {
	return func(h *hostModule) {
		h.LayerAdd(fn)
	}
}

func EnvLookup(prefix string) Lookup {
	return func(key string) (string, bool) {
		return os.LookupEnv(prefix + key)
	}
}
